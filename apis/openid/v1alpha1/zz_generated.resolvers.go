// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/trois-six/provider-keycloak/apis/keycloak/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Client.
func (mg *Client) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientAggregatePolicy.
func (mg *ClientAggregatePolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientAuthorizationPermission.
func (mg *ClientAuthorizationPermission) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientAuthorizationResource.
func (mg *ClientAuthorizationResource) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientAuthorizationScope.
func (mg *ClientAuthorizationScope) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientClientPolicy.
func (mg *ClientClientPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientDefaultScopes.
func (mg *ClientDefaultScopes) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClientIDRef,
		Selector:     mg.Spec.ForProvider.ClientIDSelector,
		To: reference.To{
			List:    &ClientList{},
			Managed: &Client{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientID")
	}
	mg.Spec.ForProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClientID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClientIDRef,
		Selector:     mg.Spec.InitProvider.ClientIDSelector,
		To: reference.To{
			List:    &ClientList{},
			Managed: &Client{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClientID")
	}
	mg.Spec.InitProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientGroupPolicy.
func (mg *ClientGroupPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientJsPolicy.
func (mg *ClientJsPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientOptionalScopes.
func (mg *ClientOptionalScopes) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClientIDRef,
		Selector:     mg.Spec.ForProvider.ClientIDSelector,
		To: reference.To{
			List:    &ClientList{},
			Managed: &Client{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientID")
	}
	mg.Spec.ForProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClientID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClientIDRef,
		Selector:     mg.Spec.InitProvider.ClientIDSelector,
		To: reference.To{
			List:    &ClientList{},
			Managed: &Client{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClientID")
	}
	mg.Spec.InitProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientPermissions.
func (mg *ClientPermissions) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClientIDRef,
		Selector:     mg.Spec.ForProvider.ClientIDSelector,
		To: reference.To{
			List:    &ClientList{},
			Managed: &Client{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientID")
	}
	mg.Spec.ForProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientRolePolicy.
func (mg *ClientRolePolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientScope.
func (mg *ClientScope) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientServiceAccountRealmRole.
func (mg *ClientServiceAccountRealmRole) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RoleRef,
		Selector:     mg.Spec.ForProvider.RoleSelector,
		To: reference.To{
			List:    &v1alpha1.RoleList{},
			Managed: &v1alpha1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountUserID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountUserIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountUserIDSelector,
		To: reference.To{
			List:    &ClientList{},
			Managed: &Client{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountUserID")
	}
	mg.Spec.ForProvider.ServiceAccountUserID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountUserIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Role),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RoleRef,
		Selector:     mg.Spec.InitProvider.RoleSelector,
		To: reference.To{
			List:    &v1alpha1.RoleList{},
			Managed: &v1alpha1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Role")
	}
	mg.Spec.InitProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientServiceAccountRole.
func (mg *ClientServiceAccountRole) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClientIDRef,
		Selector:     mg.Spec.ForProvider.ClientIDSelector,
		To: reference.To{
			List:    &ClientList{},
			Managed: &Client{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientID")
	}
	mg.Spec.ForProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RoleRef,
		Selector:     mg.Spec.ForProvider.RoleSelector,
		To: reference.To{
			List:    &v1alpha1.RoleList{},
			Managed: &v1alpha1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountUserID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountUserIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountUserIDSelector,
		To: reference.To{
			List:    &ClientList{},
			Managed: &Client{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountUserID")
	}
	mg.Spec.ForProvider.ServiceAccountUserID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountUserIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Role),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RoleRef,
		Selector:     mg.Spec.InitProvider.RoleSelector,
		To: reference.To{
			List:    &v1alpha1.RoleList{},
			Managed: &v1alpha1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Role")
	}
	mg.Spec.InitProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientTimePolicy.
func (mg *ClientTimePolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClientUserPolicy.
func (mg *ClientUserPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RealmIDRef,
		Selector:     mg.Spec.ForProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RealmID")
	}
	mg.Spec.ForProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RealmIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RealmID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RealmIDRef,
		Selector:     mg.Spec.InitProvider.RealmIDSelector,
		To: reference.To{
			List:    &v1alpha1.RealmList{},
			Managed: &v1alpha1.Realm{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RealmID")
	}
	mg.Spec.InitProvider.RealmID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RealmIDRef = rsp.ResolvedReference

	return nil
}
