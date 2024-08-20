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

type ClientDefaultScopesInitParameters struct {

	// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
	// +crossplane:generate:reference:type=github.com/trois-six/provider-keycloak/apis/openid/v1alpha1.Client
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a Client in openid to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a Client in openid to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// An array of client scope names to attach to this client.
	// +listType=set
	DefaultScopes []*string `json:"defaultScopes,omitempty" tf:"default_scopes,omitempty"`

	// The realm this client and scopes exists in.
	// +crossplane:generate:reference:type=github.com/trois-six/provider-keycloak/apis/keycloak/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in keycloak to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in keycloak to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

type ClientDefaultScopesObservation struct {

	// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// An array of client scope names to attach to this client.
	// +listType=set
	DefaultScopes []*string `json:"defaultScopes,omitempty" tf:"default_scopes,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The realm this client and scopes exists in.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`
}

type ClientDefaultScopesParameters struct {

	// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
	// +crossplane:generate:reference:type=github.com/trois-six/provider-keycloak/apis/openid/v1alpha1.Client
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a Client in openid to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a Client in openid to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// An array of client scope names to attach to this client.
	// +kubebuilder:validation:Optional
	// +listType=set
	DefaultScopes []*string `json:"defaultScopes,omitempty" tf:"default_scopes,omitempty"`

	// The realm this client and scopes exists in.
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

// ClientDefaultScopesSpec defines the desired state of ClientDefaultScopes
type ClientDefaultScopesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClientDefaultScopesParameters `json:"forProvider"`
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
	InitProvider ClientDefaultScopesInitParameters `json:"initProvider,omitempty"`
}

// ClientDefaultScopesStatus defines the observed state of ClientDefaultScopes.
type ClientDefaultScopesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClientDefaultScopesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ClientDefaultScopes is the Schema for the ClientDefaultScopess API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type ClientDefaultScopes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultScopes) || (has(self.initProvider) && has(self.initProvider.defaultScopes))",message="spec.forProvider.defaultScopes is a required parameter"
	Spec   ClientDefaultScopesSpec   `json:"spec"`
	Status ClientDefaultScopesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClientDefaultScopesList contains a list of ClientDefaultScopess
type ClientDefaultScopesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClientDefaultScopes `json:"items"`
}

// Repository type metadata.
var (
	ClientDefaultScopes_Kind             = "ClientDefaultScopes"
	ClientDefaultScopes_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClientDefaultScopes_Kind}.String()
	ClientDefaultScopes_KindAPIVersion   = ClientDefaultScopes_Kind + "." + CRDGroupVersion.String()
	ClientDefaultScopes_GroupVersionKind = CRDGroupVersion.WithKind(ClientDefaultScopes_Kind)
)

func init() {
	SchemeBuilder.Register(&ClientDefaultScopes{}, &ClientDefaultScopesList{})
}
