package runners

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/services/analysers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
)

type DebugUnassignedRunner struct {
	analyser *analysers.UnassignedTokenAnalyser
}

func NewDebugUnassignedRunner(analyser *analysers.UnassignedTokenAnalyser) *DebugUnassignedRunner {
	return &DebugUnassignedRunner{analyser: analyser}
}

// Run - returns are there any unassigned tokens?
func (d *DebugUnassignedRunner) Run(output services.OutputInterface) (bool, error) {
	unassignedTokens, err := d.analyser.FindUnassignedTokens()
	if err != nil {
		return false, apperrors.NewCommandRunExceptionAnalyserException(err)
	}

	if len(unassignedTokens) == 0 {
		output.WriteLineFormatted(services.StringOrArrayOfStrings{String: "There are no unassigned tokens."})
		return false, nil
	}

	output.WriteLineFormatted(services.StringOrArrayOfStrings{Strings: unassignedTokens})

	return true, err
}
