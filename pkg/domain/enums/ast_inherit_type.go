package enums

type AstInheritType string

const (
	AstInheritTypeExtends    AstInheritType = "Extends"
	AstInheritTypeImplements AstInheritType = "Implements"
	AstInheritTypeUses       AstInheritType = "Uses"
)
