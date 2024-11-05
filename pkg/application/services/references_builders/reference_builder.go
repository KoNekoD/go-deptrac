package references_builders

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type ReferenceBuilder struct {
	Dependencies   []*dependencies.DependencyToken
	tokenTemplates []string
	Filepath       string
	ref            *tokens_references.FileReference
}

type ReferenceBuilderInterface interface {
	GetTokenTemplates() []string
	CreateContext(occursAtLine int, dependencyType enums.DependencyType) *dependencies.DependencyContext
	UnresolvedFunctionCall(functionName string, occursAtLine int) *ReferenceBuilder
	Variable(classLikeName string, occursAtLine int) *ReferenceBuilder
	ReturnType(classLikeName string, occursAtLine int) *ReferenceBuilder
	ThrowStatement(classLikeName string, occursAtLine int) *ReferenceBuilder
	AnonymousClassExtends(classLikeName string, occursAtLine int) *ReferenceBuilder
	AnonymousClassTrait(classLikeName string, occursAtLine int) *ReferenceBuilder
	ConstFetch(classLikeName string, occursAtLine int) *ReferenceBuilder
	AnonymousClassImplements(classLikeName string, occursAtLine int) *ReferenceBuilder
	Parameter(classLikeName string, occursAtLine int) *ReferenceBuilder
	Attribute(classLikeName string, occursAtLine int) *ReferenceBuilder
	Instanceof(classLikeName string, occursAtLine int) *ReferenceBuilder
	NewStatement(classLikeName string, occursAtLine int) *ReferenceBuilder
	StaticProperty(classLikeName string, occursAtLine int) *ReferenceBuilder
	StaticMethod(classLikeName string, occursAtLine int) *ReferenceBuilder
	CatchStmt(classLikeName string, occursAtLine int) *ReferenceBuilder
	AddTokenTemplate(tokenTemplate string)
	RemoveTokenTemplate(tokenTemplate string)
}

func NewReferenceBuilder(tokenTemplates []string, filepath string) *ReferenceBuilder {
	return &ReferenceBuilder{
		Dependencies:   make([]*dependencies.DependencyToken, 0),
		tokenTemplates: tokenTemplates,
		Filepath:       filepath,
	}
}

func (r *ReferenceBuilder) GetTokenTemplates() []string {
	return r.tokenTemplates
}

func (r *ReferenceBuilder) CreateContext(occursAtLine int, dependencyType enums.DependencyType) *dependencies.DependencyContext {
	return dependencies.NewDependencyContext(dtos.NewFileOccurrence(r.Filepath, occursAtLine), dependencyType)
}

// UnresolvedFunctionCall - Unqualified function and constant names inside a namespace cannot be statically resolved. Inside a namespace Foo, a call to strlen() may either refer to the namespaced \Foo\strlen(), or the global \strlen(). Because PHP-ParserInterface does not have the necessary context to decide this, such names are left unresolved.
func (r *ReferenceBuilder) UnresolvedFunctionCall(functionName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewFunctionTokenFromFQCN(functionName), r.CreateContext(occursAtLine, enums.DependencyTypeUnresolvedFunctionCall)))
	return r
}

func (r *ReferenceBuilder) Variable(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, enums.DependencyTypeVariable)))
	return r
}

func (r *ReferenceBuilder) ReturnType(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, enums.DependencyTypeReturnType)))
	return r
}

func (r *ReferenceBuilder) ThrowStatement(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, enums.DependencyTypeThrow)))
	return r
}

func (r *ReferenceBuilder) AnonymousClassExtends(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, enums.DependencyTypeAnonymousClassExtends)))
	return r
}

func (r *ReferenceBuilder) AnonymousClassTrait(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, enums.DependencyTypeAnonymousClassTrait)))
	return r
}

func (r *ReferenceBuilder) ConstFetch(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, enums.DependencyTypeConst)))
	return r
}

func (r *ReferenceBuilder) AnonymousClassImplements(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, enums.DependencyTypeAnonymousClassImplements)))
	return r
}

func (r *ReferenceBuilder) Parameter(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, enums.DependencyTypeParameter)))
	return r
}

func (r *ReferenceBuilder) Attribute(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, enums.DependencyTypeAttribute)))
	return r
}

func (r *ReferenceBuilder) Instanceof(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, enums.DependencyTypeInstanceof)))
	return r
}

func (r *ReferenceBuilder) NewStatement(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, enums.DependencyTypeNew)))
	return r
}

func (r *ReferenceBuilder) StaticProperty(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, enums.DependencyTypeStaticProperty)))
	return r
}

func (r *ReferenceBuilder) StaticMethod(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, enums.DependencyTypeStaticMethod)))
	return r
}

func (r *ReferenceBuilder) CatchStmt(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, dependencies.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, enums.DependencyTypeCatch)))
	return r
}

func (r *ReferenceBuilder) AddTokenTemplate(tokenTemplate string) {
	r.tokenTemplates = append(r.tokenTemplates, tokenTemplate)
}

func (r *ReferenceBuilder) RemoveTokenTemplate(tokenTemplate string) {
	withoutTokenTemplate := make([]string, 0)
	for _, token := range r.tokenTemplates {
		if token != tokenTemplate {
			withoutTokenTemplate = append(withoutTokenTemplate, token)
		}
	}
	r.tokenTemplates = withoutTokenTemplate
}
