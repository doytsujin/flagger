/*
Copyright 2020 The Flux authors

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

	v1beta1 "github.com/fluxcd/flagger/pkg/apis/istio/v1beta1"
	scheme "github.com/fluxcd/flagger/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// VirtualServicesGetter has a method to return a VirtualServiceInterface.
// A group's client should implement this interface.
type VirtualServicesGetter interface {
	VirtualServices(namespace string) VirtualServiceInterface
}

// VirtualServiceInterface has methods to work with VirtualService resources.
type VirtualServiceInterface interface {
	Create(ctx context.Context, virtualService *v1beta1.VirtualService, opts v1.CreateOptions) (*v1beta1.VirtualService, error)
	Update(ctx context.Context, virtualService *v1beta1.VirtualService, opts v1.UpdateOptions) (*v1beta1.VirtualService, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VirtualService, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VirtualServiceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualService, err error)
	VirtualServiceExpansion
}

// virtualServices implements VirtualServiceInterface
type virtualServices struct {
	*gentype.ClientWithList[*v1beta1.VirtualService, *v1beta1.VirtualServiceList]
}

// newVirtualServices returns a VirtualServices
func newVirtualServices(c *NetworkingV1beta1Client, namespace string) *virtualServices {
	return &virtualServices{
		gentype.NewClientWithList[*v1beta1.VirtualService, *v1beta1.VirtualServiceList](
			"virtualservices",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.VirtualService { return &v1beta1.VirtualService{} },
			func() *v1beta1.VirtualServiceList { return &v1beta1.VirtualServiceList{} }),
	}
}
