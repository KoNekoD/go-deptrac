package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
)

type ClassLikeReference struct {
	Type          *ClassLikeType
	classLikeName *ClassLikeToken

	Inherits      []*AstInherit
	Dependencies  []*DependencyToken
	fileReference *FileReference
	*TaggedTokenReference
}

func NewClassLikeReference(classLikeName *ClassLikeToken, classLikeType *ClassLikeType, inherits []*AstInherit, dependencies []*DependencyToken, tags map[string][]string, fileReference *FileReference) *ClassLikeReference {
	if classLikeType == nil {
		classLikeTypeTmp := TypeClasslike
		classLikeType = &classLikeTypeTmp
	}

	return &ClassLikeReference{
		Type:                 classLikeType,
		classLikeName:        classLikeName,
		Inherits:             inherits,
		Dependencies:         dependencies,
		fileReference:        fileReference,
		TaggedTokenReference: NewTaggedTokenReference(tags),
	}
}

func (c *ClassLikeReference) WithFileReference(astFileReference *FileReference) *ClassLikeReference {
	return NewClassLikeReference(c.classLikeName, c.Type, c.Inherits, c.Dependencies, c.Tags, astFileReference)
}

func (c *ClassLikeReference) GetFilepath() *string {
	return c.fileReference.Filepath
}

func (c *ClassLikeReference) GetToken() ast.TokenInterface {
	return c.classLikeName
}

func (c *ClassLikeReference) GetDependencies() []*DependencyToken {
	return c.Dependencies
}
