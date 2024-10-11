package commands

import (
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services/runners"
	"github.com/pkg/errors"
)

// DebugUnassignedCommand - debug:unassigned - Lists tokens that are not assigned to any layer
type DebugUnassignedCommand struct {
	runner *runners.DebugUnassignedRunner
}

func NewDebugUnassignedCommand(runner *runners.DebugUnassignedRunner) *DebugUnassignedCommand {
	return &DebugUnassignedCommand{runner: runner}
}

func (c *DebugUnassignedCommand) Run(output services.OutputInterface) error {
	result, err := c.runner.Run(output)
	if err != nil {
		output.GetStyle().Error(services.StringOrArrayOfStrings{String: "<fg=red>Unassigned token debugging failed.</> error: " + err.Error()})
		return errors.WithStack(err)
	}

	if result {
		output.WriteLineFormatted(services.StringOrArrayOfStrings{String: "There are unassigned tokens."})

		return errors.New("There are unassigned tokens.")
	}

	return nil
}
