// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "k8s.io/api/authorization/v1beta1"
)

// SelfSubjectAccessReviewInterface is an autogenerated mock type for the SelfSubjectAccessReviewInterface type
type SelfSubjectAccessReviewInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, selfSubjectAccessReview, opts
func (_m *SelfSubjectAccessReviewInterface) Create(ctx context.Context, selfSubjectAccessReview *v1beta1.SelfSubjectAccessReview, opts v1.CreateOptions) (*v1beta1.SelfSubjectAccessReview, error) {
	ret := _m.Called(ctx, selfSubjectAccessReview, opts)

	var r0 *v1beta1.SelfSubjectAccessReview
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.SelfSubjectAccessReview, v1.CreateOptions) *v1beta1.SelfSubjectAccessReview); ok {
		r0 = rf(ctx, selfSubjectAccessReview, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.SelfSubjectAccessReview)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.SelfSubjectAccessReview, v1.CreateOptions) error); ok {
		r1 = rf(ctx, selfSubjectAccessReview, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
