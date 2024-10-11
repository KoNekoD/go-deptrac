package runners

import (
	services2 "github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/analysers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
)

type DebugUnassignedRunner struct {
	analyser *analysers.UnassignedTokenAnalyser
}

func NewDebugUnassignedRunner(analyser *analysers.UnassignedTokenAnalyser) *DebugUnassignedRunner {
	return &DebugUnassignedRunner{analyser: analyser}
}

// Run - returns are there any unassigned tokens?
func (d *DebugUnassignedRunner) Run(output services2.OutputInterface) (bool, error) {
	unassignedTokens, err := d.analyser.FindUnassignedTokens()
	if err != nil {
		return false, apperrors.NewCommandRunExceptionAnalyserException(err)
	}

	if len(unassignedTokens) == 0 {
		output.WriteLineFormatted(services2.StringOrArrayOfStrings{String: "There are no unassigned tokens."})
		return false, nil
	}

	output.WriteLineFormatted(services2.StringOrArrayOfStrings{Strings: unassignedTokens})

	return true, err
}
