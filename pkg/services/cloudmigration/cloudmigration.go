package cloudmigration

import (
	"context"

	"github.com/grafana/grafana/pkg/services/cloudmigration/cmsclient"
)

type Service interface {
	CreateToken(context.Context) (CreateAccessTokenResponse, error)
	ValidateToken(context.Context, Base64EncodedTokenPayload) error

	CreateMigration(context.Context, CloudMigrationRequest) (*CloudMigrationResponse, error)
	GetMigration(context.Context, int64) (*CloudMigration, error)
	DeleteMigration(context.Context, int64) (*CloudMigration, error)
	UpdateMigration(context.Context, int64, CloudMigrationRequest) (*CloudMigrationResponse, error)
	GetMigrationList(context.Context) ([]CloudMigrationResponse, error)

	RunMigration(context.Context, int64) (*cmsclient.MigrateDataResponseDTO, error)
	SaveMigrationRun(context.Context, *CloudMigrationRun) (int64, error)
	GetMigrationStatus(context.Context, string, string) (*CloudMigrationRun, error)
	GetMigrationRunList(context.Context, string) (*CloudMigrationRunList, error)
}
