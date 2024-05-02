package cloudmigration

import (
	"context"

	"github.com/grafana/grafana/pkg/services/cloudmigration/cmsclient"
)

type Service interface {
	CreateToken(context.Context) (CreateAccessTokenResponse, error)
	ValidateToken(context.Context, Base64EncodedTokenPayload) error

	CreateMigration(context.Context, CloudMigrationRequest) (*CloudMigrationResponse, error)
	GetMigration(ctx context.Context, uid string) (*CloudMigration, error)
	DeleteMigration(ctx context.Context, uid string) (*CloudMigration, error)
	UpdateMigration(ctx context.Context, uid string, request CloudMigrationRequest) (*CloudMigrationResponse, error)
	GetMigrationList(context.Context) ([]CloudMigrationResponse, error)

	RunMigration(ctx context.Context, uid string) (*cmsclient.MigrateDataResponseDTO, error)
	CreateMigrationRun(context.Context, CloudMigrationRun) (string, error)
	GetMigrationStatus(ctx context.Context, runUID string) (*CloudMigrationRun, error)
	GetMigrationRunList(context.Context, string) (*CloudMigrationRunList, error)
}
