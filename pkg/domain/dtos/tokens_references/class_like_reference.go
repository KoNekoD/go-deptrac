package tokens_references

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_map"
	tokens2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ClassLikeReference struct {
	Type          *enums.ClassLikeType
	classLikeName *tokens2.ClassLikeToken

	Inherits      []*ast_map.AstInherit
	Dependencies  []*tokens2.DependencyToken
	fileReference *FileReference
	*TaggedTokenReference
}

func NewClassLikeReference(classLikeName *tokens2.ClassLikeToken, classLikeType *enums.ClassLikeType, inherits []*ast_map.AstInherit, dependencies []*tokens2.DependencyToken, tags map[string][]string, fileReference *FileReference) *ClassLikeReference {
	if classLikeType == nil {
		classLikeTypeTmp := enums.TypeClasslike
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

func (c *ClassLikeReference) GetToken() tokens2.TokenInterface {
	return c.classLikeName
}

func (c *ClassLikeReference) GetDependencies() []*tokens2.DependencyToken {
	return c.Dependencies
}
