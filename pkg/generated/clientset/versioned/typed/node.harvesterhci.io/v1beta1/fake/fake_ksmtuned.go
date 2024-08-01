/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/harvester/node-manager/pkg/apis/node.harvesterhci.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKsmtuneds implements KsmtunedInterface
type FakeKsmtuneds struct {
	Fake *FakeNodeV1beta1
}

var ksmtunedsResource = v1beta1.SchemeGroupVersion.WithResource("ksmtuneds")

var ksmtunedsKind = v1beta1.SchemeGroupVersion.WithKind("Ksmtuned")

// Get takes name of the ksmtuned, and returns the corresponding ksmtuned object, and an error if there is any.
func (c *FakeKsmtuneds) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Ksmtuned, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ksmtunedsResource, name), &v1beta1.Ksmtuned{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ksmtuned), err
}

// List takes label and field selectors, and returns the list of Ksmtuneds that match those selectors.
func (c *FakeKsmtuneds) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.KsmtunedList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ksmtunedsResource, ksmtunedsKind, opts), &v1beta1.KsmtunedList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.KsmtunedList{ListMeta: obj.(*v1beta1.KsmtunedList).ListMeta}
	for _, item := range obj.(*v1beta1.KsmtunedList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ksmtuneds.
func (c *FakeKsmtuneds) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ksmtunedsResource, opts))
}

// Create takes the representation of a ksmtuned and creates it.  Returns the server's representation of the ksmtuned, and an error, if there is any.
func (c *FakeKsmtuneds) Create(ctx context.Context, ksmtuned *v1beta1.Ksmtuned, opts v1.CreateOptions) (result *v1beta1.Ksmtuned, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ksmtunedsResource, ksmtuned), &v1beta1.Ksmtuned{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ksmtuned), err
}

// Update takes the representation of a ksmtuned and updates it. Returns the server's representation of the ksmtuned, and an error, if there is any.
func (c *FakeKsmtuneds) Update(ctx context.Context, ksmtuned *v1beta1.Ksmtuned, opts v1.UpdateOptions) (result *v1beta1.Ksmtuned, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ksmtunedsResource, ksmtuned), &v1beta1.Ksmtuned{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ksmtuned), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKsmtuneds) UpdateStatus(ctx context.Context, ksmtuned *v1beta1.Ksmtuned, opts v1.UpdateOptions) (*v1beta1.Ksmtuned, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(ksmtunedsResource, "status", ksmtuned), &v1beta1.Ksmtuned{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ksmtuned), err
}

// Delete takes name of the ksmtuned and deletes it. Returns an error if one occurs.
func (c *FakeKsmtuneds) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ksmtunedsResource, name, opts), &v1beta1.Ksmtuned{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKsmtuneds) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ksmtunedsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.KsmtunedList{})
	return err
}

// Patch applies the patch and returns the patched ksmtuned.
func (c *FakeKsmtuneds) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Ksmtuned, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ksmtunedsResource, name, pt, data, subresources...), &v1beta1.Ksmtuned{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ksmtuned), err
}
