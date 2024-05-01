package cloudmigrationimpl

import (
	"context"

	"github.com/grafana/grafana/pkg/services/cloudmigration"
	"github.com/grafana/grafana/pkg/services/cloudmigration/cmsclient"
)

// CloudMigrationsServiceImpl Define the Service Implementation.
type NoopServiceImpl struct{}

var _ cloudmigration.Service = (*NoopServiceImpl)(nil)

func (s *NoopServiceImpl) CreateToken(ctx context.Context) (cloudmigration.CreateAccessTokenResponse, error) {
	return cloudmigration.CreateAccessTokenResponse{}, cloudmigration.ErrFeatureDisabledError
}
func (s *NoopServiceImpl) ValidateToken(ctx context.Context, token cloudmigration.Base64EncodedTokenPayload) error {
	return cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) GetMigration(ctx context.Context, id int64) (*cloudmigration.CloudMigration, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) GetMigrationList(ctx context.Context) ([]cloudmigration.CloudMigrationResponse, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) CreateMigration(ctx context.Context, cm cloudmigration.CloudMigrationRequest) (*cloudmigration.CloudMigrationResponse, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) UpdateMigration(ctx context.Context, id int64, cm cloudmigration.CloudMigrationRequest) (*cloudmigration.CloudMigrationResponse, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) GetMigrationStatus(ctx context.Context, id string, runID string) (*cloudmigration.CloudMigrationRun, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) GetMigrationRunList(ctx context.Context, id string) (*cloudmigration.CloudMigrationRunList, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) DeleteMigration(ctx context.Context, id int64) (*cloudmigration.CloudMigration, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}

func (s *NoopServiceImpl) SaveMigrationRun(ctx context.Context, cmr *cloudmigration.CloudMigrationRun) (int64, error) {
	return -1, cloudmigration.ErrInternalNotImplementedError
}

func (s *NoopServiceImpl) RunMigration(context.Context, int64) (*cmsclient.MigrateDataResponseDTO, error) {
	return nil, cloudmigration.ErrFeatureDisabledError
}
