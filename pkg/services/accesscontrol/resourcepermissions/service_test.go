package resourcepermissions

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/grafana/grafana/pkg/api/routing"
	"github.com/grafana/grafana/pkg/infra/db"
	"github.com/grafana/grafana/pkg/infra/tracing"
	"github.com/grafana/grafana/pkg/services/accesscontrol"
	"github.com/grafana/grafana/pkg/services/accesscontrol/acimpl"
	"github.com/grafana/grafana/pkg/services/accesscontrol/actest"
	"github.com/grafana/grafana/pkg/services/featuremgmt"
	"github.com/grafana/grafana/pkg/services/licensing/licensingtest"
	"github.com/grafana/grafana/pkg/services/org/orgimpl"
	"github.com/grafana/grafana/pkg/services/quota/quotatest"
	"github.com/grafana/grafana/pkg/services/supportbundles/supportbundlestest"
	"github.com/grafana/grafana/pkg/services/team"
	"github.com/grafana/grafana/pkg/services/team/teamimpl"
	"github.com/grafana/grafana/pkg/services/user"
	"github.com/grafana/grafana/pkg/services/user/userimpl"
	"github.com/grafana/grafana/pkg/setting"
)

type setUserPermissionTest struct {
	desc     string
	callHook bool
}

func TestService_SetUserPermission(t *testing.T) {
	tests := []setUserPermissionTest{
		{
			desc:     "should call hook when updating user permissions",
			callHook: true,
		},
		{
			desc:     "should not call hook when updating user permissions",
			callHook: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			service, usrSvc, _ := setupTestEnvironment(t, Options{
				Resource:             "dashboards",
				Assignments:          Assignments{Users: true},
				PermissionsToActions: nil,
			})

			// seed user
			user, err := usrSvc.Create(context.Background(), &user.CreateUserCommand{Login: "test", OrgID: 1})
			require.NoError(t, err)

			var hookCalled bool
			if tt.callHook {
				service.options.OnSetUser = func(session *db.Session, orgID int64, user accesscontrol.User, resourceID, permission string) error {
					hookCalled = true
					return nil
				}
			}

			_, err = service.SetUserPermission(context.Background(), user.OrgID, accesscontrol.User{ID: user.ID}, "1", "")
			require.NoError(t, err)
			assert.Equal(t, tt.callHook, hookCalled)
		})
	}
}

type setTeamPermissionTest struct {
	desc     string
	callHook bool
}

func TestService_SetTeamPermission(t *testing.T) {
	tests := []setTeamPermissionTest{
		{
			desc:     "should call hook when updating user permissions",
			callHook: true,
		},
		{
			desc:     "should not call hook when updating user permissions",
			callHook: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			service, _, teamSvc := setupTestEnvironment(t, Options{
				Resource:             "dashboards",
				Assignments:          Assignments{Teams: true},
				PermissionsToActions: nil,
			})

			// seed team
			team, err := teamSvc.CreateTeam(context.Background(), "test", "test@test.com", 1)
			require.NoError(t, err)

			var hookCalled bool
			if tt.callHook {
				service.options.OnSetTeam = func(session *db.Session, orgID, teamID int64, resourceID, permission string) error {
					hookCalled = true
					return nil
				}
			}

			_, err = service.SetTeamPermission(context.Background(), team.OrgID, team.ID, "1", "")
			require.NoError(t, err)
			assert.Equal(t, tt.callHook, hookCalled)
		})
	}
}

type setBuiltInRolePermissionTest struct {
	desc     string
	callHook bool
}

func TestService_SetBuiltInRolePermission(t *testing.T) {
	tests := []setBuiltInRolePermissionTest{
		{
			desc:     "should call hook when updating user permissions",
			callHook: true,
		},
		{
			desc:     "should not call hook when updating user permissions",
			callHook: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			service, _, _ := setupTestEnvironment(t, Options{
				Resource:             "dashboards",
				Assignments:          Assignments{BuiltInRoles: true},
				PermissionsToActions: nil,
			})

			var hookCalled bool
			if tt.callHook {
				service.options.OnSetBuiltInRole = func(session *db.Session, orgID int64, builtInRole, resourceID, permission string) error {
					hookCalled = true
					return nil
				}
			}

			_, err := service.SetBuiltInRolePermission(context.Background(), 1, "Viewer", "1", "")
			require.NoError(t, err)
			assert.Equal(t, tt.callHook, hookCalled)
		})
	}
}

type setPermissionsTest struct {
	desc      string
	options   Options
	commands  []accesscontrol.SetResourcePermissionCommand
	expectErr bool
}

func TestService_SetPermissions(t *testing.T) {
	tests := []setPermissionsTest{
		{
			desc: "should set all permissions",
			options: Options{
				Resource: "dashboards",
				Assignments: Assignments{
					Users:        true,
					Teams:        true,
					BuiltInRoles: true,
				},
				PermissionsToActions: map[string][]string{
					"View": {"dashboards:read"},
				},
			},
			commands: []accesscontrol.SetResourcePermissionCommand{
				{UserID: 1, Permission: "View"},
				{TeamID: 1, Permission: "View"},
				{BuiltinRole: "Editor", Permission: "View"},
			},
		},
		{
			desc: "should return error for invalid permission",
			options: Options{
				Resource: "dashboards",
				Assignments: Assignments{
					Users:        true,
					Teams:        true,
					BuiltInRoles: true,
				},
				PermissionsToActions: map[string][]string{
					"View": {"dashboards:read"},
				},
			},
			commands: []accesscontrol.SetResourcePermissionCommand{
				{UserID: 1, Permission: "View"},
				{TeamID: 1, Permission: "View"},
				{BuiltinRole: "Editor", Permission: "Not real permission"},
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			service, usrSvc, teamSvc := setupTestEnvironment(t, tt.options)

			// seed user
			_, err := usrSvc.Create(context.Background(), &user.CreateUserCommand{Login: "user", OrgID: 1})
			require.NoError(t, err)
			_, err = teamSvc.CreateTeam(context.Background(), "team", "", 1)
			require.NoError(t, err)

			permissions, err := service.SetPermissions(context.Background(), 1, "1", tt.commands...)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, permissions, len(tt.commands))
			}
		})
	}
}

func TestService_RegisterActionSets(t *testing.T) {
	type registerActionSetsTest struct {
		desc               string
		actionSetsEnabled  bool
		options            Options
		expectedActionSets []ActionSet
	}

	tests := []registerActionSetsTest{
		{
			desc:              "should register folder action sets if action sets are enabled",
			actionSetsEnabled: true,
			options: Options{
				Resource: "folders",
				PermissionsToActions: map[string][]string{
					"View": {"folders:read", "dashboards:read"},
					"Edit": {"folders:read", "dashboards:read", "folders:write", "dashboards:write"},
				},
			},
			expectedActionSets: []ActionSet{
				{
					Action:  "folders:view",
					Actions: []string{"folders:read", "dashboards:read"},
				},
				{
					Action:  "folders:edit",
					Actions: []string{"folders:read", "dashboards:read", "folders:write", "dashboards:write"},
				},
			},
		},
		{
			desc:              "should register dashboard action set if action sets are enabled",
			actionSetsEnabled: true,
			options: Options{
				Resource: "dashboards",
				PermissionsToActions: map[string][]string{
					"View": {"dashboards:read"},
				},
			},
			expectedActionSets: []ActionSet{
				{
					Action:  "dashboards:view",
					Actions: []string{"dashboards:read"},
				},
			},
		},
		{
			desc:              "should not register dashboard action set if action sets are not enabled",
			actionSetsEnabled: false,
			options: Options{
				Resource: "dashboards",
				PermissionsToActions: map[string][]string{
					"View": {"dashboards:read"},
				},
			},
			expectedActionSets: []ActionSet{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			cfg := setting.NewCfg()
			cfg.IsFeatureToggleEnabled = func(ft string) bool {
				if ft == featuremgmt.FlagAccessActionSets {
					return tt.actionSetsEnabled
				}
				return false
			}
			ac := acimpl.ProvideAccessControl(cfg)
			actionSets := NewActionSetService(ac)
			features := featuremgmt.WithFeatures()
			if tt.actionSetsEnabled {
				features = featuremgmt.WithFeatures(featuremgmt.FlagAccessActionSets)
			}
			_, err := New(
				setting.NewCfg(), tt.options, features, routing.NewRouteRegister(), licensingtest.NewFakeLicensing(),
				ac, &actest.FakeService{}, db.InitTestDB(t), nil, nil, actionSets,
			)
			require.NoError(t, err)

			if len(tt.expectedActionSets) > 0 {
				for _, expectedActionSet := range tt.expectedActionSets {
					actionSet := actionSets.GetActionSet(expectedActionSet.Action)
					assert.ElementsMatch(t, expectedActionSet.Actions, actionSet)
				}
			} else {
				// Check that action sets have not been registered
				for permission := range tt.options.PermissionsToActions {
					actionSetName := actionSets.GetActionSetName(tt.options.Resource, permission)
					assert.Nil(t, actionSets.GetActionSet(actionSetName))
				}
			}
		})
	}
}

func setupTestEnvironment(t *testing.T, ops Options) (*Service, user.Service, team.Service) {
	t.Helper()

	sql := db.InitTestDB(t)
	cfg := setting.NewCfg()
	tracer := tracing.InitializeTracerForTest()

	teamSvc, err := teamimpl.ProvideService(sql, cfg, tracer)
	require.NoError(t, err)

	orgSvc, err := orgimpl.ProvideService(sql, cfg, quotatest.New(false, nil))
	require.NoError(t, err)

	userSvc, err := userimpl.ProvideService(
		sql, orgSvc, cfg, teamSvc, nil, tracer,
		quotatest.New(false, nil), supportbundlestest.NewFakeBundleService(),
	)
	require.NoError(t, err)

	license := licensingtest.NewFakeLicensing()
	license.On("FeatureEnabled", "accesscontrol.enforcement").Return(true).Maybe()
	ac := acimpl.ProvideAccessControl(setting.NewCfg())
	acService := &actest.FakeService{}
	service, err := New(
		cfg, ops, featuremgmt.WithFeatures(), routing.NewRouteRegister(), license,
		ac, acService, sql, teamSvc, userSvc, NewActionSetService(ac),
	)
	require.NoError(t, err)

	return service, userSvc, teamSvc
}
