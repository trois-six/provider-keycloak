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

type ManageMembersScopeInitParameters struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ManageMembersScopeObservation struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ManageMembersScopeParameters struct {

	// Decision strategy of the permission.
	// +kubebuilder:validation:Optional
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ManageMembershipScopeInitParameters struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ManageMembershipScopeObservation struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ManageMembershipScopeParameters struct {

	// Decision strategy of the permission.
	// +kubebuilder:validation:Optional
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ManageScopeInitParameters struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ManageScopeObservation struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ManageScopeParameters struct {

	// Decision strategy of the permission.
	// +kubebuilder:validation:Optional
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type PermissionsInitParameters struct {

	// The id of the group.
	// +crossplane:generate:reference:type=github.com/trois-six/provider-keycloak/apis/keycloak/v1alpha1.Group
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group in keycloak to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group in keycloak to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// Policies that decide if the admin can manage the users that belong to this group.
	ManageMembersScope []ManageMembersScopeInitParameters `json:"manageMembersScope,omitempty" tf:"manage_members_scope,omitempty"`

	// Policies that decide if an admin can change the membership of the group. Add or remove members from the group.
	ManageMembershipScope []ManageMembershipScopeInitParameters `json:"manageMembershipScope,omitempty" tf:"manage_membership_scope,omitempty"`

	// Policies that decide if the admin can manage the configuration of the group.
	ManageScope []ManageScopeInitParameters `json:"manageScope,omitempty" tf:"manage_scope,omitempty"`

	// Policies that decide if the admin can view the user details of members of the group.
	ViewMembersScope []ViewMembersScopeInitParameters `json:"viewMembersScope,omitempty" tf:"view_members_scope,omitempty"`

	// Policies that decide if the admin can view information about the group.
	ViewScope []ViewScopeInitParameters `json:"viewScope,omitempty" tf:"view_scope,omitempty"`
}

type PermissionsObservation struct {

	// Resource server id representing the realm management client on which these permissions are managed.
	// Resource server id representing the realm management client on which this permission is managed
	AuthorizationResourceServerID *string `json:"authorizationResourceServerId,omitempty" tf:"authorization_resource_server_id,omitempty"`

	// When true, this indicates that fine-grained role permissions are enabled. This will always be true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The id of the group.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Policies that decide if the admin can manage the users that belong to this group.
	ManageMembersScope []ManageMembersScopeObservation `json:"manageMembersScope,omitempty" tf:"manage_members_scope,omitempty"`

	// Policies that decide if an admin can change the membership of the group. Add or remove members from the group.
	ManageMembershipScope []ManageMembershipScopeObservation `json:"manageMembershipScope,omitempty" tf:"manage_membership_scope,omitempty"`

	// Policies that decide if the admin can manage the configuration of the group.
	ManageScope []ManageScopeObservation `json:"manageScope,omitempty" tf:"manage_scope,omitempty"`

	// The realm in which to manage fine-grained role permissions.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Policies that decide if the admin can view the user details of members of the group.
	ViewMembersScope []ViewMembersScopeObservation `json:"viewMembersScope,omitempty" tf:"view_members_scope,omitempty"`

	// Policies that decide if the admin can view information about the group.
	ViewScope []ViewScopeObservation `json:"viewScope,omitempty" tf:"view_scope,omitempty"`
}

type PermissionsParameters struct {

	// The id of the group.
	// +crossplane:generate:reference:type=github.com/trois-six/provider-keycloak/apis/keycloak/v1alpha1.Group
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group in keycloak to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group in keycloak to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// Policies that decide if the admin can manage the users that belong to this group.
	// +kubebuilder:validation:Optional
	ManageMembersScope []ManageMembersScopeParameters `json:"manageMembersScope,omitempty" tf:"manage_members_scope,omitempty"`

	// Policies that decide if an admin can change the membership of the group. Add or remove members from the group.
	// +kubebuilder:validation:Optional
	ManageMembershipScope []ManageMembershipScopeParameters `json:"manageMembershipScope,omitempty" tf:"manage_membership_scope,omitempty"`

	// Policies that decide if the admin can manage the configuration of the group.
	// +kubebuilder:validation:Optional
	ManageScope []ManageScopeParameters `json:"manageScope,omitempty" tf:"manage_scope,omitempty"`

	// The realm in which to manage fine-grained role permissions.
	// +crossplane:generate:reference:type=github.com/trois-six/provider-keycloak/apis/keycloak/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in keycloak to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in keycloak to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// Policies that decide if the admin can view the user details of members of the group.
	// +kubebuilder:validation:Optional
	ViewMembersScope []ViewMembersScopeParameters `json:"viewMembersScope,omitempty" tf:"view_members_scope,omitempty"`

	// Policies that decide if the admin can view information about the group.
	// +kubebuilder:validation:Optional
	ViewScope []ViewScopeParameters `json:"viewScope,omitempty" tf:"view_scope,omitempty"`
}

type ViewMembersScopeInitParameters struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ViewMembersScopeObservation struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ViewMembersScopeParameters struct {

	// Decision strategy of the permission.
	// +kubebuilder:validation:Optional
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ViewScopeInitParameters struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ViewScopeObservation struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ViewScopeParameters struct {

	// Decision strategy of the permission.
	// +kubebuilder:validation:Optional
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

// PermissionsSpec defines the desired state of Permissions
type PermissionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PermissionsParameters `json:"forProvider"`
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
	InitProvider PermissionsInitParameters `json:"initProvider,omitempty"`
}

// PermissionsStatus defines the observed state of Permissions.
type PermissionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PermissionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Permissions is the Schema for the Permissionss API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type Permissions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PermissionsSpec   `json:"spec"`
	Status            PermissionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionsList contains a list of Permissionss
type PermissionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Permissions `json:"items"`
}

// Repository type metadata.
var (
	Permissions_Kind             = "Permissions"
	Permissions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Permissions_Kind}.String()
	Permissions_KindAPIVersion   = Permissions_Kind + "." + CRDGroupVersion.String()
	Permissions_GroupVersionKind = CRDGroupVersion.WithKind(Permissions_Kind)
)

func init() {
	SchemeBuilder.Register(&Permissions{}, &PermissionsList{})
}
