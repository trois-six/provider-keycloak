// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type KeystoreEcdsaGeneratedInitParameters struct {

	// When false, key in not used for signing. Defaults to true.
	// Set if the keys can be used for signing
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Elliptic Curve used in ECDSA. Defaults to P-256.
	// Elliptic Curve used in ECDSA
	EllipticCurveKey *string `json:"ellipticCurveKey,omitempty" tf:"elliptic_curve_key,omitempty"`

	// When false, key is not accessible in this realm. Defaults to true.
	// Set if the keys are enabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Display name of provider when linked in admin console.
	// Display name of provider when linked in admin console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Priority for the provider. Defaults to 0
	// Priority for the provider
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type KeystoreEcdsaGeneratedObservation struct {

	// When false, key in not used for signing. Defaults to true.
	// Set if the keys can be used for signing
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Elliptic Curve used in ECDSA. Defaults to P-256.
	// Elliptic Curve used in ECDSA
	EllipticCurveKey *string `json:"ellipticCurveKey,omitempty" tf:"elliptic_curve_key,omitempty"`

	// When false, key is not accessible in this realm. Defaults to true.
	// Set if the keys are enabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Display name of provider when linked in admin console.
	// Display name of provider when linked in admin console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Priority for the provider. Defaults to 0
	// Priority for the provider
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The realm this keystore exists in.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`
}

type KeystoreEcdsaGeneratedParameters struct {

	// When false, key in not used for signing. Defaults to true.
	// Set if the keys can be used for signing
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Elliptic Curve used in ECDSA. Defaults to P-256.
	// Elliptic Curve used in ECDSA
	// +kubebuilder:validation:Optional
	EllipticCurveKey *string `json:"ellipticCurveKey,omitempty" tf:"elliptic_curve_key,omitempty"`

	// When false, key is not accessible in this realm. Defaults to true.
	// Set if the keys are enabled
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Display name of provider when linked in admin console.
	// Display name of provider when linked in admin console.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Priority for the provider. Defaults to 0
	// Priority for the provider
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The realm this keystore exists in.
	// +crossplane:generate:reference:type=github.com/trois-six/provider-keycloak/apis/keycloak/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in keycloak to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in keycloak to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

// KeystoreEcdsaGeneratedSpec defines the desired state of KeystoreEcdsaGenerated
type KeystoreEcdsaGeneratedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeystoreEcdsaGeneratedParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider KeystoreEcdsaGeneratedInitParameters `json:"initProvider,omitempty"`
}

// KeystoreEcdsaGeneratedStatus defines the observed state of KeystoreEcdsaGenerated.
type KeystoreEcdsaGeneratedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeystoreEcdsaGeneratedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// KeystoreEcdsaGenerated is the Schema for the KeystoreEcdsaGenerateds API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type KeystoreEcdsaGenerated struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   KeystoreEcdsaGeneratedSpec   `json:"spec"`
	Status KeystoreEcdsaGeneratedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeystoreEcdsaGeneratedList contains a list of KeystoreEcdsaGenerateds
type KeystoreEcdsaGeneratedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeystoreEcdsaGenerated `json:"items"`
}

// Repository type metadata.
var (
	KeystoreEcdsaGenerated_Kind             = "KeystoreEcdsaGenerated"
	KeystoreEcdsaGenerated_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KeystoreEcdsaGenerated_Kind}.String()
	KeystoreEcdsaGenerated_KindAPIVersion   = KeystoreEcdsaGenerated_Kind + "." + CRDGroupVersion.String()
	KeystoreEcdsaGenerated_GroupVersionKind = CRDGroupVersion.WithKind(KeystoreEcdsaGenerated_Kind)
)

func init() {
	SchemeBuilder.Register(&KeystoreEcdsaGenerated{}, &KeystoreEcdsaGeneratedList{})
}