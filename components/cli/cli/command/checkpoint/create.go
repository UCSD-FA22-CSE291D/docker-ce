package checkpoint

import (
	"context"
	"fmt"

	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/docker/docker/api/types"
	"github.com/spf13/cobra"
)

type createOptions struct {
	container          string
	checkpoint         string
	checkpointDir      string
	leaveRunning       bool
	predump            bool
	parentCheckpointId string
}

func newCreateCommand(dockerCli command.Cli) *cobra.Command {
	var opts createOptions

	cmd := &cobra.Command{
		Use:   "create [OPTIONS] CONTAINER CHECKPOINT",
		Short: "Create a checkpoint from a running container",
		Args:  cli.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.container = args[0]
			opts.checkpoint = args[1]
			return runCreate(dockerCli, opts)
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&opts.leaveRunning, "leave-running", false, "Leave the container running after checkpoint")
	flags.StringVarP(&opts.checkpointDir, "checkpoint-dir", "", "", "Use a custom checkpoint storage directory")
	flags.BoolVar(&opts.predump, "pre-dump", false, "Create a pre-dump of the container")
	flags.StringVarP(&opts.parentCheckpointId, "parent-checkpoint-id", "", "", "ID of the parent checkpoint")

	return cmd
}

func runCreate(dockerCli command.Cli, opts createOptions) error {
	client := dockerCli.Client()

	checkpointOpts := types.CheckpointCreateOptions{
		CheckpointID:       opts.checkpoint,
		CheckpointDir:      opts.checkpointDir,
		Exit:               !opts.leaveRunning || !opts.predump,
		Predump:            opts.predump,
		ParentCheckpointID: opts.parentCheckpointId,
	}

	err := client.CheckpointCreate(context.Background(), opts.container, checkpointOpts)
	if err != nil {
		return err
	}

	fmt.Fprintf(dockerCli.Out(), "%s\n", opts.checkpoint)
	return nil
}
