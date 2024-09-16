package config_contract

import "errors"

type CollectorType string

const (
	TypeAttribute      CollectorType = "attribute"
	TypeBool           CollectorType = "bool"
	TypeClass          CollectorType = "struct"
	TypeClasslike      CollectorType = "structLike"
	TypeClassNameRegex CollectorType = "structNameRegex"
	TypeTagValueRegex  CollectorType = "tagValueRegex"
	TypeDirectory      CollectorType = "directory"
	TypeExtends        CollectorType = "extends"
	TypeFunctionName   CollectorType = "functionName"
	TypeGlob           CollectorType = "glob"
	TypeImplements     CollectorType = "implements"
	TypeInheritance    CollectorType = "inheritanceLevel"
	TypeInherits       CollectorType = "inherits"
	TypeInterface      CollectorType = "interface"
	TypeLayer          CollectorType = "layer_contract"
	TypeMethod         CollectorType = "method"
	TypeSuperGlobal    CollectorType = "superGlobal"
	TypeGlobal         CollectorType = "global"
	TypeTrait          CollectorType = "trait"
	TypeUses           CollectorType = "uses"
	TypePhpInternal    CollectorType = "php_internal"
	TypeComposer       CollectorType = "composer"
)

var availableTypes = []CollectorType{
	TypeAttribute,
	TypeBool,
	TypeClass,
	TypeClasslike,
	TypeClassNameRegex,
	TypeTagValueRegex,
	TypeDirectory,
	TypeExtends,
	TypeFunctionName,
	TypeGlob,
	TypeImplements,
	TypeInheritance,
	TypeInherits,
	TypeInterface,
	TypeLayer,
	TypeMethod,
	TypeSuperGlobal,
	TypeGlobal,
	TypeTrait,
	TypeUses,
	TypePhpInternal,
	TypeComposer,
}

func NewCollectorTypeFromString(collectorType string) (CollectorType, error) {
	for _, availableType := range availableTypes {
		if string(availableType) == collectorType {
			return availableType, nil
		}
	}

	return "", errors.New("invalid collector type")
}
