package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyContext"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/FileOccurrence"
)

type ReferenceBuilder struct {
	Dependencies   []*DependencyToken
	tokenTemplates []string
	Filepath       string
	ref            *FileReference
}

type ReferenceBuilderInterface interface {
	GetTokenTemplates() []string
	CreateContext(occursAtLine int, dependencyType DependencyType.DependencyType) *DependencyContext.DependencyContext
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
		Dependencies:   make([]*DependencyToken, 0),
		tokenTemplates: tokenTemplates,
		Filepath:       filepath,
	}
}

func (r *ReferenceBuilder) GetTokenTemplates() []string {
	return r.tokenTemplates
}

func (r *ReferenceBuilder) CreateContext(occursAtLine int, dependencyType DependencyType.DependencyType) *DependencyContext.DependencyContext {
	return DependencyContext.NewDependencyContext(FileOccurrence.NewFileOccurrence(r.Filepath, occursAtLine), dependencyType)
}

// UnresolvedFunctionCall - Unqualified function and constant names inside a namespace cannot be statically resolved. Inside a namespace Foo, a call to strlen() may either refer to the namespaced \Foo\strlen(), or the global \strlen(). Because PHP-ParserInterface does not have the necessary context to decide this, such names are left unresolved.
func (r *ReferenceBuilder) UnresolvedFunctionCall(functionName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewFunctionTokenFromFQCN(functionName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeUnresolvedFunctionCall)))
	return r
}

func (r *ReferenceBuilder) Variable(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeVariable)))
	return r
}

func (r *ReferenceBuilder) Superglobal(superglobalName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewSuperGlobalToken(superglobalName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeSuperGlobalVariable)))
	return r
}

func (r *ReferenceBuilder) ReturnType(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeReturnType)))
	return r
}

func (r *ReferenceBuilder) ThrowStatement(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeThrow)))
	return r
}

func (r *ReferenceBuilder) AnonymousClassExtends(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeAnonymousClassExtends)))
	return r
}

func (r *ReferenceBuilder) AnonymousClassTrait(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeAnonymousClassTrait)))
	return r
}

func (r *ReferenceBuilder) ConstFetch(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeConst)))
	return r
}

func (r *ReferenceBuilder) AnonymousClassImplements(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeAnonymousClassImplements)))
	return r
}

func (r *ReferenceBuilder) Parameter(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeParameter)))
	return r
}

func (r *ReferenceBuilder) Attribute(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeAttribute)))
	return r
}

func (r *ReferenceBuilder) Instanceof(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeInstanceof)))
	return r
}

func (r *ReferenceBuilder) NewStatement(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeNew)))
	return r
}

func (r *ReferenceBuilder) StaticProperty(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeStaticProperty)))
	return r
}

func (r *ReferenceBuilder) StaticMethod(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeStaticMethod)))
	return r
}

func (r *ReferenceBuilder) CatchStmt(classLikeName string, occursAtLine int) *ReferenceBuilder {
	r.Dependencies = append(r.Dependencies, NewDependencyToken(NewClassLikeTokenFromFQCN(classLikeName), r.CreateContext(occursAtLine, DependencyType.DependencyTypeCatch)))
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
