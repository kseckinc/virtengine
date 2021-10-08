// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
)

// CSIStorageCapacitiesGetter is an autogenerated mock type for the CSIStorageCapacitiesGetter type
type CSIStorageCapacitiesGetter struct {
	mock.Mock
}

// CSIStorageCapacities provides a mock function with given fields: namespace
func (_m *CSIStorageCapacitiesGetter) CSIStorageCapacities(namespace string) v1alpha1.CSIStorageCapacityInterface {
	ret := _m.Called(namespace)

	var r0 v1alpha1.CSIStorageCapacityInterface
	if rf, ok := ret.Get(0).(func(string) v1alpha1.CSIStorageCapacityInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.CSIStorageCapacityInterface)
		}
	}

	return r0
}