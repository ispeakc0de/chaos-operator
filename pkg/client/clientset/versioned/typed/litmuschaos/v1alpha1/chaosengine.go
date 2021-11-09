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
	"context"
	"time"

	v1alpha1 "github.com/litmuschaos/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
	scheme "github.com/litmuschaos/chaos-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ChaosEnginesGetter has a method to return a ChaosEngineInterface.
// A group's client should implement this interface.
type ChaosEnginesGetter interface {
	ChaosEngines(namespace string) ChaosEngineInterface
}

// ChaosEngineInterface has methods to work with ChaosEngine resources.
type ChaosEngineInterface interface {
	Create(*v1alpha1.ChaosEngine) (*v1alpha1.ChaosEngine, error)
	Update(*v1alpha1.ChaosEngine) (*v1alpha1.ChaosEngine, error)
	UpdateStatus(*v1alpha1.ChaosEngine) (*v1alpha1.ChaosEngine, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ChaosEngine, error)
	List(opts v1.ListOptions) (*v1alpha1.ChaosEngineList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ChaosEngine, err error)
	ChaosEngineExpansion
}

// chaosEngines implements ChaosEngineInterface
type chaosEngines struct {
	client rest.Interface
	ns     string
}

// newChaosEngines returns a ChaosEngines
func newChaosEngines(c *LitmuschaosV1alpha1Client, namespace string) *chaosEngines {
	return &chaosEngines{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the chaosEngine, and returns the corresponding chaosEngine object, and an error if there is any.
func (c *chaosEngines) Get(name string, options v1.GetOptions) (result *v1alpha1.ChaosEngine, err error) {
	result = &v1alpha1.ChaosEngine{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chaosengines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ChaosEngines that match those selectors.
func (c *chaosEngines) List(opts v1.ListOptions) (result *v1alpha1.ChaosEngineList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ChaosEngineList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chaosengines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.TODO()).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested chaosEngines.
func (c *chaosEngines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("chaosengines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(context.TODO())
}

// Create takes the representation of a chaosEngine and creates it.  Returns the server's representation of the chaosEngine, and an error, if there is any.
func (c *chaosEngines) Create(chaosEngine *v1alpha1.ChaosEngine) (result *v1alpha1.ChaosEngine, err error) {
	result = &v1alpha1.ChaosEngine{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("chaosengines").
		Body(chaosEngine).
		Do(context.TODO()).
		Into(result)
	return
}

// Update takes the representation of a chaosEngine and updates it. Returns the server's representation of the chaosEngine, and an error, if there is any.
func (c *chaosEngines) Update(chaosEngine *v1alpha1.ChaosEngine) (result *v1alpha1.ChaosEngine, err error) {
	result = &v1alpha1.ChaosEngine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("chaosengines").
		Name(chaosEngine.Name).
		Body(chaosEngine).
		Do(context.TODO()).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *chaosEngines) UpdateStatus(chaosEngine *v1alpha1.ChaosEngine) (result *v1alpha1.ChaosEngine, err error) {
	result = &v1alpha1.ChaosEngine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("chaosengines").
		Name(chaosEngine.Name).
		SubResource("status").
		Body(chaosEngine).
		Do(context.TODO()).
		Into(result)
	return
}

// Delete takes name of the chaosEngine and deletes it. Returns an error if one occurs.
func (c *chaosEngines) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chaosengines").
		Name(name).
		Body(options).
		Do(context.TODO()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *chaosEngines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chaosengines").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(context.TODO()).
		Error()
}

// Patch applies the patch and returns the patched chaosEngine.
func (c *chaosEngines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ChaosEngine, err error) {
	result = &v1alpha1.ChaosEngine{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("chaosengines").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.TODO()).
		Into(result)
	return
}
