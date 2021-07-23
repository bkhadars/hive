// Code generated by MockGen. DO NOT EDIT.
// Source: ./hibernation_actuator.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/openshift/hive/apis/hive/v1"
	logrus "github.com/sirupsen/logrus"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockHibernationActuator is a mock of HibernationActuator interface
type MockHibernationActuator struct {
	ctrl     *gomock.Controller
	recorder *MockHibernationActuatorMockRecorder
}

// MockHibernationActuatorMockRecorder is the mock recorder for MockHibernationActuator
type MockHibernationActuatorMockRecorder struct {
	mock *MockHibernationActuator
}

// NewMockHibernationActuator creates a new mock instance
func NewMockHibernationActuator(ctrl *gomock.Controller) *MockHibernationActuator {
	mock := &MockHibernationActuator{ctrl: ctrl}
	mock.recorder = &MockHibernationActuatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHibernationActuator) EXPECT() *MockHibernationActuatorMockRecorder {
	return m.recorder
}

// CanHandle mocks base method
func (m *MockHibernationActuator) CanHandle(cd *v1.ClusterDeployment) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanHandle", cd)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanHandle indicates an expected call of CanHandle
func (mr *MockHibernationActuatorMockRecorder) CanHandle(cd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanHandle", reflect.TypeOf((*MockHibernationActuator)(nil).CanHandle), cd)
}

// StopMachines mocks base method
func (m *MockHibernationActuator) StopMachines(cd *v1.ClusterDeployment, hiveClient client.Client, logger logrus.FieldLogger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopMachines", cd, hiveClient, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopMachines indicates an expected call of StopMachines
func (mr *MockHibernationActuatorMockRecorder) StopMachines(cd, hiveClient, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopMachines", reflect.TypeOf((*MockHibernationActuator)(nil).StopMachines), cd, hiveClient, logger)
}

// StartMachines mocks base method
func (m *MockHibernationActuator) StartMachines(cd *v1.ClusterDeployment, hiveClient client.Client, logger logrus.FieldLogger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartMachines", cd, hiveClient, logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartMachines indicates an expected call of StartMachines
func (mr *MockHibernationActuatorMockRecorder) StartMachines(cd, hiveClient, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMachines", reflect.TypeOf((*MockHibernationActuator)(nil).StartMachines), cd, hiveClient, logger)
}

// MachinesRunning mocks base method
func (m *MockHibernationActuator) MachinesRunning(cd *v1.ClusterDeployment, hiveClient client.Client, logger logrus.FieldLogger) (bool, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MachinesRunning", cd, hiveClient, logger)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MachinesRunning indicates an expected call of MachinesRunning
func (mr *MockHibernationActuatorMockRecorder) MachinesRunning(cd, hiveClient, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachinesRunning", reflect.TypeOf((*MockHibernationActuator)(nil).MachinesRunning), cd, hiveClient, logger)
}

// MachinesStopped mocks base method
func (m *MockHibernationActuator) MachinesStopped(cd *v1.ClusterDeployment, hiveClient client.Client, logger logrus.FieldLogger) (bool, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MachinesStopped", cd, hiveClient, logger)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MachinesStopped indicates an expected call of MachinesStopped
func (mr *MockHibernationActuatorMockRecorder) MachinesStopped(cd, hiveClient, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachinesStopped", reflect.TypeOf((*MockHibernationActuator)(nil).MachinesStopped), cd, hiveClient, logger)
}

// MockHibernationPreemptibleMachines is a mock of HibernationPreemptibleMachines interface
type MockHibernationPreemptibleMachines struct {
	ctrl     *gomock.Controller
	recorder *MockHibernationPreemptibleMachinesMockRecorder
}

// MockHibernationPreemptibleMachinesMockRecorder is the mock recorder for MockHibernationPreemptibleMachines
type MockHibernationPreemptibleMachinesMockRecorder struct {
	mock *MockHibernationPreemptibleMachines
}

// NewMockHibernationPreemptibleMachines creates a new mock instance
func NewMockHibernationPreemptibleMachines(ctrl *gomock.Controller) *MockHibernationPreemptibleMachines {
	mock := &MockHibernationPreemptibleMachines{ctrl: ctrl}
	mock.recorder = &MockHibernationPreemptibleMachinesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHibernationPreemptibleMachines) EXPECT() *MockHibernationPreemptibleMachinesMockRecorder {
	return m.recorder
}

// ReplaceMachines mocks base method
func (m *MockHibernationPreemptibleMachines) ReplaceMachines(cd *v1.ClusterDeployment, remoteClient client.Client, logger logrus.FieldLogger) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplaceMachines", cd, remoteClient, logger)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReplaceMachines indicates an expected call of ReplaceMachines
func (mr *MockHibernationPreemptibleMachinesMockRecorder) ReplaceMachines(cd, remoteClient, logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceMachines", reflect.TypeOf((*MockHibernationPreemptibleMachines)(nil).ReplaceMachines), cd, remoteClient, logger)
}
