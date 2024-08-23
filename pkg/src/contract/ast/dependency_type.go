package ast

// DependencyType - Specifies the type of AST dependency. You can use this information to enrich the displayed output to the user in your output formatter.
type DependencyType string

const (
	DependencyTypeUse                      DependencyType = "use"
	DependencyTypeInherit                  DependencyType = "inherit"
	DependencyTypeReturnType               DependencyType = "returntype"
	DependencyTypeParameter                DependencyType = "parameter"
	DependencyTypeNew                      DependencyType = "new"
	DependencyTypeStaticProperty           DependencyType = "static_property"
	DependencyTypeStaticMethod             DependencyType = "static_method"
	DependencyTypeInstanceof               DependencyType = "instanceof"
	DependencyTypeCatch                    DependencyType = "catch"
	DependencyTypeVariable                 DependencyType = "variable"
	DependencyTypeThrow                    DependencyType = "throw"
	DependencyTypeConst                    DependencyType = "const"
	DependencyTypeAnonymousClassExtends    DependencyType = "anonymous_class_extends"
	DependencyTypeAnonymousClassImplements DependencyType = "anonymous_class_implements"
	DependencyTypeAnonymousClassTrait      DependencyType = "anonymous_class_trait"
	DependencyTypeAttribute                DependencyType = "attribute"
	DependencyTypeSuperGlobalVariable      DependencyType = "superglobal_variable"
	DependencyTypeUnresolvedFunctionCall   DependencyType = "unresolved_function_call"
)
