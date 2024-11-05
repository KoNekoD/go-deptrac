package tokens

import (
	"regexp"
	"strings"
)

type FunctionToken struct {
	FunctionName string
}

func NewFunctionToken(functionName string) *FunctionToken {
	if functionName == "" {
		panic("1")
	}
	return &FunctionToken{
		FunctionName: functionName,
	}
}

func NewFunctionTokenFromFQCN(functionName string) *FunctionToken {
	if functionName == "" {
		panic("1")
	}
	return &FunctionToken{
		FunctionName: strings.TrimPrefix(functionName, "\\"),
	}
}

func (t *FunctionToken) Match(pattern string) bool {
	r, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	matches := r.FindStringSubmatch(t.FunctionName)
	return len(matches) > 0
}

func (t *FunctionToken) ToString() string {
	return t.FunctionName + "()"
}

func (t *FunctionToken) Equals(token *FunctionToken) bool {
	return t.FunctionName == token.FunctionName
}

func (t *FunctionToken) tokenInterface() {}
