package runners

import (
	"fmt"
	services2 "github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/analysers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type DebugTokenRunner struct {
	analyser *analysers.LayerForTokenAnalyser
}

func NewDebugTokenRunner(analyser *analysers.LayerForTokenAnalyser) *DebugTokenRunner {
	return &DebugTokenRunner{analyser: analyser}
}

func (r *DebugTokenRunner) Run(output services2.OutputInterface, tokenName string, tokenType enums.TokenType) error {
	matches, err := r.analyser.FindLayerForToken(tokenName, tokenType)
	if err != nil {
		return apperrors.NewCommandRunExceptionAnalyserException(err)
	}

	if len(matches) == 0 {
		output.WriteLineFormatted(services2.StringOrArrayOfStrings{String: fmt.Sprintf("Could not find a token matching \"%s\"", tokenName)})
		return nil
	}

	headers := []string{"matching token", "layers"}
	rows := make([][]string, 0)
	for token, layers := range matches {
		layersJoined := "---"

		if len(layers) > 0 {
			layersJoined = fmt.Sprintf("%s", layers)
		}

		rows = append(rows, []string{token, layersJoined})
	}
	output.GetStyle().Table(headers, rows)
	return nil
}
