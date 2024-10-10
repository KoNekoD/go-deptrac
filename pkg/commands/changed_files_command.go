package commands

import (
	"github.com/KoNekoD/go-deptrac/pkg/formatters"
	"github.com/KoNekoD/go-deptrac/pkg/results"
	"github.com/KoNekoD/go-deptrac/pkg/runners"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type ChangedFilesCommand struct {
	runner *runners.ChangedFilesRunner
}

const (
	argWithDependencies = "with-dependencies"
	argFiles            = "files"
)

func NewChangedFilesCommand(runner *runners.ChangedFilesRunner) *cobra.Command {
	cmd := &ChangedFilesCommand{runner: runner}

	cobraCmd := &cobra.Command{
		Use:   "debug:changed-files",
		Short: "Lists layers corresponding to the changed files",
		RunE:  cmd.run,
	}

	cobraCmd.Flags().BoolP("verbose", "v", false, "verbose output")
	cobraCmd.Flags().BoolP("debug", "d", false, "debug output")

	cobraCmd.Flags().Bool(argWithDependencies, false, "show dependent layers")
	cobraCmd.Flags().StringArray(argFiles, []string{}, "List of changed files")
	_ = cobraCmd.MarkFlagRequired(argFiles)

	return cobraCmd
}

func (cmd *ChangedFilesCommand) run(cobraCmd *cobra.Command, args []string) error {
	symfonyOutput := results.NewSymfonyOutput(formatters.NewStyle(cobraCmd.Flags().Changed("verbose"), cobraCmd.Flags().Changed("debug")))

	files, err := cobraCmd.Flags().GetStringArray(argFiles)
	if err != nil {
		return errors.WithStack(err)
	}

	err = cmd.runner.Run(files, cobraCmd.Flags().Changed(argWithDependencies), symfonyOutput)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
