package time_interval

import (
	"context"
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/kube-openapi/pkg/common"

	"github.com/grafana/grafana/pkg/apis/alerting/notifications"
	timeInterval "github.com/grafana/grafana/pkg/apis/alerting/notifications/time-intervals/v0alpha1"
	"github.com/grafana/grafana/pkg/apiserver/builder"
	grafanarest "github.com/grafana/grafana/pkg/apiserver/rest"
	"github.com/grafana/grafana/pkg/infra/appcontext"
	"github.com/grafana/grafana/pkg/services/accesscontrol"
	"github.com/grafana/grafana/pkg/services/apiserver/endpoints/request"
	"github.com/grafana/grafana/pkg/services/apiserver/utils"
	"github.com/grafana/grafana/pkg/services/ngalert"
	"github.com/grafana/grafana/pkg/services/ngalert/api/tooling/definitions"
	"github.com/grafana/grafana/pkg/setting"
)

var _ builder.APIGroupBuilder = (*TimeIntervalsAPIBuilder)(nil)

type TimeIntervalService interface {
	GetMuteTimings(ctx context.Context, orgID int64) ([]definitions.MuteTimeInterval, error)
	GetMuteTiming(ctx context.Context, name string, orgID int64) (definitions.MuteTimeInterval, error)
	CreateMuteTiming(ctx context.Context, mt definitions.MuteTimeInterval, orgID int64) (definitions.MuteTimeInterval, error)
	UpdateMuteTiming(ctx context.Context, mt definitions.MuteTimeInterval, orgID int64) (definitions.MuteTimeInterval, error)
	DeleteMuteTiming(ctx context.Context, name string, orgID int64) error
}

// This is used just so wire has something unique to return
type TimeIntervalsAPIBuilder struct {
	authz      accesscontrol.AccessControl
	service    TimeIntervalService
	namespacer request.NamespaceMapper
	gv         schema.GroupVersion
}

func RegisterAPIService(
	apiregistration builder.APIRegistrar,
	cfg *setting.Cfg,
	ng *ngalert.AlertNG,
) *TimeIntervalsAPIBuilder {
	if ng.IsDisabled() {
		return nil
	}

	builder := &TimeIntervalsAPIBuilder{
		service:    ng.Api.MuteTimings,
		namespacer: request.GetNamespaceMapper(cfg),
		gv:         timeInterval.TimeIntervalResourceInfo.GroupVersion(),
		authz:      ng.Api.AccessControl,
	}
	apiregistration.RegisterAPI(builder)
	return builder
}

func (t TimeIntervalsAPIBuilder) GetGroupVersion() schema.GroupVersion {
	return t.gv
}

func (t TimeIntervalsAPIBuilder) InstallSchema(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(t.gv,
		&timeInterval.TimeIntervals{},
		&timeInterval.TimeIntervalsList{},
	)
	metav1.AddToGroupVersion(scheme, t.gv)
	return scheme.SetVersionPriority(t.gv)
}

func (t TimeIntervalsAPIBuilder) GetAPIGroupInfo(scheme *runtime.Scheme, codecs serializer.CodecFactory, optsGetter generic.RESTOptionsGetter, dualWrite bool) (*genericapiserver.APIGroupInfo, error) {
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(timeInterval.GROUP, scheme, metav1.ParameterCodec, codecs)

	legacyStore := &legacyStorage{
		service:    t.service,
		namespacer: t.namespacer,
		tableConverter: utils.NewTableConverter(
			resourceInfo.GroupResource(),
			[]metav1.TableColumnDefinition{
				{Name: "Name", Type: "string", Format: "name"},
				// {Name: "Intervals", Type: "string", Format: "string", Description: "The display name"},
			},
			func(obj any) ([]interface{}, error) {
				r, ok := obj.(*timeInterval.TimeIntervals)
				if ok {
					return []interface{}{
						r.Name,
						// r.Spec, //TODO implement formatting for Spec, same as UI?
					}, nil
				}
				return nil, fmt.Errorf("expected resource or info")
			}),
	}

	storage := map[string]rest.Storage{
		resourceInfo.StoragePath(): legacyStore,
	}

	// enable dual writes if a RESTOptionsGetter is provided
	if dualWrite && optsGetter != nil {
		store, err := newTimeIntervalStorage(scheme, optsGetter, legacyStore)
		if err != nil {
			return nil, err
		}
		storage[resourceInfo.StoragePath()] = grafanarest.NewDualWriter(legacyStore, store)
	}

	apiGroupInfo.VersionedResourcesStorageMap[timeInterval.VERSION] = storage
	return &apiGroupInfo, nil
}

func (t TimeIntervalsAPIBuilder) GetOpenAPIDefinitions() common.GetOpenAPIDefinitions {
	// TODO Figure out better way
	return func(callback common.ReferenceCallback) map[string]common.OpenAPIDefinition {
		all := notifications.GetOpenAPIDefinitions(callback)
		for k := range all {
			if !strings.Contains(k, "github.com/grafana/grafana/pkg/apis/alerting/notifications/time-intervals") {
				delete(all, k)
			}
		}
		return all
	}
}

func (t TimeIntervalsAPIBuilder) GetAPIRoutes() *builder.APIRoutes {
	return nil
}

func (t TimeIntervalsAPIBuilder) GetAuthorizer() authorizer.Authorizer {
	return authorizer.AuthorizerFunc(
		func(ctx context.Context, attr authorizer.Attributes) (authorized authorizer.Decision, reason string, err error) {
			user, err := appcontext.User(ctx)
			if err != nil {
				return authorizer.DecisionDeny, "valid user is required", err
			}

			var action accesscontrol.Evaluator
			switch attr.GetVerb() {
			case "patch":
				fallthrough
			case "create":
				fallthrough
			case "update":
				action = accesscontrol.EvalAny(
					accesscontrol.EvalPermission(accesscontrol.ActionAlertingNotificationsTimeIntervalsWrite),
					accesscontrol.EvalPermission(accesscontrol.ActionAlertingNotificationsWrite),
				)
			case "deletecollection":
				fallthrough
			case "delete":
				action = accesscontrol.EvalAny(
					accesscontrol.EvalPermission(accesscontrol.ActionAlertingNotificationsTimeIntervalsDelete),
					accesscontrol.EvalPermission(accesscontrol.ActionAlertingNotificationsWrite),
				)
			}

			eval := accesscontrol.EvalAny(
				accesscontrol.EvalPermission(accesscontrol.ActionAlertingNotificationsTimeIntervalsRead),
				accesscontrol.EvalPermission(accesscontrol.ActionAlertingNotificationsRead),
			)
			if action != nil {
				eval = accesscontrol.EvalAll(eval, action)
			}

			ok, err := t.authz.Evaluate(ctx, user, eval)
			if ok {
				return authorizer.DecisionAllow, "", nil
			}
			return authorizer.DecisionDeny, "time-interval", err
		})
}
