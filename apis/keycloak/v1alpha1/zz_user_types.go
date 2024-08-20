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

type FederatedIdentityInitParameters struct {

	// The name of the identity provider
	IdentityProvider *string `json:"identityProvider,omitempty" tf:"identity_provider,omitempty"`

	// The user name of the user defined in the identity provider
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type FederatedIdentityObservation struct {

	// The name of the identity provider
	IdentityProvider *string `json:"identityProvider,omitempty" tf:"identity_provider,omitempty"`

	// The ID of the user defined in the identity provider
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// The user name of the user defined in the identity provider
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type FederatedIdentityParameters struct {

	// The name of the identity provider
	// +kubebuilder:validation:Optional
	IdentityProvider *string `json:"identityProvider" tf:"identity_provider,omitempty"`

	// The ID of the user defined in the identity provider
	// +kubebuilder:validation:Required
	UserID *string `json:"userId" tf:"user_id,omitempty"`

	// The user name of the user defined in the identity provider
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName" tf:"user_name,omitempty"`
}

type InitialPasswordInitParameters struct {

	// If set to true, the initial password is set up for renewal on first use. Default to false.
	Temporary *bool `json:"temporary,omitempty" tf:"temporary,omitempty"`

	// The initial password.
	ValueSecretRef v1.SecretKeySelector `json:"valueSecretRef" tf:"-"`
}

type InitialPasswordObservation struct {

	// If set to true, the initial password is set up for renewal on first use. Default to false.
	Temporary *bool `json:"temporary,omitempty" tf:"temporary,omitempty"`
}

type InitialPasswordParameters struct {

	// If set to true, the initial password is set up for renewal on first use. Default to false.
	// +kubebuilder:validation:Optional
	Temporary *bool `json:"temporary,omitempty" tf:"temporary,omitempty"`

	// The initial password.
	// +kubebuilder:validation:Optional
	ValueSecretRef v1.SecretKeySelector `json:"valueSecretRef" tf:"-"`
}

type UserInitParameters struct {

	// A map representing attributes for the user. In order to add multivalue attributes, use ## to seperate the values. Max length for each value is 255 chars
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// The user's email.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Whether the email address was validated or not. Default to false.
	EmailVerified *bool `json:"emailVerified,omitempty" tf:"email_verified,omitempty"`

	// When false, this user cannot log in. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// When specified, the user will be linked to a federated identity provider. Refer to the federated user example for more details.
	FederatedIdentity []FederatedIdentityInitParameters `json:"federatedIdentity,omitempty" tf:"federated_identity,omitempty"`

	// The user's first name.
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// When given, the user's initial password will be set. This attribute is only respected during initial user creation.
	InitialPassword []InitialPasswordInitParameters `json:"initialPassword,omitempty" tf:"initial_password,omitempty"`

	// The user's last name.
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// A list of required user actions.
	// +listType=set
	RequiredActions []*string `json:"requiredActions,omitempty" tf:"required_actions,omitempty"`

	// The unique username of this user.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserObservation struct {

	// A map representing attributes for the user. In order to add multivalue attributes, use ## to seperate the values. Max length for each value is 255 chars
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// The user's email.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Whether the email address was validated or not. Default to false.
	EmailVerified *bool `json:"emailVerified,omitempty" tf:"email_verified,omitempty"`

	// When false, this user cannot log in. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// When specified, the user will be linked to a federated identity provider. Refer to the federated user example for more details.
	FederatedIdentity []FederatedIdentityObservation `json:"federatedIdentity,omitempty" tf:"federated_identity,omitempty"`

	// The user's first name.
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// When given, the user's initial password will be set. This attribute is only respected during initial user creation.
	InitialPassword []InitialPasswordObservation `json:"initialPassword,omitempty" tf:"initial_password,omitempty"`

	// The user's last name.
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// The realm this user belongs to.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// A list of required user actions.
	// +listType=set
	RequiredActions []*string `json:"requiredActions,omitempty" tf:"required_actions,omitempty"`

	// The unique username of this user.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserParameters struct {

	// A map representing attributes for the user. In order to add multivalue attributes, use ## to seperate the values. Max length for each value is 255 chars
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// The user's email.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Whether the email address was validated or not. Default to false.
	// +kubebuilder:validation:Optional
	EmailVerified *bool `json:"emailVerified,omitempty" tf:"email_verified,omitempty"`

	// When false, this user cannot log in. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// When specified, the user will be linked to a federated identity provider. Refer to the federated user example for more details.
	// +kubebuilder:validation:Optional
	FederatedIdentity []FederatedIdentityParameters `json:"federatedIdentity,omitempty" tf:"federated_identity,omitempty"`

	// The user's first name.
	// +kubebuilder:validation:Optional
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// When given, the user's initial password will be set. This attribute is only respected during initial user creation.
	// +kubebuilder:validation:Optional
	InitialPassword []InitialPasswordParameters `json:"initialPassword,omitempty" tf:"initial_password,omitempty"`

	// The user's last name.
	// +kubebuilder:validation:Optional
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// The realm this user belongs to.
	// +kubebuilder:validation:Required
	RealmID *string `json:"realmId" tf:"realm_id,omitempty"`

	// A list of required user actions.
	// +kubebuilder:validation:Optional
	// +listType=set
	RequiredActions []*string `json:"requiredActions,omitempty" tf:"required_actions,omitempty"`

	// The unique username of this user.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
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
	InitProvider UserInitParameters `json:"initProvider,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// User is the Schema for the Users API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   UserSpec   `json:"spec"`
	Status UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
