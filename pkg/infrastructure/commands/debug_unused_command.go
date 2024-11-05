package commands

import (
	services2 "github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/commands_options"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/runners"
	"github.com/pkg/errors"
)

// DebugUnusedCommand - debug:unused - Lists unused (or barely used) layer dependencies
type DebugUnusedCommand struct {
	runner  *runners.DebugUnusedRunner
	options *commands_options.DebugUnusedOptions
}

func NewDebugUnusedCommand(runner *runners.DebugUnusedRunner, options *commands_options.DebugUnusedOptions) *DebugUnusedCommand {
	return &DebugUnusedCommand{runner: runner, options: options}
}

func (c *DebugUnusedCommand) Run(output services2.OutputInterface) error {
	err := c.runner.Run(output, c.options.Limit)
	if err != nil {
		output.GetStyle().Error(services2.StringOrArrayOfStrings{String: "<fg=red>Dependency debugging failed.</> error: " + err.Error()})
		return errors.WithStack(err)
	}

	return nil
}
