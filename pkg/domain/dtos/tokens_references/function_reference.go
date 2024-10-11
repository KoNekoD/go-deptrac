package tokens_references

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
)

type FunctionReference struct {
	*TaggedTokenReference
	functionName  *tokens.FunctionToken
	Dependencies  []*dependencies.DependencyToken
	fileReference *FileReference
}

func NewFunctionReference(functionName *tokens.FunctionToken, dependencies []*dependencies.DependencyToken, tags map[string][]string, fileReference *FileReference) *FunctionReference {
	for _, dependency := range dependencies {
		if dependency.Token.ToString() == "" {
			panic("1")
		}
	}

	return &FunctionReference{
		functionName:         functionName,
		Dependencies:         dependencies,
		TaggedTokenReference: NewTaggedTokenReference(tags),
		fileReference:        fileReference,
	}
}

func (r *FunctionReference) WithFileReference(astFileReference *FileReference) *FunctionReference {
	return NewFunctionReference(r.functionName, r.Dependencies, r.Tags, astFileReference)
}

func (r *FunctionReference) GetFilepath() *string {
	return r.fileReference.Filepath
}

func (r *FunctionReference) GetToken() tokens.TokenInterface {
	return r.functionName
}

func (r *FunctionReference) GetDependencies() []*dependencies.DependencyToken {
	return r.Dependencies
}
