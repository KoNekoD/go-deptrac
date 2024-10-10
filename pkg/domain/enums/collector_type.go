package enums

import "github.com/pkg/errors"

type CollectorType string

const (
	CollectorTypeTypeAttribute      CollectorType = "attribute"
	CollectorTypeTypeBool           CollectorType = "bool"
	CollectorTypeTypeClass          CollectorType = "struct"
	CollectorTypeTypeClasslike      CollectorType = "structLike"
	CollectorTypeTypeClassNameRegex CollectorType = "structNameRegex"
	CollectorTypeTypeTagValueRegex  CollectorType = "tagValueRegex"
	CollectorTypeTypeDirectory      CollectorType = "directory"
	CollectorTypeTypeExtends        CollectorType = "extends"
	CollectorTypeTypeFunctionName   CollectorType = "functionName"
	CollectorTypeTypeGlob           CollectorType = "glob"
	CollectorTypeTypeImplements     CollectorType = "implements"
	CollectorTypeTypeInheritance    CollectorType = "inheritanceLevel"
	CollectorTypeTypeInherits       CollectorType = "inherits"
	CollectorTypeTypeInterface      CollectorType = "interface"
	CollectorTypeTypeLayer          CollectorType = "layer_contract"
	CollectorTypeTypeMethod         CollectorType = "method"
	CollectorTypeTypeSuperGlobal    CollectorType = "superGlobal"
	CollectorTypeTypeGlobal         CollectorType = "global"
	CollectorTypeTypeTrait          CollectorType = "trait"
	CollectorTypeTypeUses           CollectorType = "uses"
	CollectorTypeTypePhpInternal    CollectorType = "php_internal"
	CollectorTypeTypeComposer       CollectorType = "composer"
)

var availableTypes = []CollectorType{
	CollectorTypeTypeAttribute,
	CollectorTypeTypeBool,
	CollectorTypeTypeClass,
	CollectorTypeTypeClasslike,
	CollectorTypeTypeClassNameRegex,
	CollectorTypeTypeTagValueRegex,
	CollectorTypeTypeDirectory,
	CollectorTypeTypeExtends,
	CollectorTypeTypeFunctionName,
	CollectorTypeTypeGlob,
	CollectorTypeTypeImplements,
	CollectorTypeTypeInheritance,
	CollectorTypeTypeInherits,
	CollectorTypeTypeInterface,
	CollectorTypeTypeLayer,
	CollectorTypeTypeMethod,
	CollectorTypeTypeSuperGlobal,
	CollectorTypeTypeGlobal,
	CollectorTypeTypeTrait,
	CollectorTypeTypeUses,
	CollectorTypeTypePhpInternal,
	CollectorTypeTypeComposer,
}

func NewCollectorTypeFromString(collectorType string) (CollectorType, error) {
	for _, availableType := range availableTypes {
		if string(availableType) == collectorType {
			return availableType, nil
		}
	}

	return "", errors.New("invalid collector type")
}
