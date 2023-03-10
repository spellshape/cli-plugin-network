// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/spellshape/network/x/monitoringc/types"
	"google.golang.org/grpc"
)

// MonitoringcClient is an autogenerated mock type for the QueryClient type
type MonitoringcClient struct {
	mock.Mock
}

// LaunchIDFromChannelID provides a mock function with given fields: ctx, in, opts
func (_m *MonitoringcClient) LaunchIDFromChannelID(ctx context.Context, in *types.QueryGetLaunchIDFromChannelIDRequest, opts ...grpc.CallOption) (*types.QueryGetLaunchIDFromChannelIDResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetLaunchIDFromChannelIDResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetLaunchIDFromChannelIDRequest, ...grpc.CallOption) *types.QueryGetLaunchIDFromChannelIDResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetLaunchIDFromChannelIDResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetLaunchIDFromChannelIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LaunchIDFromChannelIDAll provides a mock function with given fields: ctx, in, opts
func (_m *MonitoringcClient) LaunchIDFromChannelIDAll(ctx context.Context, in *types.QueryAllLaunchIDFromChannelIDRequest, opts ...grpc.CallOption) (*types.QueryAllLaunchIDFromChannelIDResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllLaunchIDFromChannelIDResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllLaunchIDFromChannelIDRequest, ...grpc.CallOption) *types.QueryAllLaunchIDFromChannelIDResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllLaunchIDFromChannelIDResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllLaunchIDFromChannelIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MonitoringHistory provides a mock function with given fields: ctx, in, opts
func (_m *MonitoringcClient) MonitoringHistory(ctx context.Context, in *types.QueryGetMonitoringHistoryRequest, opts ...grpc.CallOption) (*types.QueryGetMonitoringHistoryResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetMonitoringHistoryResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetMonitoringHistoryRequest, ...grpc.CallOption) *types.QueryGetMonitoringHistoryResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetMonitoringHistoryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetMonitoringHistoryRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Params provides a mock function with given fields: ctx, in, opts
func (_m *MonitoringcClient) Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryParamsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) *types.QueryParamsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryParamsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderClientID provides a mock function with given fields: ctx, in, opts
func (_m *MonitoringcClient) ProviderClientID(ctx context.Context, in *types.QueryGetProviderClientIDRequest, opts ...grpc.CallOption) (*types.QueryGetProviderClientIDResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetProviderClientIDResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetProviderClientIDRequest, ...grpc.CallOption) *types.QueryGetProviderClientIDResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetProviderClientIDResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetProviderClientIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderClientIDAll provides a mock function with given fields: ctx, in, opts
func (_m *MonitoringcClient) ProviderClientIDAll(ctx context.Context, in *types.QueryAllProviderClientIDRequest, opts ...grpc.CallOption) (*types.QueryAllProviderClientIDResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllProviderClientIDResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllProviderClientIDRequest, ...grpc.CallOption) *types.QueryAllProviderClientIDResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllProviderClientIDResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllProviderClientIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifiedClientIds provides a mock function with given fields: ctx, in, opts
func (_m *MonitoringcClient) VerifiedClientIds(ctx context.Context, in *types.QueryGetVerifiedClientIdsRequest, opts ...grpc.CallOption) (*types.QueryGetVerifiedClientIdsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetVerifiedClientIdsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetVerifiedClientIdsRequest, ...grpc.CallOption) *types.QueryGetVerifiedClientIdsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetVerifiedClientIdsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetVerifiedClientIdsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMonitoringcClient creates a new instance of MonitoringcClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMonitoringcClient(t testing.TB) *MonitoringcClient {
	mock := &MonitoringcClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
