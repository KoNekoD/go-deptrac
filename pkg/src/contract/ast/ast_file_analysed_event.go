package ast

// AstFileAnalysedEvent - Event triggered after parsing the AST of a file has been completed.
type AstFileAnalysedEvent struct {
	File string
}

func NewAstFileAnalysedEvent(file string) *AstFileAnalysedEvent {
	return &AstFileAnalysedEvent{File: file}
}
