// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
)

// AdmissionregistrationV1Interface is an autogenerated mock type for the AdmissionregistrationV1Interface type
type AdmissionregistrationV1Interface struct {
	mock.Mock
}

// MutatingWebhookConfigurations provides a mock function with given fields:
func (_m *AdmissionregistrationV1Interface) MutatingWebhookConfigurations() v1.MutatingWebhookConfigurationInterface {
	ret := _m.Called()

	var r0 v1.MutatingWebhookConfigurationInterface
	if rf, ok := ret.Get(0).(func() v1.MutatingWebhookConfigurationInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.MutatingWebhookConfigurationInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *AdmissionregistrationV1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}

// ValidatingWebhookConfigurations provides a mock function with given fields:
func (_m *AdmissionregistrationV1Interface) ValidatingWebhookConfigurations() v1.ValidatingWebhookConfigurationInterface {
	ret := _m.Called()

	var r0 v1.ValidatingWebhookConfigurationInterface
	if rf, ok := ret.Get(0).(func() v1.ValidatingWebhookConfigurationInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.ValidatingWebhookConfigurationInterface)
		}
	}

	return r0
}
