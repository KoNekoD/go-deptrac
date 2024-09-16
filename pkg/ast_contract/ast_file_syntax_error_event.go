package ast_contract

// AstFileSyntaxErrorEvent - Event triggered when parsing the AST failed on syntax error in the PHP file_supportive.
type AstFileSyntaxErrorEvent struct {
	File        string
	SyntaxError string
}

func NewAstFileSyntaxErrorEvent(file string, syntaxError string) *AstFileSyntaxErrorEvent {
	return &AstFileSyntaxErrorEvent{File: file, SyntaxError: syntaxError}
}
