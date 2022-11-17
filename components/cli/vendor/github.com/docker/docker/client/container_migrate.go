package client // import "github.com/docker/docker/client"

import (
	"context"

	"github.com/docker/docker/api/types"
)

// ContainerMigrate migrates the given container to the given target address
func (cli *Client) ContainerMigrate(ctx context.Context, container string, options types.MigrationOptions) error {
	resp, err := cli.post(ctx, "/containers/"+container+"/migrate", nil, options, nil)
	ensureReaderClosed(resp)
	return err
}
