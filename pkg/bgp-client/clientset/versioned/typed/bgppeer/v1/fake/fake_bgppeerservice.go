/*
Copyright The Kubernetes Authors.

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
	"context"
	json "encoding/json"
	"fmt"

	bgppeerv1 "github.com/loxilb-io/kube-loxilb/pkg/bgp-client/applyconfiguration/bgppeer/v1"
	v1 "github.com/loxilb-io/kube-loxilb/pkg/crds/bgppeer/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBGPPeerServices implements BGPPeerServiceInterface
type FakeBGPPeerServices struct {
	Fake *FakeBgppeerV1
}

var bgppeerservicesResource = v1.SchemeGroupVersion.WithResource("bgppeerservices")

var bgppeerservicesKind = v1.SchemeGroupVersion.WithKind("BGPPeerService")

// Get takes name of the bGPPeerService, and returns the corresponding bGPPeerService object, and an error if there is any.
func (c *FakeBGPPeerServices) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.BGPPeerService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(bgppeerservicesResource, name), &v1.BGPPeerService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.BGPPeerService), err
}

// List takes label and field selectors, and returns the list of BGPPeerServices that match those selectors.
func (c *FakeBGPPeerServices) List(ctx context.Context, opts metav1.ListOptions) (result *v1.BGPPeerServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(bgppeerservicesResource, bgppeerservicesKind, opts), &v1.BGPPeerServiceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.BGPPeerServiceList{ListMeta: obj.(*v1.BGPPeerServiceList).ListMeta}
	for _, item := range obj.(*v1.BGPPeerServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bGPPeerServices.
func (c *FakeBGPPeerServices) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(bgppeerservicesResource, opts))
}

// Create takes the representation of a bGPPeerService and creates it.  Returns the server's representation of the bGPPeerService, and an error, if there is any.
func (c *FakeBGPPeerServices) Create(ctx context.Context, bGPPeerService *v1.BGPPeerService, opts metav1.CreateOptions) (result *v1.BGPPeerService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(bgppeerservicesResource, bGPPeerService), &v1.BGPPeerService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.BGPPeerService), err
}

// Update takes the representation of a bGPPeerService and updates it. Returns the server's representation of the bGPPeerService, and an error, if there is any.
func (c *FakeBGPPeerServices) Update(ctx context.Context, bGPPeerService *v1.BGPPeerService, opts metav1.UpdateOptions) (result *v1.BGPPeerService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(bgppeerservicesResource, bGPPeerService), &v1.BGPPeerService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.BGPPeerService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBGPPeerServices) UpdateStatus(ctx context.Context, bGPPeerService *v1.BGPPeerService, opts metav1.UpdateOptions) (*v1.BGPPeerService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(bgppeerservicesResource, "status", bGPPeerService), &v1.BGPPeerService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.BGPPeerService), err
}

// Delete takes name of the bGPPeerService and deletes it. Returns an error if one occurs.
func (c *FakeBGPPeerServices) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(bgppeerservicesResource, name, opts), &v1.BGPPeerService{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBGPPeerServices) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(bgppeerservicesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.BGPPeerServiceList{})
	return err
}

// Patch applies the patch and returns the patched bGPPeerService.
func (c *FakeBGPPeerServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.BGPPeerService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(bgppeerservicesResource, name, pt, data, subresources...), &v1.BGPPeerService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.BGPPeerService), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied bGPPeerService.
func (c *FakeBGPPeerServices) Apply(ctx context.Context, bGPPeerService *bgppeerv1.BGPPeerServiceApplyConfiguration, opts metav1.ApplyOptions) (result *v1.BGPPeerService, err error) {
	if bGPPeerService == nil {
		return nil, fmt.Errorf("bGPPeerService provided to Apply must not be nil")
	}
	data, err := json.Marshal(bGPPeerService)
	if err != nil {
		return nil, err
	}
	name := bGPPeerService.Name
	if name == nil {
		return nil, fmt.Errorf("bGPPeerService.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(bgppeerservicesResource, *name, types.ApplyPatchType, data), &v1.BGPPeerService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.BGPPeerService), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeBGPPeerServices) ApplyStatus(ctx context.Context, bGPPeerService *bgppeerv1.BGPPeerServiceApplyConfiguration, opts metav1.ApplyOptions) (result *v1.BGPPeerService, err error) {
	if bGPPeerService == nil {
		return nil, fmt.Errorf("bGPPeerService provided to Apply must not be nil")
	}
	data, err := json.Marshal(bGPPeerService)
	if err != nil {
		return nil, err
	}
	name := bGPPeerService.Name
	if name == nil {
		return nil, fmt.Errorf("bGPPeerService.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(bgppeerservicesResource, *name, types.ApplyPatchType, data, "status"), &v1.BGPPeerService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.BGPPeerService), err
}