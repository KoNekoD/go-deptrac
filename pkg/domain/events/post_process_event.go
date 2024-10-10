package events

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/analysis_results"
)

// PostProcessEvent - Event fired after the analysis is complete. Useful if you want to change the result_contract of the analysis after it has completed and before it is returned for output processing.
type PostProcessEvent struct {
	result *analysis_results.AnalysisResult
}

func NewPostProcessEvent(result *analysis_results.AnalysisResult) *PostProcessEvent {
	return &PostProcessEvent{
		result: result,
	}
}

func (e *PostProcessEvent) GetResult() *analysis_results.AnalysisResult {
	return e.result
}

func (e *PostProcessEvent) ReplaceResult(result *analysis_results.AnalysisResult) {
	e.result = result
}
