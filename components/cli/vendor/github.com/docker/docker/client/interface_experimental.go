package client // import "github.com/docker/docker/client"

import (
	"context"

	"github.com/docker/docker/api/types"
)

type apiClientExperimental interface {
	MigrationAPIClient
	CheckpointAPIClient
}

// MigrationAPIClient defines API client methods for the migration
type MigrationAPIClient interface {
	ContainerMigrate(ctx context.Context, container string, options types.MigrationOptions) error
}

// CheckpointAPIClient defines API client methods for the checkpoints
type CheckpointAPIClient interface {
	CheckpointCreate(ctx context.Context, container string, options types.CheckpointCreateOptions) error
	CheckpointDelete(ctx context.Context, container string, options types.CheckpointDeleteOptions) error
	CheckpointList(ctx context.Context, container string, options types.CheckpointListOptions) ([]types.Checkpoint, error)
}
