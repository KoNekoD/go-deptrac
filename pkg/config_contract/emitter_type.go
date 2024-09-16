package config_contract

import "errors"

type EmitterType string

const (
	ClassToken               EmitterType = "class"
	ClassSuperGlobalToken    EmitterType = "class_superglobal"
	FileToken                EmitterType = "file_supportive"
	FunctionToken            EmitterType = "function"
	FunctionCall             EmitterType = "function_call"
	FunctionSuperGlobalToken EmitterType = "function_superglobal"
	UseToken                 EmitterType = "use"
)

func emitterTypeValues() []EmitterType {
	return []EmitterType{
		ClassToken,
		ClassSuperGlobalToken,
		FileToken,
		FunctionToken,
		FunctionCall,
		FunctionSuperGlobalToken,
		UseToken,
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
