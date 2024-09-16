package ast_contract

// PreCreateAstMapEvent - Event triggered before the AST map and parsing of all files has started.
type PreCreateAstMapEvent struct {
	ExpectedFileCount int
}

func NewPreCreateAstMapEvent(expectedFileCount int) *PreCreateAstMapEvent {
	return &PreCreateAstMapEvent{ExpectedFileCount: expectedFileCount}
}
