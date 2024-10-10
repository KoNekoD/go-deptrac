package events

import (
	"github.com/KoNekoD/go-deptrac/pkg/rules"
)

// PostProcessEvent - Event fired after the analysis is complete. Useful if you want to change the result_contract of the analysis after it has completed and before it is returned for output processing.
type PostProcessEvent struct {
	result *rules.AnalysisResult
}

func NewPostProcessEvent(result *rules.AnalysisResult) *PostProcessEvent {
	return &PostProcessEvent{
		result: result,
	}
}

func (e *PostProcessEvent) GetResult() *rules.AnalysisResult {
	return e.result
}

func (e *PostProcessEvent) ReplaceResult(result *rules.AnalysisResult) {
	e.result = result
}
