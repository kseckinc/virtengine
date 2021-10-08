// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
)

// StatefulSetsGetter is an autogenerated mock type for the StatefulSetsGetter type
type StatefulSetsGetter struct {
	mock.Mock
}

// StatefulSets provides a mock function with given fields: namespace
func (_m *StatefulSetsGetter) StatefulSets(namespace string) v1beta1.StatefulSetInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.StatefulSetInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.StatefulSetInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.StatefulSetInterface)
		}
	}

	return r0
}
