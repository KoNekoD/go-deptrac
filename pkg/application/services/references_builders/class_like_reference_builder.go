package references_builders

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_inherits"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ClassLikeReferenceBuilder struct {
	*ReferenceBuilder

	inherits []*ast_inherits.AstInherit

	classLikeToken *tokens.ClassLikeToken
	classLikeType  *enums.ClassLikeType
	tags           map[string][]string
}

func NewClassLikeReferenceBuilder(tokenTemplates []string, filepath string, classLikeToken *tokens.ClassLikeToken, classLikeType *enums.ClassLikeType, tags map[string][]string) *ClassLikeReferenceBuilder {
	return &ClassLikeReferenceBuilder{
		ReferenceBuilder: NewReferenceBuilder(tokenTemplates, filepath),
		inherits:         make([]*ast_inherits.AstInherit, 0),
		classLikeToken:   classLikeToken,
		classLikeType:    classLikeType,
		tags:             tags,
	}
}

func CreateClassLikeReferenceBuilderClassLike(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeClassLike := enums.TypeClasslike
	return NewClassLikeReferenceBuilder(classTemplates, filepath, tokens.NewClassLikeTokenFromFQCN(classLikeName), &typeClassLike, tags)
}

func CreateClassLikeReferenceBuilderClass(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeClass := enums.TypeClass
	return NewClassLikeReferenceBuilder(classTemplates, filepath, tokens.NewClassLikeTokenFromFQCN(classLikeName), &typeClass, tags)
}

func CreateClassLikeReferenceBuilderTrait(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeTrait := enums.TypeTrait
	return NewClassLikeReferenceBuilder(classTemplates, filepath, tokens.NewClassLikeTokenFromFQCN(classLikeName), &typeTrait, tags)
}

func CreateClassLikeReferenceBuilderInterface(filepath string, classLikeName string, classTemplates []string, tags map[string][]string) *ClassLikeReferenceBuilder {
	typeInterface := enums.TypeInterface
	return NewClassLikeReferenceBuilder(classTemplates, filepath, tokens.NewClassLikeTokenFromFQCN(classLikeName), &typeInterface, tags)
}

// Build - Internal
func (b *ClassLikeReferenceBuilder) Build() *tokens_references.ClassLikeReference {
	return tokens_references.NewClassLikeReference(b.classLikeToken, b.classLikeType, b.inherits, b.Dependencies, b.tags, nil)
}

func (b *ClassLikeReferenceBuilder) Extends(classLikeName string, occursAtLine int) *ClassLikeReferenceBuilder {
	b.inherits = append(b.inherits, ast_inherits.NewAstInherit(tokens.NewClassLikeTokenFromFQCN(classLikeName), dtos.NewFileOccurrence(b.Filepath, occursAtLine), enums.AstInheritTypeExtends, make([]*ast_inherits.AstInherit, 0)))
	return b
}

func (b *ClassLikeReferenceBuilder) Implements(classLikeName string, occursAtLine int) *ClassLikeReferenceBuilder {
	b.inherits = append(b.inherits, ast_inherits.NewAstInherit(tokens.NewClassLikeTokenFromFQCN(classLikeName), dtos.NewFileOccurrence(b.Filepath, occursAtLine), enums.AstInheritTypeImplements, make([]*ast_inherits.AstInherit, 0)))
	return b
}

func (b *ClassLikeReferenceBuilder) Trait(classLikeName string, occursAtLine int) *ClassLikeReferenceBuilder {
	b.inherits = append(b.inherits, ast_inherits.NewAstInherit(tokens.NewClassLikeTokenFromFQCN(classLikeName), dtos.NewFileOccurrence(b.Filepath, occursAtLine), enums.AstInheritTypeUses, make([]*ast_inherits.AstInherit, 0)))
	return b
}
