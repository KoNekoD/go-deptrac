package commands

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/commands_options"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/runners"
	"github.com/pkg/errors"
)

// DebugDependenciesCommand - debug:dependencies - List layer dependencies
type DebugDependenciesCommand struct {
	runner  *runners.DebugDependenciesRunner
	options *commands_options.DebugDependenciesOptions
}

func NewDebugDependenciesCommand(debugDependenciesRunner *runners.DebugDependenciesRunner, options *commands_options.DebugDependenciesOptions) *DebugDependenciesCommand {
	return &DebugDependenciesCommand{runner: debugDependenciesRunner, options: options}
}

func (c *DebugDependenciesCommand) Run(output services.OutputStyleInterface) error {
	err := c.runner.Run(output, c.options.Layer, c.options.TargetLayer)
	if err != nil {
		output.Error(services.StringOrArrayOfStrings{String: "<fg=red>Dependency debugging failed.</> error: " + err.Error()})
		return errors.WithStack(err)
	}

	return nil
}
