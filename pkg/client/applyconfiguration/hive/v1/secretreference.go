// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// SecretReferenceApplyConfiguration represents an declarative configuration of the SecretReference type for use
// with apply.
type SecretReferenceApplyConfiguration struct {
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// SecretReferenceApplyConfiguration constructs an declarative configuration of the SecretReference type for use with
// apply.
func SecretReference() *SecretReferenceApplyConfiguration {
	return &SecretReferenceApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *SecretReferenceApplyConfiguration) WithName(value string) *SecretReferenceApplyConfiguration {
	b.Name = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *SecretReferenceApplyConfiguration) WithNamespace(value string) *SecretReferenceApplyConfiguration {
	b.Namespace = &value
	return b
}
