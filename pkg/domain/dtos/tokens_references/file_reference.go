package tokens_references

import (
	tokens2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
)

type FileReference struct {
	Filepath            *string
	ClassLikeReferences []*ClassLikeReference
	FunctionReferences  []*FunctionReference
	Dependencies        []*tokens2.DependencyToken
}

func NewFileReference(filepath *string, structLikeReferences []*ClassLikeReference, functionReferences []*FunctionReference, dependencies []*tokens2.DependencyToken) *FileReference {
	structLikeReferencesMapped := make([]*ClassLikeReference, 0)
	functionReferencesMapped := make([]*FunctionReference, 0)

	ref := &FileReference{}

	for _, structLikeReference := range structLikeReferences {
		structLikeReferencesMapped = append(structLikeReferencesMapped, structLikeReference.WithFileReference(ref))
	}

	for _, functionReference := range functionReferences {
		functionReferencesMapped = append(functionReferencesMapped, functionReference.WithFileReference(ref))
	}

	ref.ClassLikeReferences = structLikeReferencesMapped
	ref.FunctionReferences = functionReferencesMapped
	ref.Filepath = filepath
	ref.Dependencies = dependencies

	return ref
}

func (r *FileReference) GetFilepath() *string {
	return r.Filepath
}

func (r *FileReference) GetToken() tokens2.TokenInterface {
	return tokens2.NewFileToken(r.Filepath)
}

func (r *FileReference) GetDependencies() []*tokens2.DependencyToken {
	return r.Dependencies
}
