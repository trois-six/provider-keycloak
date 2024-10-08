// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/trois-six/provider-keycloak/apis/default/v1alpha1"
	v1alpha1group "github.com/trois-six/provider-keycloak/apis/group/v1alpha1"
	v1alpha1keycloak "github.com/trois-six/provider-keycloak/apis/keycloak/v1alpha1"
	v1alpha1openid "github.com/trois-six/provider-keycloak/apis/openid/v1alpha1"
	v1alpha1realm "github.com/trois-six/provider-keycloak/apis/realm/v1alpha1"
	v1alpha1apis "github.com/trois-six/provider-keycloak/apis/v1alpha1"
	v1beta1 "github.com/trois-six/provider-keycloak/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1group.SchemeBuilder.AddToScheme,
		v1alpha1keycloak.SchemeBuilder.AddToScheme,
		v1alpha1openid.SchemeBuilder.AddToScheme,
		v1alpha1realm.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
