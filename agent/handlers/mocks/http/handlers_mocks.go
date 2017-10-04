// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: net/http (interfaces: ResponseWriter)

package mock_http

import (
	http "net/http"

	gomock "github.com/golang/mock/gomock"
)

// Mock of ResponseWriter interface
type MockResponseWriter struct {
	ctrl     *gomock.Controller
	recorder *_MockResponseWriterRecorder
}

// Recorder for MockResponseWriter (not exported)
type _MockResponseWriterRecorder struct {
	mock *MockResponseWriter
}

func NewMockResponseWriter(ctrl *gomock.Controller) *MockResponseWriter {
	mock := &MockResponseWriter{ctrl: ctrl}
	mock.recorder = &_MockResponseWriterRecorder{mock}
	return mock
}

func (_m *MockResponseWriter) EXPECT() *_MockResponseWriterRecorder {
	return _m.recorder
}

func (_m *MockResponseWriter) Header() http.Header {
	ret := _m.ctrl.Call(_m, "Header")
	ret0, _ := ret[0].(http.Header)
	return ret0
}

func (_mr *_MockResponseWriterRecorder) Header() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Header")
}

func (_m *MockResponseWriter) Write(_param0 []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Write", _param0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockResponseWriterRecorder) Write(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0)
}

func (_m *MockResponseWriter) WriteHeader(_param0 int) {
	_m.ctrl.Call(_m, "WriteHeader", _param0)
}

func (_mr *_MockResponseWriterRecorder) WriteHeader(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteHeader", arg0)
}
