package cmsclient

import (
	"context"

	"github.com/grafana/grafana/pkg/util/errutil"
)

type ValidateKeyInput struct {
	AuthToken   string `json:"-"`
	StackID     int    `json:"stackID"`
	ClusterSlug string `json:"clusterSlug"`
}

type MigrateDataInput struct {
	AuthToken   string `json:"-"`
	StackID     int    `json:"stackID"`
	ClusterSlug string `json:"clusterSlug"`
}

// swagger:enum MigrateDataType
type MigrateDataType string

type MigrateDataRequestDTO struct {
	Items []MigrateDataRequestItemDTO `json:"items"`
}

type MigrateDataRequestItemDTO struct {
	Type  MigrateDataType `json:"type"`
	RefID string          `json:"refId"`
	Name  string          `json:"name"`
	Data  interface{}     `json:"data"`
}

type MigrateDataResponseDTO struct {
	RunID int64                        `json:"id"`
	Items []MigrateDataResponseItemDTO `json:"items"`
}

type MigrateDataResponseItemDTO struct {
	// required:true
	Type MigrateDataType `json:"type"`
	// required:true
	RefID string `json:"refId"`
	// required:true
	Status ItemStatus `json:"status"`
	Error  string     `json:"error,omitempty"`
}

// swagger:enum ItemStatus
type ItemStatus string

type Client interface {
	ValidateKey(context.Context, ValidateKeyInput) error
	MigrateData(context.Context, MigrateDataInput, MigrateDataRequestDTO) (*MigrateDataResponseDTO, error)
}

const logPrefix = "cloudmigration.cmsclient"

var ErrMigrationNotDeleted = errutil.Internal("cloudmigrations.developerModeEnabled", errutil.WithPublicMessage("Developer mode enabled"))
