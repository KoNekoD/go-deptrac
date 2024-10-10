package references

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/types"
	"github.com/KoNekoD/go-deptrac/pkg/violations"
)

type ClassLikeReferenceBuilder struct {
	*ReferenceBuilder

	inherits []*ast_map.AstInherit

	classLikeToken *tokens.ClassLikeToken
	classLikeType  *types.ClassLikeType
	tags           map[string][]string
}

func NewClassLikeReferenceBuilder(tokenTemplates []string, filepath string, classLikeToken *tokens.ClassLikeToken, classLikeType *types.ClassLikeType, tags map[string][]string) *ClassLikeReferenceBuilder {
	return &ClassLikeReferenceBuilder{
		ReferenceBuilder: NewReferenceBuilder(tokenTemplates, filepath),
		inherits:         make([]*ast_map.AstInherit, 0),
		classLikeToken:   classLikeToken,
		classLikeType:    classLikeType,
		tags:             tags,
	}
}

func CreateClassLikeReferenceBuilderClassLike(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeClassLike := types.TypeClasslike
	return NewClassLikeReferenceBuilder(classTemplates, filepath, tokens.NewClassLikeTokenFromFQCN(classLikeName), &typeClassLike, tags)
}

func CreateClassLikeReferenceBuilderClass(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeClass := types.TypeClass
	return NewClassLikeReferenceBuilder(classTemplates, filepath, tokens.NewClassLikeTokenFromFQCN(classLikeName), &typeClass, tags)
}

func CreateClassLikeReferenceBuilderTrait(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeTrait := types.TypeTrait
	return NewClassLikeReferenceBuilder(classTemplates, filepath, tokens.NewClassLikeTokenFromFQCN(classLikeName), &typeTrait, tags)
}

func CreateClassLikeReferenceBuilderInterface(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeInterface := types.TypeInterface
	return NewClassLikeReferenceBuilder(classTemplates, filepath, tokens.NewClassLikeTokenFromFQCN(classLikeName), &typeInterface, tags)
}

// Build - Internal
func (b *ClassLikeReferenceBuilder) Build() *ClassLikeReference {
	return NewClassLikeReference(b.classLikeToken, b.classLikeType, b.inherits, b.Dependencies, b.tags, nil)
}

func (b *ClassLikeReferenceBuilder) Extends(classLikeName string, occursAtLine int) *ClassLikeReferenceBuilder {
	b.inherits = append(b.inherits, ast_map.NewAstInherit(tokens.NewClassLikeTokenFromFQCN(classLikeName), violations.NewFileOccurrence(b.Filepath, occursAtLine), ast_map.AstInheritTypeExtends, make([]*ast_map.AstInherit, 0)))
	return b
}

func (b *ClassLikeReferenceBuilder) Implements(classLikeName string, occursAtLine int) *ClassLikeReferenceBuilder {
	b.inherits = append(b.inherits, ast_map.NewAstInherit(tokens.NewClassLikeTokenFromFQCN(classLikeName), violations.NewFileOccurrence(b.Filepath, occursAtLine), ast_map.AstInheritTypeImplements, make([]*ast_map.AstInherit, 0)))
	return b
}

func (b *ClassLikeReferenceBuilder) Trait(classLikeName string, occursAtLine int) *ClassLikeReferenceBuilder {
	b.inherits = append(b.inherits, ast_map.NewAstInherit(tokens.NewClassLikeTokenFromFQCN(classLikeName), violations.NewFileOccurrence(b.Filepath, occursAtLine), ast_map.AstInheritTypeUses, make([]*ast_map.AstInherit, 0)))
	return b
}
