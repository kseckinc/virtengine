// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// PersistentVolumesGetter is an autogenerated mock type for the PersistentVolumesGetter type
type PersistentVolumesGetter struct {
	mock.Mock
}

// PersistentVolumes provides a mock function with given fields:
func (_m *PersistentVolumesGetter) PersistentVolumes() v1.PersistentVolumeInterface {
	ret := _m.Called()

	var r0 v1.PersistentVolumeInterface
	if rf, ok := ret.Get(0).(func() v1.PersistentVolumeInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.PersistentVolumeInterface)
		}
	}

	return r0
}
