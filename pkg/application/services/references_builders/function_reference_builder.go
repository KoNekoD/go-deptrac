package references_builders

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
)

type FunctionReferenceBuilder struct {
	*ReferenceBuilder
	functionName string
	tags         map[string][]string
}

func newFunctionReferenceBuilder(tokenTemplates []string, filepath string, functionName string, tags map[string][]string) *FunctionReferenceBuilder {
	if functionName == "" {
		panic("1")
	}

	return &FunctionReferenceBuilder{ReferenceBuilder: NewReferenceBuilder(tokenTemplates, filepath), functionName: functionName, tags: tags}
}

func CreateFunctionReferenceBuilder(filepath string, functionName string, functionTemplates []string, tags map[string][]string) *FunctionReferenceBuilder {
	return newFunctionReferenceBuilder(functionTemplates, filepath, functionName, tags)
}

// Build - Internal
func (b *FunctionReferenceBuilder) Build() *tokens_references.FunctionReference {
	return tokens_references.NewFunctionReference(tokens.NewFunctionTokenFromFQCN(b.functionName), b.Dependencies, b.tags, nil)
}
