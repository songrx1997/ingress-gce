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
	serviceattachmentv1 "k8s.io/ingress-gce/pkg/apis/serviceattachment/v1"
)

// FakeServiceAttachments implements ServiceAttachmentInterface
type FakeServiceAttachments struct {
	Fake *FakeNetworkingV1
	ns   string
}

var serviceattachmentsResource = schema.GroupVersionResource{Group: "networking.gke.io", Version: "v1", Resource: "serviceattachments"}

var serviceattachmentsKind = schema.GroupVersionKind{Group: "networking.gke.io", Version: "v1", Kind: "ServiceAttachment"}

// Get takes name of the serviceAttachment, and returns the corresponding serviceAttachment object, and an error if there is any.
func (c *FakeServiceAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *serviceattachmentv1.ServiceAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serviceattachmentsResource, c.ns, name), &serviceattachmentv1.ServiceAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*serviceattachmentv1.ServiceAttachment), err
}

// List takes label and field selectors, and returns the list of ServiceAttachments that match those selectors.
func (c *FakeServiceAttachments) List(ctx context.Context, opts v1.ListOptions) (result *serviceattachmentv1.ServiceAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serviceattachmentsResource, serviceattachmentsKind, c.ns, opts), &serviceattachmentv1.ServiceAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &serviceattachmentv1.ServiceAttachmentList{ListMeta: obj.(*serviceattachmentv1.ServiceAttachmentList).ListMeta}
	for _, item := range obj.(*serviceattachmentv1.ServiceAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceAttachments.
func (c *FakeServiceAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serviceattachmentsResource, c.ns, opts))

}

// Create takes the representation of a serviceAttachment and creates it.  Returns the server's representation of the serviceAttachment, and an error, if there is any.
func (c *FakeServiceAttachments) Create(ctx context.Context, serviceAttachment *serviceattachmentv1.ServiceAttachment, opts v1.CreateOptions) (result *serviceattachmentv1.ServiceAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serviceattachmentsResource, c.ns, serviceAttachment), &serviceattachmentv1.ServiceAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*serviceattachmentv1.ServiceAttachment), err
}

// Update takes the representation of a serviceAttachment and updates it. Returns the server's representation of the serviceAttachment, and an error, if there is any.
func (c *FakeServiceAttachments) Update(ctx context.Context, serviceAttachment *serviceattachmentv1.ServiceAttachment, opts v1.UpdateOptions) (result *serviceattachmentv1.ServiceAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serviceattachmentsResource, c.ns, serviceAttachment), &serviceattachmentv1.ServiceAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*serviceattachmentv1.ServiceAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceAttachments) UpdateStatus(ctx context.Context, serviceAttachment *serviceattachmentv1.ServiceAttachment, opts v1.UpdateOptions) (*serviceattachmentv1.ServiceAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serviceattachmentsResource, "status", c.ns, serviceAttachment), &serviceattachmentv1.ServiceAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*serviceattachmentv1.ServiceAttachment), err
}

// Delete takes name of the serviceAttachment and deletes it. Returns an error if one occurs.
func (c *FakeServiceAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serviceattachmentsResource, c.ns, name), &serviceattachmentv1.ServiceAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serviceattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &serviceattachmentv1.ServiceAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched serviceAttachment.
func (c *FakeServiceAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *serviceattachmentv1.ServiceAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serviceattachmentsResource, c.ns, name, pt, data, subresources...), &serviceattachmentv1.ServiceAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*serviceattachmentv1.ServiceAttachment), err
}
