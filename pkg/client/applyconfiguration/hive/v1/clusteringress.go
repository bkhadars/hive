// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	configv1 "github.com/openshift/api/config/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterIngressApplyConfiguration represents an declarative configuration of the ClusterIngress type for use
// with apply.
type ClusterIngressApplyConfiguration struct {
	Name               *string                          `json:"name,omitempty"`
	Domain             *string                          `json:"domain,omitempty"`
	NamespaceSelector  *v1.LabelSelector                `json:"namespaceSelector,omitempty"`
	RouteSelector      *v1.LabelSelector                `json:"routeSelector,omitempty"`
	ServingCertificate *string                          `json:"servingCertificate,omitempty"`
	HttpErrorCodePages *configv1.ConfigMapNameReference `json:"httpErrorCodePages,omitempty"`
}

// ClusterIngressApplyConfiguration constructs an declarative configuration of the ClusterIngress type for use with
// apply.
func ClusterIngress() *ClusterIngressApplyConfiguration {
	return &ClusterIngressApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ClusterIngressApplyConfiguration) WithName(value string) *ClusterIngressApplyConfiguration {
	b.Name = &value
	return b
}

// WithDomain sets the Domain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Domain field is set to the value of the last call.
func (b *ClusterIngressApplyConfiguration) WithDomain(value string) *ClusterIngressApplyConfiguration {
	b.Domain = &value
	return b
}

// WithNamespaceSelector sets the NamespaceSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NamespaceSelector field is set to the value of the last call.
func (b *ClusterIngressApplyConfiguration) WithNamespaceSelector(value v1.LabelSelector) *ClusterIngressApplyConfiguration {
	b.NamespaceSelector = &value
	return b
}

// WithRouteSelector sets the RouteSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RouteSelector field is set to the value of the last call.
func (b *ClusterIngressApplyConfiguration) WithRouteSelector(value v1.LabelSelector) *ClusterIngressApplyConfiguration {
	b.RouteSelector = &value
	return b
}

// WithServingCertificate sets the ServingCertificate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServingCertificate field is set to the value of the last call.
func (b *ClusterIngressApplyConfiguration) WithServingCertificate(value string) *ClusterIngressApplyConfiguration {
	b.ServingCertificate = &value
	return b
}

// WithHttpErrorCodePages sets the HttpErrorCodePages field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HttpErrorCodePages field is set to the value of the last call.
func (b *ClusterIngressApplyConfiguration) WithHttpErrorCodePages(value configv1.ConfigMapNameReference) *ClusterIngressApplyConfiguration {
	b.HttpErrorCodePages = &value
	return b
}
