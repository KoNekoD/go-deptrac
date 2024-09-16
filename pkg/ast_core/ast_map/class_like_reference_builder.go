package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_contract"
)

type ClassLikeReferenceBuilder struct {
	*ReferenceBuilder

	inherits []*AstInherit

	classLikeToken *ClassLikeToken
	classLikeType  *ClassLikeType
	tags           map[string][]string
}

func NewClassLikeReferenceBuilder(tokenTemplates []string, filepath string, classLikeToken *ClassLikeToken, classLikeType *ClassLikeType, tags map[string][]string) *ClassLikeReferenceBuilder {
	return &ClassLikeReferenceBuilder{
		ReferenceBuilder: NewReferenceBuilder(tokenTemplates, filepath),
		inherits:         make([]*AstInherit, 0),
		classLikeToken:   classLikeToken,
		classLikeType:    classLikeType,
		tags:             tags,
	}
}

func CreateClassLikeReferenceBuilderClassLike(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeClassLike := TypeClasslike
	return NewClassLikeReferenceBuilder(classTemplates, filepath, NewClassLikeTokenFromFQCN(classLikeName), &typeClassLike, tags)
}

func CreateClassLikeReferenceBuilderClass(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeClass := TypeClass
	return NewClassLikeReferenceBuilder(classTemplates, filepath, NewClassLikeTokenFromFQCN(classLikeName), &typeClass, tags)
}

func CreateClassLikeReferenceBuilderTrait(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeTrait := TypeTrait
	return NewClassLikeReferenceBuilder(classTemplates, filepath, NewClassLikeTokenFromFQCN(classLikeName), &typeTrait, tags)
}

func CreateClassLikeReferenceBuilderInterface(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeInterface := TypeInterface
	return NewClassLikeReferenceBuilder(classTemplates, filepath, NewClassLikeTokenFromFQCN(classLikeName), &typeInterface, tags)
}

// Build - Internal
func (b *ClassLikeReferenceBuilder) Build() *ClassLikeReference {
	return NewClassLikeReference(b.classLikeToken, b.classLikeType, b.inherits, b.Dependencies, b.tags, nil)
}

func (b *ClassLikeReferenceBuilder) Extends(classLikeName string, occursAtLine int) *ClassLikeReferenceBuilder {
	b.inherits = append(b.inherits, NewAstInherit(NewClassLikeTokenFromFQCN(classLikeName), ast_contract.NewFileOccurrence(b.Filepath, occursAtLine), AstInheritTypeExtends, make([]*AstInherit, 0)))
	return b
}

func (b *ClassLikeReferenceBuilder) Implements(classLikeName string, occursAtLine int) *ClassLikeReferenceBuilder {
	b.inherits = append(b.inherits, NewAstInherit(NewClassLikeTokenFromFQCN(classLikeName), ast_contract.NewFileOccurrence(b.Filepath, occursAtLine), AstInheritTypeImplements, make([]*AstInherit, 0)))
	return b
}

func (b *ClassLikeReferenceBuilder) Trait(classLikeName string, occursAtLine int) *ClassLikeReferenceBuilder {
	b.inherits = append(b.inherits, NewAstInherit(NewClassLikeTokenFromFQCN(classLikeName), ast_contract.NewFileOccurrence(b.Filepath, occursAtLine), AstInheritTypeUses, make([]*AstInherit, 0)))
	return b
}
