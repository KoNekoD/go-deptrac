package enums

import "github.com/pkg/errors"

type EmitterType string

const (
	EmitterTypeClassToken               EmitterType = "class"
	EmitterTypeClassSuperGlobalToken    EmitterType = "class_superglobal"
	EmitterTypeFileToken                EmitterType = "file_supportive"
	EmitterTypeFunctionToken            EmitterType = "function"
	EmitterTypeFunctionCall             EmitterType = "function_call"
	EmitterTypeFunctionSuperGlobalToken EmitterType = "function_superglobal"
	EmitterTypeUseToken                 EmitterType = "use"
)

func emitterTypeValues() []EmitterType {
	return []EmitterType{
		EmitterTypeClassToken,
		EmitterTypeClassSuperGlobalToken,
		EmitterTypeFileToken,
		EmitterTypeFunctionToken,
		EmitterTypeFunctionCall,
		EmitterTypeFunctionSuperGlobalToken,
		EmitterTypeUseToken,
	}
}

func NewEmitterTypeFromString(input string) (EmitterType, error) {
	for _, emitterType := range emitterTypeValues() {
		if string(emitterType) == input {
			return emitterType, nil
		}
	}

	return "", errors.New("invalid emitter type string: " + input)
}
