// Copyright (c) 2018-2021 Splunk Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	spltest "github.com/splunk/splunk-operator/pkg/splunk/test"
	splutil "github.com/splunk/splunk-operator/pkg/splunk/util"

	enterpriseApi "github.com/splunk/splunk-operator/pkg/apis/enterprise/v2"
	splcommon "github.com/splunk/splunk-operator/pkg/splunk/common"
)

func TestApplyStatefulSet(t *testing.T) {
	funcCalls := []spltest.MockFuncCall{{MetaName: "*v1.StatefulSet-test-splunk-stack1-indexer"}}
	createCalls := map[string][]spltest.MockFuncCall{"Get": funcCalls, "Create": funcCalls}
	updateCalls := map[string][]spltest.MockFuncCall{"Get": funcCalls, "Update": funcCalls}
	var replicas int32 = 1
	current := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "splunk-stack1-indexer",
			Namespace: "test",
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: &replicas,
		},
	}
	revised := current.DeepCopy()
	revised.Spec.Template.ObjectMeta.Labels = map[string]string{"one": "two"}
	reconcile := func(c *spltest.MockClient, cr interface{}) error {
		_, err := ApplyStatefulSet(c, cr.(*appsv1.StatefulSet))
		return err
	}
	spltest.ReconcileTester(t, "TestApplyStatefulSet", current, revised, createCalls, updateCalls, reconcile, false)
}

func TestDefaultStatefulSetPodManager(t *testing.T) {
	// test for updating
	mgr := DefaultStatefulSetPodManager{}
	method := "DefaultStatefulSetPodManager.Update"
	spltest.PodManagerTester(t, method, &mgr)
}

func updateStatefulSetPodsTester(t *testing.T, mgr splcommon.StatefulSetPodManager, statefulSet *appsv1.StatefulSet, desiredReplicas int32, initObjects ...runtime.Object) (splcommon.Phase, error) {
	// initialize client
	c := spltest.NewMockClient()
	c.AddObjects(initObjects)
	phase, err := UpdateStatefulSetPods(c, statefulSet, mgr, desiredReplicas)
	return phase, err
}

func TestUpdateStatefulSetPods(t *testing.T) {
	mgr := DefaultStatefulSetPodManager{}
	var replicas int32 = 1
	statefulSet := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "splunk-stack1",
			Namespace: "test",
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: &replicas,
			VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
				{ObjectMeta: metav1.ObjectMeta{Name: "pvc-etc", Namespace: "test"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "pvc-var", Namespace: "test"}},
			},
		},
		Status: appsv1.StatefulSetStatus{
			Replicas:        replicas,
			ReadyReplicas:   replicas,
			UpdatedReplicas: replicas,
			UpdateRevision:  "v1",
		},
	}
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "splunk-stack1-0",
			Namespace: "test",
			Labels: map[string]string{
				"controller-revision-hash": "v0",
			},
		},
		Status: corev1.PodStatus{
			Phase: corev1.PodRunning,
			ContainerStatuses: []corev1.ContainerStatus{
				{Ready: true},
			},
		},
	}

	var phase splcommon.Phase
	phase, err := updateStatefulSetPodsTester(t, &mgr, statefulSet, 1 /*desiredReplicas*/, statefulSet, pod)
	if err != nil && phase != splcommon.PhaseUpdating {
		t.Errorf("UpdateStatefulSetPods should not have returned error=%s with phase=%s", err, phase)
	}

	// readyReplicas < replicas
	replicas = 3
	statefulSet.Status.ReadyReplicas = 2
	statefulSet.Spec.Replicas = &replicas
	phase, err = updateStatefulSetPodsTester(t, &mgr, statefulSet, 1 /*desiredReplicas*/, statefulSet, pod)
	if err != nil && phase != splcommon.PhaseUpdating {
		t.Errorf("UpdateStatefulSetPods should not have returned error=%s with phase=%s", err, phase)
	}

	// CurrentRevision = UpdateRevision
	statefulSet.Status.CurrentRevision = "v1"
	phase, err = updateStatefulSetPodsTester(t, &mgr, statefulSet, 1 /*desiredReplicas*/, statefulSet, pod)
	if err == nil && phase != splcommon.PhaseScalingUp {
		t.Errorf("UpdateStatefulSetPods should have returned error or phase should have been PhaseError, but we got phase=%s", phase)
	}

	// readyReplicas > replicas
	replicas = 2
	statefulSet.Status.ReadyReplicas = 3
	statefulSet.Spec.Replicas = &replicas
	statefulSet.Status.CurrentRevision = ""
	phase, err = updateStatefulSetPodsTester(t, &mgr, statefulSet, 1 /*desiredReplicas*/, statefulSet, pod)
	if err == nil && phase != splcommon.PhaseScalingDown {
		t.Errorf("UpdateStatefulSetPods should have returned error or phase should have been PhaseError, but we got phase=%s", phase)
	}

	// CurrentRevision = UpdateRevision
	statefulSet.Status.CurrentRevision = "v1"
	phase, err = updateStatefulSetPodsTester(t, &mgr, statefulSet, 1 /*desiredReplicas*/, statefulSet, pod)
	if err == nil && phase != splcommon.PhaseScalingDown {
		t.Errorf("UpdateStatefulSetPods should have returned error or phase should have been PhaseError, but we got phase=%s", phase)
	}
}

func TestSetStatefulSetOwnerRef(t *testing.T) {
	cr := enterpriseApi.Standalone{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "stack1",
			Namespace: "test",
		},
	}

	c := spltest.NewMockClient()
	current := appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "splunk-test-monitoring-console",
			Namespace: "test",
		},
	}
	namespacedName := types.NamespacedName{Namespace: "test", Name: "splunk-test-monitoring-console"}

	err := SetStatefulSetOwnerRef(c, &cr, namespacedName)
	if err.Error() != "NotFound" {
		t.Errorf("Couldn't detect resource %s", current.GetName())
	}

	// Create statefulset
	err = splutil.CreateResource(c, &current)
	if err != nil {
		t.Errorf("Failed to create owner reference  %s", current.GetName())
	}

	// Test existing owner reference
	err = SetStatefulSetOwnerRef(c, &cr, namespacedName)
	if err != nil {
		t.Errorf("Couldn't set owner ref for resource %s", current.GetName())
	}

	// Try adding same owner again
	err = SetStatefulSetOwnerRef(c, &cr, namespacedName)
	if err != nil {
		t.Errorf("Couldn't set owner ref for statefulset %s", current.GetName())
	}
}

func TestGetStatefulSetByName(t *testing.T) {
	c := spltest.NewMockClient()

	current := appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "splunk-test-monitoring-console",
			Namespace: "test",
		},
	}

	_, err := ApplyStatefulSet(c, &current)
	if err != nil {
		return
	}

	namespacedName := types.NamespacedName{Namespace: "test", Name: "splunk-test-monitoring-console"}
	_, err = GetStatefulSetByName(c, namespacedName)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestIsStatefulSetScalingUp(t *testing.T) {
	var replicas int32 = 1
	statefulSetName := "splunk-stand1-standalone"

	cr := enterpriseApi.Standalone{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "stand1",
			Namespace: "test",
		},
	}

	current := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      statefulSetName,
			Namespace: "test",
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: &replicas,
		},
	}

	c := spltest.NewMockClient()

	*current.Spec.Replicas = 2
	_, err := IsStatefulSetScalingUp(c, &cr, statefulSetName, replicas)
	if err == nil {
		t.Errorf("IsStatefulSetScalingUp should have returned error as we have not yet added statefulset to client.")
	}

	c.AddObject(current)
	_, err = IsStatefulSetScalingUp(c, &cr, statefulSetName, replicas)
	if err != nil {
		t.Errorf("IsStatefulSetScalingUp should not have returned error")
	}
}
