// Code generated by MockGen. DO NOT EDIT.
// Source: /usr/local/go/src/net/net.go

// Package mock_net is a generated GoMock package.
package mock

import (
	reflect "reflect"
	net "net"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockAddr is a mock of Addr interface.
type MockAddr struct {
	ctrl     *gomock.Controller
	recorder *MockAddrMockRecorder
}

// MockAddrMockRecorder is the mock recorder for MockAddr.
type MockAddrMockRecorder struct {
	mock *MockAddr
}

// NewMockAddr creates a new mock instance.
func NewMockAddr(ctrl *gomock.Controller) *MockAddr {
	mock := &MockAddr{ctrl: ctrl}
	mock.recorder = &MockAddrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddr) EXPECT() *MockAddrMockRecorder {
	return m.recorder
}

// Network mocks base method.
func (m *MockAddr) Network() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Network")
	ret0, _ := ret[0].(string)
	return ret0
}

// Network indicates an expected call of Network.
func (mr *MockAddrMockRecorder) Network() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Network", reflect.TypeOf((*MockAddr)(nil).Network))
}

// String mocks base method.
func (m *MockAddr) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockAddrMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockAddr)(nil).String))
}

// MockConn is a mock of Conn interface.
type MockConn struct {
	ctrl     *gomock.Controller
	recorder *MockConnMockRecorder
}

// MockConnMockRecorder is the mock recorder for MockConn.
type MockConnMockRecorder struct {
	mock *MockConn
}

// NewMockConn creates a new mock instance.
func NewMockConn(ctrl *gomock.Controller) *MockConn {
	mock := &MockConn{ctrl: ctrl}
	mock.recorder = &MockConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConn) EXPECT() *MockConnMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConnMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConn)(nil).Close))
}

// LocalAddr mocks base method.
func (m *MockConn) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockConnMockRecorder) LocalAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockConn)(nil).LocalAddr))
}

// Read mocks base method.
func (m *MockConn) Read(b []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", b)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockConnMockRecorder) Read(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockConn)(nil).Read), b)
}

// RemoteAddr mocks base method.
func (m *MockConn) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockConnMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockConn)(nil).RemoteAddr))
}

// SetDeadline mocks base method.
func (m *MockConn) SetDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline.
func (mr *MockConnMockRecorder) SetDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadline", reflect.TypeOf((*MockConn)(nil).SetDeadline), t)
}

// SetReadDeadline mocks base method.
func (m *MockConn) SetReadDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockConnMockRecorder) SetReadDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockConn)(nil).SetReadDeadline), t)
}

// SetWriteDeadline mocks base method.
func (m *MockConn) SetWriteDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockConnMockRecorder) SetWriteDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockConn)(nil).SetWriteDeadline), t)
}

// Write mocks base method.
func (m *MockConn) Write(b []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", b)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockConnMockRecorder) Write(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockConn)(nil).Write), b)
}

// MockPacketConn is a mock of PacketConn interface.
type MockPacketConn struct {
	ctrl     *gomock.Controller
	recorder *MockPacketConnMockRecorder
}

// MockPacketConnMockRecorder is the mock recorder for MockPacketConn.
type MockPacketConnMockRecorder struct {
	mock *MockPacketConn
}

// NewMockPacketConn creates a new mock instance.
func NewMockPacketConn(ctrl *gomock.Controller) *MockPacketConn {
	mock := &MockPacketConn{ctrl: ctrl}
	mock.recorder = &MockPacketConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacketConn) EXPECT() *MockPacketConnMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockPacketConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPacketConnMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPacketConn)(nil).Close))
}

// LocalAddr mocks base method.
func (m *MockPacketConn) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockPacketConnMockRecorder) LocalAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockPacketConn)(nil).LocalAddr))
}

// ReadFrom mocks base method.
func (m *MockPacketConn) ReadFrom(p []byte) (int, net.Addr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFrom", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(net.Addr)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadFrom indicates an expected call of ReadFrom.
func (mr *MockPacketConnMockRecorder) ReadFrom(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFrom", reflect.TypeOf((*MockPacketConn)(nil).ReadFrom), p)
}

// SetDeadline mocks base method.
func (m *MockPacketConn) SetDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline.
func (mr *MockPacketConnMockRecorder) SetDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadline", reflect.TypeOf((*MockPacketConn)(nil).SetDeadline), t)
}

// SetReadDeadline mocks base method.
func (m *MockPacketConn) SetReadDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockPacketConnMockRecorder) SetReadDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockPacketConn)(nil).SetReadDeadline), t)
}

// SetWriteDeadline mocks base method.
func (m *MockPacketConn) SetWriteDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockPacketConnMockRecorder) SetWriteDeadline(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockPacketConn)(nil).SetWriteDeadline), t)
}

// WriteTo mocks base method.
func (m *MockPacketConn) WriteTo(p []byte, addr net.Addr) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTo", p, addr)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteTo indicates an expected call of WriteTo.
func (mr *MockPacketConnMockRecorder) WriteTo(p, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTo", reflect.TypeOf((*MockPacketConn)(nil).WriteTo), p, addr)
}

// MockListener is a mock of Listener interface.
type MockListener struct {
	ctrl     *gomock.Controller
	recorder *MockListenerMockRecorder
}

// MockListenerMockRecorder is the mock recorder for MockListener.
type MockListenerMockRecorder struct {
	mock *MockListener
}

// NewMockListener creates a new mock instance.
func NewMockListener(ctrl *gomock.Controller) *MockListener {
	mock := &MockListener{ctrl: ctrl}
	mock.recorder = &MockListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListener) EXPECT() *MockListenerMockRecorder {
	return m.recorder
}

// Accept mocks base method.
func (m *MockListener) Accept() (net.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept")
	ret0, _ := ret[0].(net.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Accept indicates an expected call of Accept.
func (mr *MockListenerMockRecorder) Accept() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockListener)(nil).Accept))
}

// Addr mocks base method.
func (m *MockListener) Addr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// Addr indicates an expected call of Addr.
func (mr *MockListenerMockRecorder) Addr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addr", reflect.TypeOf((*MockListener)(nil).Addr))
}

// Close mocks base method.
func (m *MockListener) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockListenerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockListener)(nil).Close))
}

// MockError is a mock of Error interface.
type MockError struct {
	ctrl     *gomock.Controller
	recorder *MockErrorMockRecorder
}

// MockErrorMockRecorder is the mock recorder for MockError.
type MockErrorMockRecorder struct {
	mock *MockError
}

// NewMockError creates a new mock instance.
func NewMockError(ctrl *gomock.Controller) *MockError {
	mock := &MockError{ctrl: ctrl}
	mock.recorder = &MockErrorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockError) EXPECT() *MockErrorMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockError) Error() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockErrorMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockError)(nil).Error))
}

// Temporary mocks base method.
func (m *MockError) Temporary() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Temporary")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Temporary indicates an expected call of Temporary.
func (mr *MockErrorMockRecorder) Temporary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Temporary", reflect.TypeOf((*MockError)(nil).Temporary))
}

// Timeout mocks base method.
func (m *MockError) Timeout() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timeout")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Timeout indicates an expected call of Timeout.
func (mr *MockErrorMockRecorder) Timeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timeout", reflect.TypeOf((*MockError)(nil).Timeout))
}

// Mocktimeout is a mock of timeout interface.
type Mocktimeout struct {
	ctrl     *gomock.Controller
	recorder *MocktimeoutMockRecorder
}

// MocktimeoutMockRecorder is the mock recorder for Mocktimeout.
type MocktimeoutMockRecorder struct {
	mock *Mocktimeout
}

// NewMocktimeout creates a new mock instance.
func NewMocktimeout(ctrl *gomock.Controller) *Mocktimeout {
	mock := &Mocktimeout{ctrl: ctrl}
	mock.recorder = &MocktimeoutMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocktimeout) EXPECT() *MocktimeoutMockRecorder {
	return m.recorder
}

// Timeout mocks base method.
func (m *Mocktimeout) Timeout() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timeout")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Timeout indicates an expected call of Timeout.
func (mr *MocktimeoutMockRecorder) Timeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timeout", reflect.TypeOf((*Mocktimeout)(nil).Timeout))
}

// Mocktemporary is a mock of temporary interface.
type Mocktemporary struct {
	ctrl     *gomock.Controller
	recorder *MocktemporaryMockRecorder
}

// MocktemporaryMockRecorder is the mock recorder for Mocktemporary.
type MocktemporaryMockRecorder struct {
	mock *Mocktemporary
}

// NewMocktemporary creates a new mock instance.
func NewMocktemporary(ctrl *gomock.Controller) *Mocktemporary {
	mock := &Mocktemporary{ctrl: ctrl}
	mock.recorder = &MocktemporaryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocktemporary) EXPECT() *MocktemporaryMockRecorder {
	return m.recorder
}

// Temporary mocks base method.
func (m *Mocktemporary) Temporary() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Temporary")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Temporary indicates an expected call of Temporary.
func (mr *MocktemporaryMockRecorder) Temporary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Temporary", reflect.TypeOf((*Mocktemporary)(nil).Temporary))
}

// MockbuffersWriter is a mock of buffersWriter interface.
type MockbuffersWriter struct {
	ctrl     *gomock.Controller
	recorder *MockbuffersWriterMockRecorder
}

// MockbuffersWriterMockRecorder is the mock recorder for MockbuffersWriter.
type MockbuffersWriterMockRecorder struct {
	mock *MockbuffersWriter
}

// NewMockbuffersWriter creates a new mock instance.
func NewMockbuffersWriter(ctrl *gomock.Controller) *MockbuffersWriter {
	mock := &MockbuffersWriter{ctrl: ctrl}
	mock.recorder = &MockbuffersWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockbuffersWriter) EXPECT() *MockbuffersWriterMockRecorder {
	return m.recorder
}

// writeBuffers mocks base method.
func (m *MockbuffersWriter) writeBuffers(arg0 *net.Buffers) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "writeBuffers", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// writeBuffers indicates an expected call of writeBuffers.
func (mr *MockbuffersWriterMockRecorder) writeBuffers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "writeBuffers", reflect.TypeOf((*MockbuffersWriter)(nil).writeBuffers), arg0)
}
