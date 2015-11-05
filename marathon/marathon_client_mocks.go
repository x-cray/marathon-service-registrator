// Automatically generated by MockGen. DO NOT EDIT!
// Source: marathon/marathon_client.go

package marathon

import (
	go_marathon "github.com/gambol99/go-marathon"
	gomock "github.com/golang/mock/gomock"
	url "net/url"
)

// Mock of MarathonClient interface
type MockMarathonClient struct {
	ctrl     *gomock.Controller
	recorder *_MockMarathonClientRecorder
}

// Recorder for MockMarathonClient (not exported)
type _MockMarathonClientRecorder struct {
	mock *MockMarathonClient
}

func NewMockMarathonClient(ctrl *gomock.Controller) *MockMarathonClient {
	mock := &MockMarathonClient{ctrl: ctrl}
	mock.recorder = &_MockMarathonClientRecorder{mock}
	return mock
}

func (_m *MockMarathonClient) EXPECT() *_MockMarathonClientRecorder {
	return _m.recorder
}

func (_m *MockMarathonClient) Applications(_param0 url.Values) (*go_marathon.Applications, error) {
	ret := _m.ctrl.Call(_m, "Applications", _param0)
	ret0, _ := ret[0].(*go_marathon.Applications)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMarathonClientRecorder) Applications(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Applications", arg0)
}

func (_m *MockMarathonClient) AddEventsListener(channel go_marathon.EventsChannel, filter int) error {
	ret := _m.ctrl.Call(_m, "AddEventsListener", channel, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMarathonClientRecorder) AddEventsListener(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddEventsListener", arg0, arg1)
}

func (_m *MockMarathonClient) RemoveEventsListener(channel go_marathon.EventsChannel) {
	_m.ctrl.Call(_m, "RemoveEventsListener", channel)
}

func (_mr *_MockMarathonClientRecorder) RemoveEventsListener(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveEventsListener", arg0)
}
