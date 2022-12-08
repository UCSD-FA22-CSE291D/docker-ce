package migrate

import (
	"context"
	"fmt"

	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/docker/docker/api/types"
	"github.com/spf13/cobra"
)

type migrateOptions struct {
	srcDockerAddr string
	dstDockerAddr string
	containerId   string
	predump       bool
}

// NewMigrateCommand creates a new `docker migrate` command
func NewMigrateCommand(dockerCli command.Cli) *cobra.Command {
	var opts migrateOptions

	cmd := &cobra.Command{
		Use:   "migrate [OPTIONS] CONTAINER MIGRATE",
		Short: "Live migration of a running container",
		Args:  cli.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.srcDockerAddr = args[0]
			opts.dstDockerAddr = args[1]
			opts.containerId = args[2]
			return runMigrate(dockerCli, opts)
		},
	}

	flags := cmd.Flags()

	flags.BoolVar(&opts.predump, "predump", false, "Use predump migration")

	return cmd
}

func runMigrate(dockerCli command.Cli, opts migrateOptions) error {
	client := dockerCli.Client()

	migrationOpts := types.MigrationOptions{
		SrcDockerdAddr: opts.srcDockerAddr,
		DstDockerdAddr: opts.dstDockerAddr,
		Predump:        opts.predump,
	}

	err := client.ContainerMigrate(context.Background(), opts.containerId, migrationOpts)
	if err != nil {
		return err
	}

	fmt.Fprintf(dockerCli.Out(), "%s\n", opts.dstDockerAddr)
	return nil
}
