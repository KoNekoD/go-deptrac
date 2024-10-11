package runners

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/services/analysers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
)

type DebugLayerRunner struct {
	analyser *analysers.TokenInLayerAnalyser
	layers   []*rules.Layer
}

func NewDebugLayerRunner(analyser *analysers.TokenInLayerAnalyser, layers []*rules.Layer) *DebugLayerRunner {
	return &DebugLayerRunner{analyser: analyser, layers: layers}
}

func (d *DebugLayerRunner) Run(output services.OutputStyleInterface, layer *string) error {
	debugLayers := make([]string, 0)

	if layer != nil {
		debugLayers = append(debugLayers, *layer)
	} else {
		debugLayers = utils.MapSlice(d.layers, func(layer *rules.Layer) string { return layer.Name })
	}

	for _, debugLayer := range debugLayers {
		found, err := d.analyser.FindTokensInLayer(debugLayer)
		if err != nil {
			return apperrors.NewCommandRunExceptionAnalyserException(err)
		}

		values := utils.Map(found, func(t1 string, t2 enums.TokenType) []string { return []string{t1, string(t2)} })

		output.Table([]string{debugLayer, "Token Type"}, values)
	}

	return nil
}
