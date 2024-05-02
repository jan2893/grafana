package cloudmigrationimpl

import (
	"context"

	"github.com/grafana/grafana/pkg/services/cloudmigration"
	"github.com/grafana/grafana/pkg/services/cloudmigration/cmsclient"
)

// NoopServiceImpl Define the Service Implementation.
type NoopServiceImpl struct{}

var _ cloudmigration.Service = (*NoopServiceImpl)(nil)

func (s *NoopServiceImpl) CreateToken(ctx context.Context) (cloudmigration.CreateAccessTokenResponse, error) {
	return cloudmigration.CreateAccessTokenResponse{}, cloudmigration.ErrFeatureDisabledError
}
func (s *NoopServiceImpl) ValidateToken(ctx context.Context, token cloudmigration.Base64EncodedTokenPayload) error {
	return cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) GetMigration(ctx context.Context, uid string) (*cloudmigration.CloudMigration, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) GetMigrationList(ctx context.Context) ([]cloudmigration.CloudMigrationResponse, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) CreateMigration(ctx context.Context, cm cloudmigration.CloudMigrationRequest) (*cloudmigration.CloudMigrationResponse, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) UpdateMigration(ctx context.Context, uid string, cm cloudmigration.CloudMigrationRequest) (*cloudmigration.CloudMigrationResponse, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) GetMigrationStatus(ctx context.Context, runUID string) (*cloudmigration.CloudMigrationRun, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) GetMigrationRunList(ctx context.Context, uid string) (*cloudmigration.CloudMigrationRunList, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) DeleteMigration(ctx context.Context, uid string) (*cloudmigration.CloudMigration, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) CreateMigrationRun(context.Context, cloudmigration.CloudMigrationRun) (string, error) {
	return "", cloudmigration.ErrInternalNotImplementedError
}

func (s *NoopServiceImpl) RunMigration(context.Context, string) (*cmsclient.MigrateDataResponseDTO, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}
