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

// Code generated by ack-generate. DO NOT EDIT.

package addon

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/eks"
	svcsdk "github.com/aws/aws-sdk-go/service/eks"
	svcsdkapi "github.com/aws/aws-sdk-go/service/eks/eksiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/eks/v1alpha1"
	awsclient "github.com/crossplane/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an Addon resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create Addon in AWS"
	errUpdate        = "cannot update Addon in AWS"
	errDescribe      = "failed to describe Addon"
	errDelete        = "failed to delete Addon"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.Addon)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := awsclient.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.Addon)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeAddonInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeAddonWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateAddon(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

	upToDate, err := e.isUpToDate(cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.Addon)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateAddonInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateAddonWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.Addon.AddonArn != nil {
		cr.Status.AtProvider.AddonARN = resp.Addon.AddonArn
	} else {
		cr.Status.AtProvider.AddonARN = nil
	}
	if resp.Addon.ClusterName != nil {
		cr.Status.AtProvider.ClusterName = resp.Addon.ClusterName
	} else {
		cr.Status.AtProvider.ClusterName = nil
	}
	if resp.Addon.CreatedAt != nil {
		cr.Status.AtProvider.CreatedAt = &metav1.Time{*resp.Addon.CreatedAt}
	} else {
		cr.Status.AtProvider.CreatedAt = nil
	}
	if resp.Addon.Health != nil {
		f5 := &svcapitypes.AddonHealth{}
		if resp.Addon.Health.Issues != nil {
			f5f0 := []*svcapitypes.AddonIssue{}
			for _, f5f0iter := range resp.Addon.Health.Issues {
				f5f0elem := &svcapitypes.AddonIssue{}
				if f5f0iter.Code != nil {
					f5f0elem.Code = f5f0iter.Code
				}
				if f5f0iter.Message != nil {
					f5f0elem.Message = f5f0iter.Message
				}
				if f5f0iter.ResourceIds != nil {
					f5f0elemf2 := []*string{}
					for _, f5f0elemf2iter := range f5f0iter.ResourceIds {
						var f5f0elemf2elem string
						f5f0elemf2elem = *f5f0elemf2iter
						f5f0elemf2 = append(f5f0elemf2, &f5f0elemf2elem)
					}
					f5f0elem.ResourceIDs = f5f0elemf2
				}
				f5f0 = append(f5f0, f5f0elem)
			}
			f5.Issues = f5f0
		}
		cr.Status.AtProvider.Health = f5
	} else {
		cr.Status.AtProvider.Health = nil
	}
	if resp.Addon.ModifiedAt != nil {
		cr.Status.AtProvider.ModifiedAt = &metav1.Time{*resp.Addon.ModifiedAt}
	} else {
		cr.Status.AtProvider.ModifiedAt = nil
	}
	if resp.Addon.Status != nil {
		cr.Status.AtProvider.Status = resp.Addon.Status
	} else {
		cr.Status.AtProvider.Status = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.Addon)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateAddonInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateAddonWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.Addon)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteAddonInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteAddonWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.EKSAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.EKSAPI
	preObserve     func(context.Context, *svcapitypes.Addon, *svcsdk.DescribeAddonInput) error
	postObserve    func(context.Context, *svcapitypes.Addon, *svcsdk.DescribeAddonOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.AddonParameters, *svcsdk.DescribeAddonOutput) error
	isUpToDate     func(*svcapitypes.Addon, *svcsdk.DescribeAddonOutput) (bool, error)
	preCreate      func(context.Context, *svcapitypes.Addon, *svcsdk.CreateAddonInput) error
	postCreate     func(context.Context, *svcapitypes.Addon, *svcsdk.CreateAddonOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.Addon, *svcsdk.DeleteAddonInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.Addon, *svcsdk.DeleteAddonOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.Addon, *svcsdk.UpdateAddonInput) error
	postUpdate     func(context.Context, *svcapitypes.Addon, *svcsdk.UpdateAddonOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.Addon, *svcsdk.DescribeAddonInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.Addon, _ *svcsdk.DescribeAddonOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.AddonParameters, *svcsdk.DescribeAddonOutput) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.Addon, *svcsdk.DescribeAddonOutput) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.Addon, *svcsdk.CreateAddonInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.Addon, _ *svcsdk.CreateAddonOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.Addon, *svcsdk.DeleteAddonInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.Addon, _ *svcsdk.DeleteAddonOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.Addon, *svcsdk.UpdateAddonInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.Addon, _ *svcsdk.UpdateAddonOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
