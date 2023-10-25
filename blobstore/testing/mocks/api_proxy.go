// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cubefs/cubefs/blobstore/api/proxy (interfaces: Client,LbMsgSender)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	blobnode "github.com/cubefs/cubefs/blobstore/api/blobnode"
	clustermgr "github.com/cubefs/cubefs/blobstore/api/clustermgr"
	proxy "github.com/cubefs/cubefs/blobstore/api/proxy"
	gomock "github.com/golang/mock/gomock"
)

// MockProxyClient is a mock of Client interface.
type MockProxyClient struct {
	ctrl     *gomock.Controller
	recorder *MockProxyClientMockRecorder
}

// MockProxyClientMockRecorder is the mock recorder for MockProxyClient.
type MockProxyClientMockRecorder struct {
	mock *MockProxyClient
}

// NewMockProxyClient creates a new mock instance.
func NewMockProxyClient(ctrl *gomock.Controller) *MockProxyClient {
	mock := &MockProxyClient{ctrl: ctrl}
	mock.recorder = &MockProxyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProxyClient) EXPECT() *MockProxyClientMockRecorder {
	return m.recorder
}

// Erase mocks base method.
func (m *MockProxyClient) Erase(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Erase", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Erase indicates an expected call of Erase.
func (mr *MockProxyClientMockRecorder) Erase(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Erase", reflect.TypeOf((*MockProxyClient)(nil).Erase), arg0, arg1, arg2)
}

// GetCacheDisk mocks base method.
func (m *MockProxyClient) GetCacheDisk(arg0 context.Context, arg1 string, arg2 *clustermgr.CacheDiskArgs) (*blobnode.DiskInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCacheDisk", arg0, arg1, arg2)
	ret0, _ := ret[0].(*blobnode.DiskInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCacheDisk indicates an expected call of GetCacheDisk.
func (mr *MockProxyClientMockRecorder) GetCacheDisk(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCacheDisk", reflect.TypeOf((*MockProxyClient)(nil).GetCacheDisk), arg0, arg1, arg2)
}

// GetCacheVolume mocks base method.
func (m *MockProxyClient) GetCacheVolume(arg0 context.Context, arg1 string, arg2 *clustermgr.CacheVolumeArgs) (*clustermgr.VersionVolume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCacheVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(*clustermgr.VersionVolume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCacheVolume indicates an expected call of GetCacheVolume.
func (mr *MockProxyClientMockRecorder) GetCacheVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCacheVolume", reflect.TypeOf((*MockProxyClient)(nil).GetCacheVolume), arg0, arg1, arg2)
}

// ListVolumes mocks base method.
func (m *MockProxyClient) ListVolumes(arg0 context.Context, arg1 string, arg2 *proxy.ListVolsArgs) (proxy.VolumeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolumes", arg0, arg1, arg2)
	ret0, _ := ret[0].(proxy.VolumeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVolumes indicates an expected call of ListVolumes.
func (mr *MockProxyClientMockRecorder) ListVolumes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolumes", reflect.TypeOf((*MockProxyClient)(nil).ListVolumes), arg0, arg1, arg2)
}

// SendDeleteMsg mocks base method.
func (m *MockProxyClient) SendDeleteMsg(arg0 context.Context, arg1 string, arg2 *proxy.DeleteArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDeleteMsg", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendDeleteMsg indicates an expected call of SendDeleteMsg.
func (mr *MockProxyClientMockRecorder) SendDeleteMsg(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDeleteMsg", reflect.TypeOf((*MockProxyClient)(nil).SendDeleteMsg), arg0, arg1, arg2)
}

// SendShardRepairMsg mocks base method.
func (m *MockProxyClient) SendShardRepairMsg(arg0 context.Context, arg1 string, arg2 *proxy.ShardRepairArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendShardRepairMsg", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendShardRepairMsg indicates an expected call of SendShardRepairMsg.
func (mr *MockProxyClientMockRecorder) SendShardRepairMsg(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendShardRepairMsg", reflect.TypeOf((*MockProxyClient)(nil).SendShardRepairMsg), arg0, arg1, arg2)
}

// VolumeAlloc mocks base method.
func (m *MockProxyClient) VolumeAlloc(arg0 context.Context, arg1 string, arg2 *proxy.AllocVolsArgs) ([]proxy.AllocRet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeAlloc", arg0, arg1, arg2)
	ret0, _ := ret[0].([]proxy.AllocRet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeAlloc indicates an expected call of VolumeAlloc.
func (mr *MockProxyClientMockRecorder) VolumeAlloc(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeAlloc", reflect.TypeOf((*MockProxyClient)(nil).VolumeAlloc), arg0, arg1, arg2)
}

// MockProxyLbRpcClient is a mock of LbMsgSender interface.
type MockProxyLbRpcClient struct {
	ctrl     *gomock.Controller
	recorder *MockProxyLbRpcClientMockRecorder
}

// MockProxyLbRpcClientMockRecorder is the mock recorder for MockProxyLbRpcClient.
type MockProxyLbRpcClientMockRecorder struct {
	mock *MockProxyLbRpcClient
}

// NewMockProxyLbRpcClient creates a new mock instance.
func NewMockProxyLbRpcClient(ctrl *gomock.Controller) *MockProxyLbRpcClient {
	mock := &MockProxyLbRpcClient{ctrl: ctrl}
	mock.recorder = &MockProxyLbRpcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProxyLbRpcClient) EXPECT() *MockProxyLbRpcClientMockRecorder {
	return m.recorder
}

// SendDeleteMsg mocks base method.
func (m *MockProxyLbRpcClient) SendDeleteMsg(arg0 context.Context, arg1 *proxy.DeleteArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDeleteMsg", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendDeleteMsg indicates an expected call of SendDeleteMsg.
func (mr *MockProxyLbRpcClientMockRecorder) SendDeleteMsg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDeleteMsg", reflect.TypeOf((*MockProxyLbRpcClient)(nil).SendDeleteMsg), arg0, arg1)
}

// SendShardRepairMsg mocks base method.
func (m *MockProxyLbRpcClient) SendShardRepairMsg(arg0 context.Context, arg1 *proxy.ShardRepairArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendShardRepairMsg", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendShardRepairMsg indicates an expected call of SendShardRepairMsg.
func (mr *MockProxyLbRpcClientMockRecorder) SendShardRepairMsg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendShardRepairMsg", reflect.TypeOf((*MockProxyLbRpcClient)(nil).SendShardRepairMsg), arg0, arg1)
}
