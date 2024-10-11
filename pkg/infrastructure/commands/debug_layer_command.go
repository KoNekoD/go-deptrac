package commands

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/commands_options"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/runners"
	"github.com/pkg/errors"
)

// DebugLayerCommand - debug:layer - Checks which tokens belong to the provided layer
type DebugLayerCommand struct {
	runner  *runners.DebugLayerRunner
	options *commands_options.DebugLayerOptions
}

func NewDebugLayerCommand(debugLayerRunner *runners.DebugLayerRunner, options *commands_options.DebugLayerOptions) *DebugLayerCommand {
	return &DebugLayerCommand{runner: debugLayerRunner, options: options}
}

func (c *DebugLayerCommand) Run(output services.OutputStyleInterface) error {
	err := c.runner.Run(output, c.options.Layer)
	if err != nil {
		output.Error(services.StringOrArrayOfStrings{String: "<fg=red>Layer debugging failed.</> error: " + err.Error()})
		return errors.WithStack(err)
	}

	return nil
}
