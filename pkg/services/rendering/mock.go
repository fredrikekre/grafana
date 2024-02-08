// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/grafana/grafana/pkg/services/rendering (interfaces: Service)

package rendering

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	models "github.com/grafana/grafana/pkg/models"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateRenderingSession mocks base method.
func (m *MockService) CreateRenderingSession(ctx context.Context, authOpts AuthOpts, sessionOpts SessionOpts) (Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRenderingSession", ctx, authOpts, sessionOpts)
	ret0, _ := ret[0].(Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRenderingSession indicates an expected call of CreateRenderingSession.
func (mr *MockServiceMockRecorder) CreateRenderingSession(ctx, authOpts, sessionOpts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRenderingSession", reflect.TypeOf((*MockService)(nil).CreateRenderingSession), ctx, authOpts, sessionOpts)
}

// GetRenderUser mocks base method.
func (m *MockService) GetRenderUser(ctx context.Context, key string) (*RenderUser, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRenderUser", ctx, key)
	ret0, _ := ret[0].(*RenderUser)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetRenderUser indicates an expected call of GetRenderUser.
func (mr *MockServiceMockRecorder) GetRenderUser(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRenderUser", reflect.TypeOf((*MockService)(nil).GetRenderUser), ctx, key)
}

// HasCapability mocks base method.
func (m *MockService) HasCapability(ctx context.Context, capability CapabilityName) (CapabilitySupportRequestResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasCapability", ctx, capability)
	ret0, _ := ret[0].(CapabilitySupportRequestResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasCapability indicates an expected call of HasCapability.
func (mr *MockServiceMockRecorder) HasCapability(ctx, capability interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasCapability", reflect.TypeOf((*MockService)(nil).HasCapability), ctx, capability)
}

// IsAvailable mocks base method.
func (m *MockService) IsAvailable(ctx context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAvailable", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAvailable indicates an expected call of IsAvailable.
func (mr *MockServiceMockRecorder) IsAvailable(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAvailable", reflect.TypeOf((*MockService)(nil).IsAvailable), ctx)
}

// Render mocks base method.
func (m *MockService) Render(ctx context.Context, renderType RenderType, opts Opts, session Session) (*RenderResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Render", ctx, renderType, opts, session)
	ret0, _ := ret[0].(*RenderResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Render indicates an expected call of Render.
func (mr *MockServiceMockRecorder) Render(ctx, renderType, opts, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockService)(nil).Render), ctx, renderType, opts, session)
}

// RenderCSV mocks base method.
func (m *MockService) RenderCSV(ctx context.Context, opts CSVOpts, session Session) (*RenderCSVResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenderCSV", ctx, opts, session)
	ret0, _ := ret[0].(*RenderCSVResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenderCSV indicates an expected call of RenderCSV.
func (mr *MockServiceMockRecorder) RenderCSV(ctx, opts, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderCSV", reflect.TypeOf((*MockService)(nil).RenderCSV), ctx, opts, session)
}

// RenderErrorImage mocks base method.
func (m *MockService) RenderErrorImage(theme models.Theme, err error) (*RenderResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenderErrorImage", theme, err)
	ret0, _ := ret[0].(*RenderResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenderErrorImage indicates an expected call of RenderErrorImage.
func (mr *MockServiceMockRecorder) RenderErrorImage(theme, error interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderErrorImage", reflect.TypeOf((*MockService)(nil).RenderErrorImage), theme, error)
}

// SanitizeSVG mocks base method.
func (m *MockService) SanitizeSVG(ctx context.Context, req *SanitizeSVGRequest) (*SanitizeSVGResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SanitizeSVG", ctx, req)
	ret0, _ := ret[0].(*SanitizeSVGResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SanitizeSVG indicates an expected call of SanitizeSVG.
func (mr *MockServiceMockRecorder) SanitizeSVG(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SanitizeSVG", reflect.TypeOf((*MockService)(nil).SanitizeSVG), ctx, req)
}

// Version mocks base method.
func (m *MockService) Version() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockServiceMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockService)(nil).Version))
}
