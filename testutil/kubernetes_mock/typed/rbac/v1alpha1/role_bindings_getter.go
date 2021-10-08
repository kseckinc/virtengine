// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
)

// RoleBindingsGetter is an autogenerated mock type for the RoleBindingsGetter type
type RoleBindingsGetter struct {
	mock.Mock
}

// RoleBindings provides a mock function with given fields: namespace
func (_m *RoleBindingsGetter) RoleBindings(namespace string) v1alpha1.RoleBindingInterface {
	ret := _m.Called(namespace)

	var r0 v1alpha1.RoleBindingInterface
	if rf, ok := ret.Get(0).(func(string) v1alpha1.RoleBindingInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.RoleBindingInterface)
		}
	}

	return r0
}