package post_process_event

import (
	"github.com/KoNekoD/go-deptrac/pkg/analyser_contract/analysis_result"
)

// PostProcessEvent - Event fired after the analysis is complete. Useful if you want to change the result_contract of the analysis after it has completed and before it is returned for output processing.
type PostProcessEvent struct {
	result *analysis_result.AnalysisResult
}

func NewPostProcessEvent(result *analysis_result.AnalysisResult) *PostProcessEvent {
	return &PostProcessEvent{
		result: result,
	}
}

func (e *PostProcessEvent) GetResult() *analysis_result.AnalysisResult {
	return e.result
}

func (e *PostProcessEvent) ReplaceResult(result *analysis_result.AnalysisResult) {
	e.result = result
}
