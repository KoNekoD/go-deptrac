package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
)

type FileReferenceBuilder struct {
	*ReferenceBuilder
	classReferences    []*ClassLikeReferenceBuilder
	functionReferences []*FunctionReferenceBuilder
}

func CreateFileReferenceBuilder(filepath string) *FileReferenceBuilder {
	return &FileReferenceBuilder{ReferenceBuilder: NewReferenceBuilder(make([]string, 0), filepath)}
}

func (b *FileReferenceBuilder) UseStatement(classLikeName string, occursAtLine int) *FileReferenceBuilder {
	b.Dependencies = append(b.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), b.CreateContext(occursAtLine, ast.DependencyTypeUse)))
	return b
}

func (b *FileReferenceBuilder) NewClass(classLikeName string, templateTypes []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	classReference := CreateClassLikeReferenceBuilderClass(b.Filepath, classLikeName, templateTypes, tags)
	b.classReferences = append(b.classReferences, classReference)
	return classReference
}

func (b *FileReferenceBuilder) NewTrait(classLikeName string, templateTypes []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	classReference := CreateClassLikeReferenceBuilderTrait(b.Filepath, classLikeName, templateTypes, tags)
	b.classReferences = append(b.classReferences, classReference)
	return classReference
}

func (b *FileReferenceBuilder) NewClassLike(classLikeName string, templateTypes []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	classReference := CreateClassLikeReferenceBuilderClassLike(b.Filepath, classLikeName, templateTypes, tags)
	b.classReferences = append(b.classReferences, classReference)
	return classReference
}

func (b *FileReferenceBuilder) NewInterface(classLikeName string, templateTypes []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	classReference := CreateClassLikeReferenceBuilderInterface(b.Filepath, classLikeName, templateTypes, tags)
	b.classReferences = append(b.classReferences, classReference)
	return classReference
}

func (b *FileReferenceBuilder) NewFunction(functionName string, templateTypes []string, tags map[string][]string) *FunctionReferenceBuilder {
	functionReference := CreateFunctionReferenceBuilder(b.Filepath, functionName, templateTypes, tags)
	b.functionReferences = append(b.functionReferences, functionReference)
	return functionReference
}

func (b *FileReferenceBuilder) Build() *FileReference {
	classReferences := make([]*ClassLikeReference, 0)
	for _, classReference := range b.classReferences {
		classReferences = append(classReferences, classReference.Build())
	}

	functionReferences := make([]*FunctionReference, 0)
	for _, functionReference := range b.functionReferences {
		functionReferences = append(functionReferences, functionReference.Build())
	}

	return NewFileReference(&b.Filepath, classReferences, functionReferences, b.Dependencies)
}
