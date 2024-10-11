package commands

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/commands_options"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/runners"
	"github.com/pkg/errors"
)

// DebugTokenCommand - debug:token|debug:class-like - Checks which layers the provided token belongs to
type DebugTokenCommand struct {
	runner  *runners.DebugTokenRunner
	options *commands_options.DebugTokenOptions
}

func NewDebugTokenCommand(runner *runners.DebugTokenRunner, options *commands_options.DebugTokenOptions) *DebugTokenCommand {
	return &DebugTokenCommand{runner: runner, options: options}
}

func (c *DebugTokenCommand) Run(output services.OutputInterface) error {
	tokenType, err := enums.NewTokenType(c.options.Type)
	if err != nil {
		return err
	}

	err = c.runner.Run(output, c.options.Token, tokenType)
	if err != nil {
		output.GetStyle().Error(services.StringOrArrayOfStrings{String: "<fg=red>Token debugging failed.</> error: " + err.Error()})
		return errors.WithStack(err)
	}

	return nil
}
