// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	compute "google.golang.org/api/compute/v1"

	mock "github.com/stretchr/testify/mock"
)

// ComputeAPI is an autogenerated mock type for the ComputeAPI type
type ComputeAPI struct {
	mock.Mock
}

// DeleteIPAddress provides a mock function with given fields: project, region, address
func (_m *ComputeAPI) DeleteIPAddress(project string, region string, address string) error {
	ret := _m.Called(project, region, address)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(project, region, address)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LookupIPAddresses provides a mock function with given fields: project, region
func (_m *ComputeAPI) LookupIPAddresses(project string, region string) ([]*compute.Address, error) {
	ret := _m.Called(project, region)

	var r0 []*compute.Address
	if rf, ok := ret.Get(0).(func(string, string) []*compute.Address); ok {
		r0 = rf(project, region)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*compute.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(project, region)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
