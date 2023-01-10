/*
Copyright The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"

	clientset "github.com/kcp-dev/kcp/pkg/client/clientset/versioned"
	apiresourcev1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/apiresource/v1alpha1"
	fakeapiresourcev1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/apiresource/v1alpha1/fake"
	apisv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/apis/v1alpha1"
	fakeapisv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/apis/v1alpha1/fake"
	corev1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/core/v1alpha1"
	fakecorev1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/core/v1alpha1/fake"
	schedulingv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/scheduling/v1alpha1"
	fakeschedulingv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/scheduling/v1alpha1/fake"
	tenancyv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/tenancy/v1alpha1"
	faketenancyv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/tenancy/v1alpha1/fake"
	tenancyv1beta1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/tenancy/v1beta1"
	faketenancyv1beta1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/tenancy/v1beta1/fake"
	topologyv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/topology/v1alpha1"
	faketopologyv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/topology/v1alpha1/fake"
	workloadv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/workload/v1alpha1"
	fakeworkloadv1alpha1 "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/workload/v1alpha1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// ApiresourceV1alpha1 retrieves the ApiresourceV1alpha1Client
func (c *Clientset) ApiresourceV1alpha1() apiresourcev1alpha1.ApiresourceV1alpha1Interface {
	return &fakeapiresourcev1alpha1.FakeApiresourceV1alpha1{Fake: &c.Fake}
}

// ApisV1alpha1 retrieves the ApisV1alpha1Client
func (c *Clientset) ApisV1alpha1() apisv1alpha1.ApisV1alpha1Interface {
	return &fakeapisv1alpha1.FakeApisV1alpha1{Fake: &c.Fake}
}

// CoreV1alpha1 retrieves the CoreV1alpha1Client
func (c *Clientset) CoreV1alpha1() corev1alpha1.CoreV1alpha1Interface {
	return &fakecorev1alpha1.FakeCoreV1alpha1{Fake: &c.Fake}
}

// SchedulingV1alpha1 retrieves the SchedulingV1alpha1Client
func (c *Clientset) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return &fakeschedulingv1alpha1.FakeSchedulingV1alpha1{Fake: &c.Fake}
}

// TenancyV1alpha1 retrieves the TenancyV1alpha1Client
func (c *Clientset) TenancyV1alpha1() tenancyv1alpha1.TenancyV1alpha1Interface {
	return &faketenancyv1alpha1.FakeTenancyV1alpha1{Fake: &c.Fake}
}

// TenancyV1beta1 retrieves the TenancyV1beta1Client
func (c *Clientset) TenancyV1beta1() tenancyv1beta1.TenancyV1beta1Interface {
	return &faketenancyv1beta1.FakeTenancyV1beta1{Fake: &c.Fake}
}

// TopologyV1alpha1 retrieves the TopologyV1alpha1Client
func (c *Clientset) TopologyV1alpha1() topologyv1alpha1.TopologyV1alpha1Interface {
	return &faketopologyv1alpha1.FakeTopologyV1alpha1{Fake: &c.Fake}
}

// WorkloadV1alpha1 retrieves the WorkloadV1alpha1Client
func (c *Clientset) WorkloadV1alpha1() workloadv1alpha1.WorkloadV1alpha1Interface {
	return &fakeworkloadv1alpha1.FakeWorkloadV1alpha1{Fake: &c.Fake}
}
