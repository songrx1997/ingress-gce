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

package v1beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1"
	scheme "k8s.io/ingress-gce/pkg/backendconfig/client/clientset/versioned/scheme"
)

// BackendConfigsGetter has a method to return a BackendConfigInterface.
// A group's client should implement this interface.
type BackendConfigsGetter interface {
	BackendConfigs(namespace string) BackendConfigInterface
}

// BackendConfigInterface has methods to work with BackendConfig resources.
type BackendConfigInterface interface {
	Create(ctx context.Context, backendConfig *v1beta1.BackendConfig, opts v1.CreateOptions) (*v1beta1.BackendConfig, error)
	Update(ctx context.Context, backendConfig *v1beta1.BackendConfig, opts v1.UpdateOptions) (*v1beta1.BackendConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.BackendConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.BackendConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackendConfig, err error)
	BackendConfigExpansion
}

// backendConfigs implements BackendConfigInterface
type backendConfigs struct {
	client rest.Interface
	ns     string
}

// newBackendConfigs returns a BackendConfigs
func newBackendConfigs(c *CloudV1beta1Client, namespace string) *backendConfigs {
	return &backendConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backendConfig, and returns the corresponding backendConfig object, and an error if there is any.
func (c *backendConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BackendConfig, err error) {
	result = &v1beta1.BackendConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backendconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackendConfigs that match those selectors.
func (c *backendConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackendConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.BackendConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backendconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backendConfigs.
func (c *backendConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backendconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a backendConfig and creates it.  Returns the server's representation of the backendConfig, and an error, if there is any.
func (c *backendConfigs) Create(ctx context.Context, backendConfig *v1beta1.BackendConfig, opts v1.CreateOptions) (result *v1beta1.BackendConfig, err error) {
	result = &v1beta1.BackendConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backendconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backendConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a backendConfig and updates it. Returns the server's representation of the backendConfig, and an error, if there is any.
func (c *backendConfigs) Update(ctx context.Context, backendConfig *v1beta1.BackendConfig, opts v1.UpdateOptions) (result *v1beta1.BackendConfig, err error) {
	result = &v1beta1.BackendConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backendconfigs").
		Name(backendConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backendConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the backendConfig and deletes it. Returns an error if one occurs.
func (c *backendConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backendconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backendConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backendconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched backendConfig.
func (c *backendConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackendConfig, err error) {
	result = &v1beta1.BackendConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backendconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
