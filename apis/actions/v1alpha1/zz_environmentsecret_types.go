/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EnvironmentSecretObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type EnvironmentSecretParameters struct {

	// +kubebuilder:validation:Optional
	EncryptedValueSecretRef *v1.SecretKeySelector `json:"encryptedValueSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Environment *string `json:"environment" tf:"environment,omitempty"`

	// +kubebuilder:validation:Optional
	PlaintextValueSecretRef *v1.SecretKeySelector `json:"plaintextValueSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Repository *string `json:"repository" tf:"repository,omitempty"`

	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

// EnvironmentSecretSpec defines the desired state of EnvironmentSecret
type EnvironmentSecretSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentSecretParameters `json:"forProvider"`
}

// EnvironmentSecretStatus defines the observed state of EnvironmentSecret.
type EnvironmentSecretStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentSecretObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentSecret is the Schema for the EnvironmentSecrets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,githubprovjet}
type EnvironmentSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EnvironmentSecretSpec   `json:"spec"`
	Status            EnvironmentSecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentSecretList contains a list of EnvironmentSecrets
type EnvironmentSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnvironmentSecret `json:"items"`
}

// Repository type metadata.
var (
	EnvironmentSecret_Kind             = "EnvironmentSecret"
	EnvironmentSecret_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EnvironmentSecret_Kind}.String()
	EnvironmentSecret_KindAPIVersion   = EnvironmentSecret_Kind + "." + CRDGroupVersion.String()
	EnvironmentSecret_GroupVersionKind = CRDGroupVersion.WithKind(EnvironmentSecret_Kind)
)

func init() {
	SchemeBuilder.Register(&EnvironmentSecret{}, &EnvironmentSecretList{})
}
