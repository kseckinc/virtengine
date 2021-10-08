// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

// IngressesGetter is an autogenerated mock type for the IngressesGetter type
type IngressesGetter struct {
	mock.Mock
}

// Ingresses provides a mock function with given fields: namespace
func (_m *IngressesGetter) Ingresses(namespace string) v1.IngressInterface {
	ret := _m.Called(namespace)

	var r0 v1.IngressInterface
	if rf, ok := ret.Get(0).(func(string) v1.IngressInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.IngressInterface)
		}
	}

	return r0
}
