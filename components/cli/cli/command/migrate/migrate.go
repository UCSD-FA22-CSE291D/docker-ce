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
	container    string
	target       string
	diskless     bool
	leaveRunning bool
}

// NewMigrateCommand creates a new `docker migrate` command
func NewMigrateCommand(dockerCli command.Cli) *cobra.Command {
	var opts migrateOptions

	cmd := &cobra.Command{
		Use:   "migrate [OPTIONS] CONTAINER MIGRATE",
		Short: "Live migration of a running container",
		Args:  cli.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.container = args[0]
			opts.target = args[1]
			return runMigrate(dockerCli, opts)
		},
	}

	flags := cmd.Flags()

	flags.BoolVar(&opts.diskless, "diskless", false, "Use diskless migration")
	flags.BoolVar(&opts.leaveRunning, "leave-running", false, "Leave the container running after migration")

	return cmd
}

func runMigrate(dockerCli command.Cli, opts migrateOptions) error {
	client := dockerCli.Client()

	migrationOpts := types.MigrationOptions{
		TargetAddr: opts.target,
		Exit:       !opts.leaveRunning,
	}

	err := client.ContainerMigrate(context.Background(), opts.container, migrationOpts)
	if err != nil {
		return err
	}

	fmt.Fprintf(dockerCli.Out(), "%s\n", opts.target)
	return nil
}
