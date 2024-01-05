/*
Copyright 2024 The Kubernetes Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	backendconfigv1 "k8s.io/ingress-gce/pkg/apis/backendconfig/v1"
)

// FakeBackendConfigs implements BackendConfigInterface
type FakeBackendConfigs struct {
	Fake *FakeCloudV1
	ns   string
}

var backendconfigsResource = schema.GroupVersionResource{Group: "cloud.google.com", Version: "v1", Resource: "backendconfigs"}

var backendconfigsKind = schema.GroupVersionKind{Group: "cloud.google.com", Version: "v1", Kind: "BackendConfig"}

// Get takes name of the backendConfig, and returns the corresponding backendConfig object, and an error if there is any.
func (c *FakeBackendConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *backendconfigv1.BackendConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backendconfigsResource, c.ns, name), &backendconfigv1.BackendConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backendconfigv1.BackendConfig), err
}

// List takes label and field selectors, and returns the list of BackendConfigs that match those selectors.
func (c *FakeBackendConfigs) List(ctx context.Context, opts v1.ListOptions) (result *backendconfigv1.BackendConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backendconfigsResource, backendconfigsKind, c.ns, opts), &backendconfigv1.BackendConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &backendconfigv1.BackendConfigList{ListMeta: obj.(*backendconfigv1.BackendConfigList).ListMeta}
	for _, item := range obj.(*backendconfigv1.BackendConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backendConfigs.
func (c *FakeBackendConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backendconfigsResource, c.ns, opts))

}

// Create takes the representation of a backendConfig and creates it.  Returns the server's representation of the backendConfig, and an error, if there is any.
func (c *FakeBackendConfigs) Create(ctx context.Context, backendConfig *backendconfigv1.BackendConfig, opts v1.CreateOptions) (result *backendconfigv1.BackendConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backendconfigsResource, c.ns, backendConfig), &backendconfigv1.BackendConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backendconfigv1.BackendConfig), err
}

// Update takes the representation of a backendConfig and updates it. Returns the server's representation of the backendConfig, and an error, if there is any.
func (c *FakeBackendConfigs) Update(ctx context.Context, backendConfig *backendconfigv1.BackendConfig, opts v1.UpdateOptions) (result *backendconfigv1.BackendConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backendconfigsResource, c.ns, backendConfig), &backendconfigv1.BackendConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backendconfigv1.BackendConfig), err
}

// Delete takes name of the backendConfig and deletes it. Returns an error if one occurs.
func (c *FakeBackendConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backendconfigsResource, c.ns, name), &backendconfigv1.BackendConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackendConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backendconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &backendconfigv1.BackendConfigList{})
	return err
}

// Patch applies the patch and returns the patched backendConfig.
func (c *FakeBackendConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *backendconfigv1.BackendConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backendconfigsResource, c.ns, name, pt, data, subresources...), &backendconfigv1.BackendConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backendconfigv1.BackendConfig), err
}
