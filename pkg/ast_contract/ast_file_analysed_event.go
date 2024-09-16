package ast_contract

// AstFileAnalysedEvent - Event triggered after parsing the AST of a file_supportive has been completed.
type AstFileAnalysedEvent struct {
	File string
}

func NewAstFileAnalysedEvent(file string) *AstFileAnalysedEvent {
	return &AstFileAnalysedEvent{File: file}
}
