package references

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	enums2 "github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/violations"
)

type ClassLikeReferenceBuilder struct {
	*ReferenceBuilder

	inherits []*ast_map.AstInherit

	classLikeToken *tokens.ClassLikeToken
	classLikeType  *enums2.ClassLikeType
	tags           map[string][]string
}

func NewClassLikeReferenceBuilder(tokenTemplates []string, filepath string, classLikeToken *tokens.ClassLikeToken, classLikeType *enums2.ClassLikeType, tags map[string][]string) *ClassLikeReferenceBuilder {
	return &ClassLikeReferenceBuilder{
		ReferenceBuilder: NewReferenceBuilder(tokenTemplates, filepath),
		inherits:         make([]*ast_map.AstInherit, 0),
		classLikeToken:   classLikeToken,
		classLikeType:    classLikeType,
		tags:             tags,
	}
}

func CreateClassLikeReferenceBuilderClassLike(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeClassLike := enums2.TypeClasslike
	return NewClassLikeReferenceBuilder(classTemplates, filepath, tokens.NewClassLikeTokenFromFQCN(classLikeName), &typeClassLike, tags)
}

func CreateClassLikeReferenceBuilderClass(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeClass := enums2.TypeClass
	return NewClassLikeReferenceBuilder(classTemplates, filepath, tokens.NewClassLikeTokenFromFQCN(classLikeName), &typeClass, tags)
}

func CreateClassLikeReferenceBuilderTrait(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeTrait := enums2.TypeTrait
	return NewClassLikeReferenceBuilder(classTemplates, filepath, tokens.NewClassLikeTokenFromFQCN(classLikeName), &typeTrait, tags)
}

func CreateClassLikeReferenceBuilderInterface(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeInterface := enums2.TypeInterface
	return NewClassLikeReferenceBuilder(classTemplates, filepath, tokens.NewClassLikeTokenFromFQCN(classLikeName), &typeInterface, tags)
}

// Build - Internal
func (b *ClassLikeReferenceBuilder) Build() *ClassLikeReference {
	return NewClassLikeReference(b.classLikeToken, b.classLikeType, b.inherits, b.Dependencies, b.tags, nil)
}

func (b *ClassLikeReferenceBuilder) Extends(classLikeName string, occursAtLine int) *ClassLikeReferenceBuilder {
	b.inherits = append(b.inherits, ast_map.NewAstInherit(tokens.NewClassLikeTokenFromFQCN(classLikeName), violations.NewFileOccurrence(b.Filepath, occursAtLine), enums2.AstInheritTypeExtends, make([]*ast_map.AstInherit, 0)))
	return b
}

func (b *ClassLikeReferenceBuilder) Implements(classLikeName string, occursAtLine int) *ClassLikeReferenceBuilder {
	b.inherits = append(b.inherits, ast_map.NewAstInherit(tokens.NewClassLikeTokenFromFQCN(classLikeName), violations.NewFileOccurrence(b.Filepath, occursAtLine), enums2.AstInheritTypeImplements, make([]*ast_map.AstInherit, 0)))
	return b
}

func (b *ClassLikeReferenceBuilder) Trait(classLikeName string, occursAtLine int) *ClassLikeReferenceBuilder {
	b.inherits = append(b.inherits, ast_map.NewAstInherit(tokens.NewClassLikeTokenFromFQCN(classLikeName), violations.NewFileOccurrence(b.Filepath, occursAtLine), enums2.AstInheritTypeUses, make([]*ast_map.AstInherit, 0)))
	return b
}
