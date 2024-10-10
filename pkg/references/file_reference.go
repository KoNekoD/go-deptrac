package references

import "github.com/KoNekoD/go-deptrac/pkg/tokens"

type FileReference struct {
	Filepath            *string
	ClassLikeReferences []*ClassLikeReference
	FunctionReferences  []*FunctionReference
	Dependencies        []*tokens.DependencyToken
}

func NewFileReference(filepath *string, structLikeReferences []*ClassLikeReference, functionReferences []*FunctionReference, dependencies []*tokens.DependencyToken) *FileReference {
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

func (r *FileReference) GetToken() tokens.TokenInterface {
	return tokens.NewFileToken(r.Filepath)
}

func (r *FileReference) GetDependencies() []*tokens.DependencyToken {
	return r.Dependencies
}
