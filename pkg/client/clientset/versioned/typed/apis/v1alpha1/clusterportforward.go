/*
Copyright 2023 The Kubernetes Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "sigs.k8s.io/kwok/pkg/apis/v1alpha1"
	scheme "sigs.k8s.io/kwok/pkg/client/clientset/versioned/scheme"
)

// ClusterPortForwardsGetter has a method to return a ClusterPortForwardInterface.
// A group's client should implement this interface.
type ClusterPortForwardsGetter interface {
	ClusterPortForwards() ClusterPortForwardInterface
}

// ClusterPortForwardInterface has methods to work with ClusterPortForward resources.
type ClusterPortForwardInterface interface {
	Create(ctx context.Context, clusterPortForward *v1alpha1.ClusterPortForward, opts v1.CreateOptions) (*v1alpha1.ClusterPortForward, error)
	Update(ctx context.Context, clusterPortForward *v1alpha1.ClusterPortForward, opts v1.UpdateOptions) (*v1alpha1.ClusterPortForward, error)
	UpdateStatus(ctx context.Context, clusterPortForward *v1alpha1.ClusterPortForward, opts v1.UpdateOptions) (*v1alpha1.ClusterPortForward, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterPortForward, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterPortForwardList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterPortForward, err error)
	ClusterPortForwardExpansion
}

// clusterPortForwards implements ClusterPortForwardInterface
type clusterPortForwards struct {
	client rest.Interface
}

// newClusterPortForwards returns a ClusterPortForwards
func newClusterPortForwards(c *KwokV1alpha1Client) *clusterPortForwards {
	return &clusterPortForwards{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterPortForward, and returns the corresponding clusterPortForward object, and an error if there is any.
func (c *clusterPortForwards) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterPortForward, err error) {
	result = &v1alpha1.ClusterPortForward{}
	err = c.client.Get().
		Resource("clusterportforwards").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterPortForwards that match those selectors.
func (c *clusterPortForwards) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterPortForwardList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterPortForwardList{}
	err = c.client.Get().
		Resource("clusterportforwards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterPortForwards.
func (c *clusterPortForwards) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterportforwards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterPortForward and creates it.  Returns the server's representation of the clusterPortForward, and an error, if there is any.
func (c *clusterPortForwards) Create(ctx context.Context, clusterPortForward *v1alpha1.ClusterPortForward, opts v1.CreateOptions) (result *v1alpha1.ClusterPortForward, err error) {
	result = &v1alpha1.ClusterPortForward{}
	err = c.client.Post().
		Resource("clusterportforwards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterPortForward).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterPortForward and updates it. Returns the server's representation of the clusterPortForward, and an error, if there is any.
func (c *clusterPortForwards) Update(ctx context.Context, clusterPortForward *v1alpha1.ClusterPortForward, opts v1.UpdateOptions) (result *v1alpha1.ClusterPortForward, err error) {
	result = &v1alpha1.ClusterPortForward{}
	err = c.client.Put().
		Resource("clusterportforwards").
		Name(clusterPortForward.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterPortForward).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterPortForwards) UpdateStatus(ctx context.Context, clusterPortForward *v1alpha1.ClusterPortForward, opts v1.UpdateOptions) (result *v1alpha1.ClusterPortForward, err error) {
	result = &v1alpha1.ClusterPortForward{}
	err = c.client.Put().
		Resource("clusterportforwards").
		Name(clusterPortForward.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterPortForward).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterPortForward and deletes it. Returns an error if one occurs.
func (c *clusterPortForwards) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterportforwards").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterPortForwards) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterportforwards").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterPortForward.
func (c *clusterPortForwards) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterPortForward, err error) {
	result = &v1alpha1.ClusterPortForward{}
	err = c.client.Patch(pt).
		Resource("clusterportforwards").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
