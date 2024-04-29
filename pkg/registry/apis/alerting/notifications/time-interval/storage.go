package time_interval

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"

	grafanaregistry "github.com/grafana/grafana/pkg/apiserver/registry/generic"
)

func newTimeIntervalStorage(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter, legacy *legacyStorage) (*genericregistry.Store, error) {
	strategy := grafanaregistry.NewStrategy(scheme)
	s := &genericregistry.Store{
		NewFunc:                   resourceInfo.NewFunc,
		NewListFunc:               resourceInfo.NewListFunc,
		PredicateFunc:             grafanaregistry.Matcher,
		DefaultQualifiedResource:  resourceInfo.GroupResource(),
		SingularQualifiedResource: resourceInfo.SingularGroupResource(),
		TableConvertor:            legacy.tableConverter,
		CreateStrategy:            strategy,
		UpdateStrategy:            strategy,
		DeleteStrategy:            strategy,
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: grafanaregistry.GetAttrs}
	if err := s.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return s, nil
}
