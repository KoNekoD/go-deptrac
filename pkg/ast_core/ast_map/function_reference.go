package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
)

type FunctionReference struct {
	*TaggedTokenReference
	functionName  *FunctionToken
	Dependencies  []*DependencyToken
	fileReference *FileReference
}

func NewFunctionReference(functionName *FunctionToken, dependencies []*DependencyToken, tags map[string][]string, fileReference *FileReference) *FunctionReference {
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

func (r *FunctionReference) GetToken() ast_contract.TokenInterface {
	return r.functionName
}

func (r *FunctionReference) GetDependencies() []*DependencyToken {
	return r.Dependencies
}
