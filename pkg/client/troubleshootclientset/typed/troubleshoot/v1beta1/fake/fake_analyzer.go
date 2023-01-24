/*
Copyright 2019 Replicated, Inc..

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

	v1beta1 "github.com/replicatedhq/troubleshoot/pkg/apis/troubleshoot/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAnalyzers implements AnalyzerInterface
type FakeAnalyzers struct {
	Fake *FakeTroubleshootV1beta1
	ns   string
}

var analyzersResource = schema.GroupVersionResource{Group: "troubleshoot.replicated.com", Version: "v1beta1", Resource: "analyzers"}

var analyzersKind = schema.GroupVersionKind{Group: "troubleshoot.replicated.com", Version: "v1beta1", Kind: "Analyzer"}

// Get takes name of the analyzer, and returns the corresponding analyzer object, and an error if there is any.
func (c *FakeAnalyzers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Analyzer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(analyzersResource, c.ns, name), &v1beta1.Analyzer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Analyzer), err
}

// List takes label and field selectors, and returns the list of Analyzers that match those selectors.
func (c *FakeAnalyzers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.AnalyzerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(analyzersResource, analyzersKind, c.ns, opts), &v1beta1.AnalyzerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.AnalyzerList{ListMeta: obj.(*v1beta1.AnalyzerList).ListMeta}
	for _, item := range obj.(*v1beta1.AnalyzerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested analyzers.
func (c *FakeAnalyzers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(analyzersResource, c.ns, opts))

}

// Create takes the representation of a analyzer and creates it.  Returns the server's representation of the analyzer, and an error, if there is any.
func (c *FakeAnalyzers) Create(ctx context.Context, analyzer *v1beta1.Analyzer, opts v1.CreateOptions) (result *v1beta1.Analyzer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(analyzersResource, c.ns, analyzer), &v1beta1.Analyzer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Analyzer), err
}

// Update takes the representation of a analyzer and updates it. Returns the server's representation of the analyzer, and an error, if there is any.
func (c *FakeAnalyzers) Update(ctx context.Context, analyzer *v1beta1.Analyzer, opts v1.UpdateOptions) (result *v1beta1.Analyzer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(analyzersResource, c.ns, analyzer), &v1beta1.Analyzer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Analyzer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAnalyzers) UpdateStatus(ctx context.Context, analyzer *v1beta1.Analyzer, opts v1.UpdateOptions) (*v1beta1.Analyzer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(analyzersResource, "status", c.ns, analyzer), &v1beta1.Analyzer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Analyzer), err
}

// Delete takes name of the analyzer and deletes it. Returns an error if one occurs.
func (c *FakeAnalyzers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(analyzersResource, c.ns, name, opts), &v1beta1.Analyzer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAnalyzers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(analyzersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.AnalyzerList{})
	return err
}

// Patch applies the patch and returns the patched analyzer.
func (c *FakeAnalyzers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Analyzer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(analyzersResource, c.ns, name, pt, data, subresources...), &v1beta1.Analyzer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Analyzer), err
}
