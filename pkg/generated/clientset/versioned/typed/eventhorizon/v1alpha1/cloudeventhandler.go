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

package v1alpha1

import (
	"time"

	v1alpha1 "acesso.io/eventhorizon/pkg/apis/eventhorizon/v1alpha1"
	scheme "acesso.io/eventhorizon/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudEventHandlersGetter has a method to return a CloudEventHandlerInterface.
// A group's client should implement this interface.
type CloudEventHandlersGetter interface {
	CloudEventHandlers() CloudEventHandlerInterface
}

// CloudEventHandlerInterface has methods to work with CloudEventHandler resources.
type CloudEventHandlerInterface interface {
	Create(*v1alpha1.CloudEventHandler) (*v1alpha1.CloudEventHandler, error)
	Update(*v1alpha1.CloudEventHandler) (*v1alpha1.CloudEventHandler, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudEventHandler, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudEventHandlerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudEventHandler, err error)
	CloudEventHandlerExpansion
}

// cloudEventHandlers implements CloudEventHandlerInterface
type cloudEventHandlers struct {
	client rest.Interface
}

// newCloudEventHandlers returns a CloudEventHandlers
func newCloudEventHandlers(c *EventhorizonV1alpha1Client) *cloudEventHandlers {
	return &cloudEventHandlers{
		client: c.RESTClient(),
	}
}

// Get takes name of the cloudEventHandler, and returns the corresponding cloudEventHandler object, and an error if there is any.
func (c *cloudEventHandlers) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudEventHandler, err error) {
	result = &v1alpha1.CloudEventHandler{}
	err = c.client.Get().
		Resource("cloudeventhandlers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudEventHandlers that match those selectors.
func (c *cloudEventHandlers) List(opts v1.ListOptions) (result *v1alpha1.CloudEventHandlerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudEventHandlerList{}
	err = c.client.Get().
		Resource("cloudeventhandlers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudEventHandlers.
func (c *cloudEventHandlers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("cloudeventhandlers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudEventHandler and creates it.  Returns the server's representation of the cloudEventHandler, and an error, if there is any.
func (c *cloudEventHandlers) Create(cloudEventHandler *v1alpha1.CloudEventHandler) (result *v1alpha1.CloudEventHandler, err error) {
	result = &v1alpha1.CloudEventHandler{}
	err = c.client.Post().
		Resource("cloudeventhandlers").
		Body(cloudEventHandler).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudEventHandler and updates it. Returns the server's representation of the cloudEventHandler, and an error, if there is any.
func (c *cloudEventHandlers) Update(cloudEventHandler *v1alpha1.CloudEventHandler) (result *v1alpha1.CloudEventHandler, err error) {
	result = &v1alpha1.CloudEventHandler{}
	err = c.client.Put().
		Resource("cloudeventhandlers").
		Name(cloudEventHandler.Name).
		Body(cloudEventHandler).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudEventHandler and deletes it. Returns an error if one occurs.
func (c *cloudEventHandlers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("cloudeventhandlers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudEventHandlers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("cloudeventhandlers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudEventHandler.
func (c *cloudEventHandlers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudEventHandler, err error) {
	result = &v1alpha1.CloudEventHandler{}
	err = c.client.Patch(pt).
		Resource("cloudeventhandlers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}