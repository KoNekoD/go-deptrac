package PostProcessEvent

import "github.com/KoNekoD/go-deptrac/pkg/src/contract/Analyser/AnalysisResult"

// PostProcessEvent - Event fired after the analysis is complete. Useful if you want to change the result of the analysis after it has completed and before it is returned for output processing.
type PostProcessEvent struct {
	result *AnalysisResult.AnalysisResult
}

func NewPostProcessEvent(result *AnalysisResult.AnalysisResult) *PostProcessEvent {
	return &PostProcessEvent{
		result: result,
	}
}

func (e *PostProcessEvent) GetResult() *AnalysisResult.AnalysisResult {
	return e.result
}

func (e *PostProcessEvent) ReplaceResult(result *AnalysisResult.AnalysisResult) {
	e.result = result
}
