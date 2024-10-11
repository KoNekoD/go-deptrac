package commands

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/commands_options"
	domainServices "github.com/KoNekoD/go-deptrac/pkg/domain/services"
	services "github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
)

// InitCommand - init - Creates a depfile template
type InitCommand struct {
	dumper      *domainServices.Dumper
	initOptions *commands_options.InitOptions
}

func NewInitCommand(dumper *domainServices.Dumper, initOptions *commands_options.InitOptions) *InitCommand {
	return &InitCommand{
		dumper:      dumper,
		initOptions: initOptions,
	}
}

func (c *InitCommand) Run(output services.OutputStyleInterface) error {
	err := c.dumper.Dump(c.initOptions.ConfigFile)
	if err != nil {
		output.Error(services.StringOrArrayOfStrings{String: "<fg=red>Depfile <info>dumping failed.</> error: " + err.Error()})
		return err
	}

	output.Success(services.StringOrArrayOfStrings{String: "Depfile <info>dumped.</>"})

	return nil
}
