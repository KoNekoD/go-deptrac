package commands

import (
	services "github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/commands_options"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/runners"
	"github.com/pkg/errors"
)

type ChangedFilesCommand struct {
	runner  *runners.ChangedFilesRunner
	options *commands_options.ChangedFilesOptions
}

func NewChangedFilesCommand(runner *runners.ChangedFilesRunner, options *commands_options.ChangedFilesOptions) *ChangedFilesCommand {
	return &ChangedFilesCommand{runner: runner, options: options}
}

func (c *ChangedFilesCommand) Run(output services.OutputInterface) error {
	err := c.runner.Run(c.options.Files, c.options.WithDependencies, output)
	if err != nil {
		output.GetStyle().Error(services.StringOrArrayOfStrings{String: "<fg=red>Changed files debugging failed.</> error: " + err.Error()})
		return errors.WithStack(err)
	}

	return nil
}
