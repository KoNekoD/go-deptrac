package references

import (
	"github.com/KoNekoD/go-deptrac/pkg/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/violations"
)

type ReferenceBuilder struct {
	Dependencies   []*tokens.DependencyToken
	tokenTemplates []string
	Filepath       string
	ref            *FileReference
}

type ReferenceBuilderInterface interface {
	GetTokenTemplates() []string
	CreateContext(occursAtLine int, dependencyType dependencies.DependencyType) *dependencies.DependencyContext
	UnresolvedFunctionCall(functionName string, occursAtLine int) *ReferenceBuilder
	Variable(classLikeName string, occursAtLine int) *ReferenceBuilder
	Superglobal(superglobalName string, occursAtLine int) *ReferenceBuilder
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
		Dependencies:   make([]*tokens.DependencyToken, 0),
		tokenTemplates: tokenTemplates,
		Filepath:       filepath,
	}
}

func (r *ReferenceBuilder) GetTokenTemplates() []string {
	return r.tokenTemplates
}

func (r *ReferenceBuilder) CreateContext(occursAtLine int, dependencyType dependencies.DependencyType) *dependencies.DependencyContext {
	return dependencies.NewDependencyContext(violations.NewFileOccurrence(r.Filepath, occursAtLine), dependencyType)
}

// UnresolvedFunctionCall - Unqualified function and constant names inside a namespace cannot be statically resolved. Inside a namespace Foo, a call to strlen() may either refer to the namespaced \Foo\strlen(), or the global \strlen(). Because PHP-ParserInterface does not have the necessary context to decide this, such names are left unresolved.
func (r *ReferenceBuilder) UnresolvedFunctionCall(functionName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewFunctionTokenFromFQCN(functionName), r.CreateContext(occursAtLine, dependencies.DependencyTypeUnresolvedFunctionCall)))
	return r
}

func (r *ReferenceBuilder) Variable(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, dependencies.DependencyTypeVariable)))
	return r
}

func (r *ReferenceBuilder) Superglobal(superglobalName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewSuperGlobalToken(superglobalName), r.CreateContext(occursAtLine, dependencies.DependencyTypeSuperGlobalVariable)))
	return r
}

func (r *ReferenceBuilder) ReturnType(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, dependencies.DependencyTypeReturnType)))
	return r
}

func (r *ReferenceBuilder) ThrowStatement(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, dependencies.DependencyTypeThrow)))
	return r
}

func (r *ReferenceBuilder) AnonymousClassExtends(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, dependencies.DependencyTypeAnonymousClassExtends)))
	return r
}

func (r *ReferenceBuilder) AnonymousClassTrait(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, dependencies.DependencyTypeAnonymousClassTrait)))
	return r
}

func (r *ReferenceBuilder) ConstFetch(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, dependencies.DependencyTypeConst)))
	return r
}

func (r *ReferenceBuilder) AnonymousClassImplements(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, dependencies.DependencyTypeAnonymousClassImplements)))
	return r
}

func (r *ReferenceBuilder) Parameter(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, dependencies.DependencyTypeParameter)))
	return r
}

func (r *ReferenceBuilder) Attribute(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, dependencies.DependencyTypeAttribute)))
	return r
}

func (r *ReferenceBuilder) Instanceof(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, dependencies.DependencyTypeInstanceof)))
	return r
}

func (r *ReferenceBuilder) NewStatement(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, dependencies.DependencyTypeNew)))
	return r
}

func (r *ReferenceBuilder) StaticProperty(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, dependencies.DependencyTypeStaticProperty)))
	return r
}

func (r *ReferenceBuilder) StaticMethod(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, dependencies.DependencyTypeStaticMethod)))
	return r
}

func (r *ReferenceBuilder) CatchStmt(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, tokens.NewDependencyToken(tokens.NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, dependencies.DependencyTypeCatch)))
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
