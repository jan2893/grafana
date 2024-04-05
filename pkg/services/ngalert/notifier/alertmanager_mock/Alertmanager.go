// Code generated by mockery v2.42.1. DO NOT EDIT.

package alertmanager_mock

import (
	context "context"

	definitions "github.com/grafana/grafana/pkg/services/ngalert/api/tooling/definitions"
	mock "github.com/stretchr/testify/mock"

	models "github.com/grafana/grafana/pkg/services/ngalert/models"

	notifier "github.com/grafana/grafana/pkg/services/ngalert/notifier"

	notify "github.com/grafana/alerting/notify"

	v2models "github.com/prometheus/alertmanager/api/v2/models"
)

// AlertmanagerMock is an autogenerated mock type for the Alertmanager type
type AlertmanagerMock struct {
	mock.Mock
}

type AlertmanagerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *AlertmanagerMock) EXPECT() *AlertmanagerMock_Expecter {
	return &AlertmanagerMock_Expecter{mock: &_m.Mock}
}

// ApplyConfig provides a mock function with given fields: _a0, _a1
func (_m *AlertmanagerMock) ApplyConfig(_a0 context.Context, _a1 *models.AlertConfiguration) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ApplyConfig")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.AlertConfiguration) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AlertmanagerMock_ApplyConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyConfig'
type AlertmanagerMock_ApplyConfig_Call struct {
	*mock.Call
}

// ApplyConfig is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *models.AlertConfiguration
func (_e *AlertmanagerMock_Expecter) ApplyConfig(_a0 interface{}, _a1 interface{}) *AlertmanagerMock_ApplyConfig_Call {
	return &AlertmanagerMock_ApplyConfig_Call{Call: _e.mock.On("ApplyConfig", _a0, _a1)}
}

func (_c *AlertmanagerMock_ApplyConfig_Call) Run(run func(_a0 context.Context, _a1 *models.AlertConfiguration)) *AlertmanagerMock_ApplyConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.AlertConfiguration))
	})
	return _c
}

func (_c *AlertmanagerMock_ApplyConfig_Call) Return(_a0 error) *AlertmanagerMock_ApplyConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertmanagerMock_ApplyConfig_Call) RunAndReturn(run func(context.Context, *models.AlertConfiguration) error) *AlertmanagerMock_ApplyConfig_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSilence provides a mock function with given fields: _a0, _a1
func (_m *AlertmanagerMock) CreateSilence(_a0 context.Context, _a1 *v2models.PostableSilence) (string, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateSilence")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v2models.PostableSilence) (string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v2models.PostableSilence) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v2models.PostableSilence) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AlertmanagerMock_CreateSilence_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSilence'
type AlertmanagerMock_CreateSilence_Call struct {
	*mock.Call
}

// CreateSilence is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *v2models.PostableSilence
func (_e *AlertmanagerMock_Expecter) CreateSilence(_a0 interface{}, _a1 interface{}) *AlertmanagerMock_CreateSilence_Call {
	return &AlertmanagerMock_CreateSilence_Call{Call: _e.mock.On("CreateSilence", _a0, _a1)}
}

func (_c *AlertmanagerMock_CreateSilence_Call) Run(run func(_a0 context.Context, _a1 *v2models.PostableSilence)) *AlertmanagerMock_CreateSilence_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v2models.PostableSilence))
	})
	return _c
}

func (_c *AlertmanagerMock_CreateSilence_Call) Return(_a0 string, _a1 error) *AlertmanagerMock_CreateSilence_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AlertmanagerMock_CreateSilence_Call) RunAndReturn(run func(context.Context, *v2models.PostableSilence) (string, error)) *AlertmanagerMock_CreateSilence_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteSilence provides a mock function with given fields: _a0, _a1
func (_m *AlertmanagerMock) DeleteSilence(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSilence")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AlertmanagerMock_DeleteSilence_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteSilence'
type AlertmanagerMock_DeleteSilence_Call struct {
	*mock.Call
}

// DeleteSilence is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *AlertmanagerMock_Expecter) DeleteSilence(_a0 interface{}, _a1 interface{}) *AlertmanagerMock_DeleteSilence_Call {
	return &AlertmanagerMock_DeleteSilence_Call{Call: _e.mock.On("DeleteSilence", _a0, _a1)}
}

func (_c *AlertmanagerMock_DeleteSilence_Call) Run(run func(_a0 context.Context, _a1 string)) *AlertmanagerMock_DeleteSilence_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AlertmanagerMock_DeleteSilence_Call) Return(_a0 error) *AlertmanagerMock_DeleteSilence_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertmanagerMock_DeleteSilence_Call) RunAndReturn(run func(context.Context, string) error) *AlertmanagerMock_DeleteSilence_Call {
	_c.Call.Return(run)
	return _c
}

// GetAlertGroups provides a mock function with given fields: ctx, active, silenced, inhibited, filter, receiver
func (_m *AlertmanagerMock) GetAlertGroups(ctx context.Context, active bool, silenced bool, inhibited bool, filter []string, receiver string) (v2models.AlertGroups, error) {
	ret := _m.Called(ctx, active, silenced, inhibited, filter, receiver)

	if len(ret) == 0 {
		panic("no return value specified for GetAlertGroups")
	}

	var r0 v2models.AlertGroups
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool, bool, bool, []string, string) (v2models.AlertGroups, error)); ok {
		return rf(ctx, active, silenced, inhibited, filter, receiver)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool, bool, bool, []string, string) v2models.AlertGroups); ok {
		r0 = rf(ctx, active, silenced, inhibited, filter, receiver)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v2models.AlertGroups)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool, bool, bool, []string, string) error); ok {
		r1 = rf(ctx, active, silenced, inhibited, filter, receiver)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AlertmanagerMock_GetAlertGroups_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAlertGroups'
type AlertmanagerMock_GetAlertGroups_Call struct {
	*mock.Call
}

// GetAlertGroups is a helper method to define mock.On call
//   - ctx context.Context
//   - active bool
//   - silenced bool
//   - inhibited bool
//   - filter []string
//   - receiver string
func (_e *AlertmanagerMock_Expecter) GetAlertGroups(ctx interface{}, active interface{}, silenced interface{}, inhibited interface{}, filter interface{}, receiver interface{}) *AlertmanagerMock_GetAlertGroups_Call {
	return &AlertmanagerMock_GetAlertGroups_Call{Call: _e.mock.On("GetAlertGroups", ctx, active, silenced, inhibited, filter, receiver)}
}

func (_c *AlertmanagerMock_GetAlertGroups_Call) Run(run func(ctx context.Context, active bool, silenced bool, inhibited bool, filter []string, receiver string)) *AlertmanagerMock_GetAlertGroups_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(bool), args[2].(bool), args[3].(bool), args[4].([]string), args[5].(string))
	})
	return _c
}

func (_c *AlertmanagerMock_GetAlertGroups_Call) Return(_a0 v2models.AlertGroups, _a1 error) *AlertmanagerMock_GetAlertGroups_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AlertmanagerMock_GetAlertGroups_Call) RunAndReturn(run func(context.Context, bool, bool, bool, []string, string) (v2models.AlertGroups, error)) *AlertmanagerMock_GetAlertGroups_Call {
	_c.Call.Return(run)
	return _c
}

// GetAlerts provides a mock function with given fields: ctx, active, silenced, inhibited, filter, receiver
func (_m *AlertmanagerMock) GetAlerts(ctx context.Context, active bool, silenced bool, inhibited bool, filter []string, receiver string) (v2models.GettableAlerts, error) {
	ret := _m.Called(ctx, active, silenced, inhibited, filter, receiver)

	if len(ret) == 0 {
		panic("no return value specified for GetAlerts")
	}

	var r0 v2models.GettableAlerts
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool, bool, bool, []string, string) (v2models.GettableAlerts, error)); ok {
		return rf(ctx, active, silenced, inhibited, filter, receiver)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool, bool, bool, []string, string) v2models.GettableAlerts); ok {
		r0 = rf(ctx, active, silenced, inhibited, filter, receiver)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v2models.GettableAlerts)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool, bool, bool, []string, string) error); ok {
		r1 = rf(ctx, active, silenced, inhibited, filter, receiver)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AlertmanagerMock_GetAlerts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAlerts'
type AlertmanagerMock_GetAlerts_Call struct {
	*mock.Call
}

// GetAlerts is a helper method to define mock.On call
//   - ctx context.Context
//   - active bool
//   - silenced bool
//   - inhibited bool
//   - filter []string
//   - receiver string
func (_e *AlertmanagerMock_Expecter) GetAlerts(ctx interface{}, active interface{}, silenced interface{}, inhibited interface{}, filter interface{}, receiver interface{}) *AlertmanagerMock_GetAlerts_Call {
	return &AlertmanagerMock_GetAlerts_Call{Call: _e.mock.On("GetAlerts", ctx, active, silenced, inhibited, filter, receiver)}
}

func (_c *AlertmanagerMock_GetAlerts_Call) Run(run func(ctx context.Context, active bool, silenced bool, inhibited bool, filter []string, receiver string)) *AlertmanagerMock_GetAlerts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(bool), args[2].(bool), args[3].(bool), args[4].([]string), args[5].(string))
	})
	return _c
}

func (_c *AlertmanagerMock_GetAlerts_Call) Return(_a0 v2models.GettableAlerts, _a1 error) *AlertmanagerMock_GetAlerts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AlertmanagerMock_GetAlerts_Call) RunAndReturn(run func(context.Context, bool, bool, bool, []string, string) (v2models.GettableAlerts, error)) *AlertmanagerMock_GetAlerts_Call {
	_c.Call.Return(run)
	return _c
}

// GetReceivers provides a mock function with given fields: ctx
func (_m *AlertmanagerMock) GetReceivers(ctx context.Context) ([]v2models.Receiver, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetReceivers")
	}

	var r0 []v2models.Receiver
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]v2models.Receiver, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []v2models.Receiver); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v2models.Receiver)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AlertmanagerMock_GetReceivers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReceivers'
type AlertmanagerMock_GetReceivers_Call struct {
	*mock.Call
}

// GetReceivers is a helper method to define mock.On call
//   - ctx context.Context
func (_e *AlertmanagerMock_Expecter) GetReceivers(ctx interface{}) *AlertmanagerMock_GetReceivers_Call {
	return &AlertmanagerMock_GetReceivers_Call{Call: _e.mock.On("GetReceivers", ctx)}
}

func (_c *AlertmanagerMock_GetReceivers_Call) Run(run func(ctx context.Context)) *AlertmanagerMock_GetReceivers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *AlertmanagerMock_GetReceivers_Call) Return(_a0 []v2models.Receiver, _a1 error) *AlertmanagerMock_GetReceivers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AlertmanagerMock_GetReceivers_Call) RunAndReturn(run func(context.Context) ([]v2models.Receiver, error)) *AlertmanagerMock_GetReceivers_Call {
	_c.Call.Return(run)
	return _c
}

// GetSilence provides a mock function with given fields: _a0, _a1
func (_m *AlertmanagerMock) GetSilence(_a0 context.Context, _a1 string) (v2models.GettableSilence, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetSilence")
	}

	var r0 v2models.GettableSilence
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (v2models.GettableSilence, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) v2models.GettableSilence); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(v2models.GettableSilence)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AlertmanagerMock_GetSilence_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSilence'
type AlertmanagerMock_GetSilence_Call struct {
	*mock.Call
}

// GetSilence is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *AlertmanagerMock_Expecter) GetSilence(_a0 interface{}, _a1 interface{}) *AlertmanagerMock_GetSilence_Call {
	return &AlertmanagerMock_GetSilence_Call{Call: _e.mock.On("GetSilence", _a0, _a1)}
}

func (_c *AlertmanagerMock_GetSilence_Call) Run(run func(_a0 context.Context, _a1 string)) *AlertmanagerMock_GetSilence_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AlertmanagerMock_GetSilence_Call) Return(_a0 v2models.GettableSilence, _a1 error) *AlertmanagerMock_GetSilence_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AlertmanagerMock_GetSilence_Call) RunAndReturn(run func(context.Context, string) (v2models.GettableSilence, error)) *AlertmanagerMock_GetSilence_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatus provides a mock function with given fields:
func (_m *AlertmanagerMock) GetStatus() definitions.GettableStatus {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStatus")
	}

	var r0 definitions.GettableStatus
	if rf, ok := ret.Get(0).(func() definitions.GettableStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(definitions.GettableStatus)
	}

	return r0
}

// AlertmanagerMock_GetStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatus'
type AlertmanagerMock_GetStatus_Call struct {
	*mock.Call
}

// GetStatus is a helper method to define mock.On call
func (_e *AlertmanagerMock_Expecter) GetStatus() *AlertmanagerMock_GetStatus_Call {
	return &AlertmanagerMock_GetStatus_Call{Call: _e.mock.On("GetStatus")}
}

func (_c *AlertmanagerMock_GetStatus_Call) Run(run func()) *AlertmanagerMock_GetStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AlertmanagerMock_GetStatus_Call) Return(_a0 definitions.GettableStatus) *AlertmanagerMock_GetStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertmanagerMock_GetStatus_Call) RunAndReturn(run func() definitions.GettableStatus) *AlertmanagerMock_GetStatus_Call {
	_c.Call.Return(run)
	return _c
}

// ListSilences provides a mock function with given fields: _a0, _a1
func (_m *AlertmanagerMock) ListSilences(_a0 context.Context, _a1 []string) (v2models.GettableSilences, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ListSilences")
	}

	var r0 v2models.GettableSilences
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) (v2models.GettableSilences, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) v2models.GettableSilences); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v2models.GettableSilences)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AlertmanagerMock_ListSilences_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSilences'
type AlertmanagerMock_ListSilences_Call struct {
	*mock.Call
}

// ListSilences is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []string
func (_e *AlertmanagerMock_Expecter) ListSilences(_a0 interface{}, _a1 interface{}) *AlertmanagerMock_ListSilences_Call {
	return &AlertmanagerMock_ListSilences_Call{Call: _e.mock.On("ListSilences", _a0, _a1)}
}

func (_c *AlertmanagerMock_ListSilences_Call) Run(run func(_a0 context.Context, _a1 []string)) *AlertmanagerMock_ListSilences_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *AlertmanagerMock_ListSilences_Call) Return(_a0 v2models.GettableSilences, _a1 error) *AlertmanagerMock_ListSilences_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AlertmanagerMock_ListSilences_Call) RunAndReturn(run func(context.Context, []string) (v2models.GettableSilences, error)) *AlertmanagerMock_ListSilences_Call {
	_c.Call.Return(run)
	return _c
}

// PutAlerts provides a mock function with given fields: _a0, _a1
func (_m *AlertmanagerMock) PutAlerts(_a0 context.Context, _a1 definitions.PostableAlerts) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for PutAlerts")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, definitions.PostableAlerts) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AlertmanagerMock_PutAlerts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutAlerts'
type AlertmanagerMock_PutAlerts_Call struct {
	*mock.Call
}

// PutAlerts is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 definitions.PostableAlerts
func (_e *AlertmanagerMock_Expecter) PutAlerts(_a0 interface{}, _a1 interface{}) *AlertmanagerMock_PutAlerts_Call {
	return &AlertmanagerMock_PutAlerts_Call{Call: _e.mock.On("PutAlerts", _a0, _a1)}
}

func (_c *AlertmanagerMock_PutAlerts_Call) Run(run func(_a0 context.Context, _a1 definitions.PostableAlerts)) *AlertmanagerMock_PutAlerts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(definitions.PostableAlerts))
	})
	return _c
}

func (_c *AlertmanagerMock_PutAlerts_Call) Return(_a0 error) *AlertmanagerMock_PutAlerts_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertmanagerMock_PutAlerts_Call) RunAndReturn(run func(context.Context, definitions.PostableAlerts) error) *AlertmanagerMock_PutAlerts_Call {
	_c.Call.Return(run)
	return _c
}

// Ready provides a mock function with given fields:
func (_m *AlertmanagerMock) Ready() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// AlertmanagerMock_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type AlertmanagerMock_Ready_Call struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
func (_e *AlertmanagerMock_Expecter) Ready() *AlertmanagerMock_Ready_Call {
	return &AlertmanagerMock_Ready_Call{Call: _e.mock.On("Ready")}
}

func (_c *AlertmanagerMock_Ready_Call) Run(run func()) *AlertmanagerMock_Ready_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AlertmanagerMock_Ready_Call) Return(_a0 bool) *AlertmanagerMock_Ready_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertmanagerMock_Ready_Call) RunAndReturn(run func() bool) *AlertmanagerMock_Ready_Call {
	_c.Call.Return(run)
	return _c
}

// SaveAndApplyConfig provides a mock function with given fields: ctx, config
func (_m *AlertmanagerMock) SaveAndApplyConfig(ctx context.Context, config *definitions.PostableUserConfig) error {
	ret := _m.Called(ctx, config)

	if len(ret) == 0 {
		panic("no return value specified for SaveAndApplyConfig")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *definitions.PostableUserConfig) error); ok {
		r0 = rf(ctx, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AlertmanagerMock_SaveAndApplyConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveAndApplyConfig'
type AlertmanagerMock_SaveAndApplyConfig_Call struct {
	*mock.Call
}

// SaveAndApplyConfig is a helper method to define mock.On call
//   - ctx context.Context
//   - config *definitions.PostableUserConfig
func (_e *AlertmanagerMock_Expecter) SaveAndApplyConfig(ctx interface{}, config interface{}) *AlertmanagerMock_SaveAndApplyConfig_Call {
	return &AlertmanagerMock_SaveAndApplyConfig_Call{Call: _e.mock.On("SaveAndApplyConfig", ctx, config)}
}

func (_c *AlertmanagerMock_SaveAndApplyConfig_Call) Run(run func(ctx context.Context, config *definitions.PostableUserConfig)) *AlertmanagerMock_SaveAndApplyConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*definitions.PostableUserConfig))
	})
	return _c
}

func (_c *AlertmanagerMock_SaveAndApplyConfig_Call) Return(_a0 error) *AlertmanagerMock_SaveAndApplyConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertmanagerMock_SaveAndApplyConfig_Call) RunAndReturn(run func(context.Context, *definitions.PostableUserConfig) error) *AlertmanagerMock_SaveAndApplyConfig_Call {
	_c.Call.Return(run)
	return _c
}

// SaveAndApplyDefaultConfig provides a mock function with given fields: ctx
func (_m *AlertmanagerMock) SaveAndApplyDefaultConfig(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SaveAndApplyDefaultConfig")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AlertmanagerMock_SaveAndApplyDefaultConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveAndApplyDefaultConfig'
type AlertmanagerMock_SaveAndApplyDefaultConfig_Call struct {
	*mock.Call
}

// SaveAndApplyDefaultConfig is a helper method to define mock.On call
//   - ctx context.Context
func (_e *AlertmanagerMock_Expecter) SaveAndApplyDefaultConfig(ctx interface{}) *AlertmanagerMock_SaveAndApplyDefaultConfig_Call {
	return &AlertmanagerMock_SaveAndApplyDefaultConfig_Call{Call: _e.mock.On("SaveAndApplyDefaultConfig", ctx)}
}

func (_c *AlertmanagerMock_SaveAndApplyDefaultConfig_Call) Run(run func(ctx context.Context)) *AlertmanagerMock_SaveAndApplyDefaultConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *AlertmanagerMock_SaveAndApplyDefaultConfig_Call) Return(_a0 error) *AlertmanagerMock_SaveAndApplyDefaultConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertmanagerMock_SaveAndApplyDefaultConfig_Call) RunAndReturn(run func(context.Context) error) *AlertmanagerMock_SaveAndApplyDefaultConfig_Call {
	_c.Call.Return(run)
	return _c
}

// StopAndWait provides a mock function with given fields:
func (_m *AlertmanagerMock) StopAndWait() {
	_m.Called()
}

// AlertmanagerMock_StopAndWait_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StopAndWait'
type AlertmanagerMock_StopAndWait_Call struct {
	*mock.Call
}

// StopAndWait is a helper method to define mock.On call
func (_e *AlertmanagerMock_Expecter) StopAndWait() *AlertmanagerMock_StopAndWait_Call {
	return &AlertmanagerMock_StopAndWait_Call{Call: _e.mock.On("StopAndWait")}
}

func (_c *AlertmanagerMock_StopAndWait_Call) Run(run func()) *AlertmanagerMock_StopAndWait_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AlertmanagerMock_StopAndWait_Call) Return() *AlertmanagerMock_StopAndWait_Call {
	_c.Call.Return()
	return _c
}

func (_c *AlertmanagerMock_StopAndWait_Call) RunAndReturn(run func()) *AlertmanagerMock_StopAndWait_Call {
	_c.Call.Return(run)
	return _c
}

// TestReceivers provides a mock function with given fields: ctx, c
func (_m *AlertmanagerMock) TestReceivers(ctx context.Context, c definitions.TestReceiversConfigBodyParams) (*notifier.TestReceiversResult, error) {
	ret := _m.Called(ctx, c)

	if len(ret) == 0 {
		panic("no return value specified for TestReceivers")
	}

	var r0 *notifier.TestReceiversResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, definitions.TestReceiversConfigBodyParams) (*notifier.TestReceiversResult, error)); ok {
		return rf(ctx, c)
	}
	if rf, ok := ret.Get(0).(func(context.Context, definitions.TestReceiversConfigBodyParams) *notifier.TestReceiversResult); ok {
		r0 = rf(ctx, c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notifier.TestReceiversResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, definitions.TestReceiversConfigBodyParams) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AlertmanagerMock_TestReceivers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TestReceivers'
type AlertmanagerMock_TestReceivers_Call struct {
	*mock.Call
}

// TestReceivers is a helper method to define mock.On call
//   - ctx context.Context
//   - c definitions.TestReceiversConfigBodyParams
func (_e *AlertmanagerMock_Expecter) TestReceivers(ctx interface{}, c interface{}) *AlertmanagerMock_TestReceivers_Call {
	return &AlertmanagerMock_TestReceivers_Call{Call: _e.mock.On("TestReceivers", ctx, c)}
}

func (_c *AlertmanagerMock_TestReceivers_Call) Run(run func(ctx context.Context, c definitions.TestReceiversConfigBodyParams)) *AlertmanagerMock_TestReceivers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(definitions.TestReceiversConfigBodyParams))
	})
	return _c
}

func (_c *AlertmanagerMock_TestReceivers_Call) Return(_a0 *notifier.TestReceiversResult, _a1 error) *AlertmanagerMock_TestReceivers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AlertmanagerMock_TestReceivers_Call) RunAndReturn(run func(context.Context, definitions.TestReceiversConfigBodyParams) (*notifier.TestReceiversResult, error)) *AlertmanagerMock_TestReceivers_Call {
	_c.Call.Return(run)
	return _c
}

// TestTemplate provides a mock function with given fields: ctx, c
func (_m *AlertmanagerMock) TestTemplate(ctx context.Context, c definitions.TestTemplatesConfigBodyParams) (*notify.TestTemplatesResults, error) {
	ret := _m.Called(ctx, c)

	if len(ret) == 0 {
		panic("no return value specified for TestTemplate")
	}

	var r0 *notify.TestTemplatesResults
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, definitions.TestTemplatesConfigBodyParams) (*notify.TestTemplatesResults, error)); ok {
		return rf(ctx, c)
	}
	if rf, ok := ret.Get(0).(func(context.Context, definitions.TestTemplatesConfigBodyParams) *notify.TestTemplatesResults); ok {
		r0 = rf(ctx, c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notify.TestTemplatesResults)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, definitions.TestTemplatesConfigBodyParams) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AlertmanagerMock_TestTemplate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TestTemplate'
type AlertmanagerMock_TestTemplate_Call struct {
	*mock.Call
}

// TestTemplate is a helper method to define mock.On call
//   - ctx context.Context
//   - c definitions.TestTemplatesConfigBodyParams
func (_e *AlertmanagerMock_Expecter) TestTemplate(ctx interface{}, c interface{}) *AlertmanagerMock_TestTemplate_Call {
	return &AlertmanagerMock_TestTemplate_Call{Call: _e.mock.On("TestTemplate", ctx, c)}
}

func (_c *AlertmanagerMock_TestTemplate_Call) Run(run func(ctx context.Context, c definitions.TestTemplatesConfigBodyParams)) *AlertmanagerMock_TestTemplate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(definitions.TestTemplatesConfigBodyParams))
	})
	return _c
}

func (_c *AlertmanagerMock_TestTemplate_Call) Return(_a0 *notify.TestTemplatesResults, _a1 error) *AlertmanagerMock_TestTemplate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AlertmanagerMock_TestTemplate_Call) RunAndReturn(run func(context.Context, definitions.TestTemplatesConfigBodyParams) (*notify.TestTemplatesResults, error)) *AlertmanagerMock_TestTemplate_Call {
	_c.Call.Return(run)
	return _c
}

// NewAlertmanagerMock creates a new instance of AlertmanagerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAlertmanagerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *AlertmanagerMock {
	mock := &AlertmanagerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
