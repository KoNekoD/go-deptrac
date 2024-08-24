package analyser

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/config"
)

type TokenType string

const (
	TokenTypeClassLike TokenType = "class-like"
	TokenTypeFunction  TokenType = "function"
	TokenTypeFile      TokenType = "file"
)

func NewTokenTypeTryFromEmitterType(emitterType config.EmitterType) *TokenType {
	if emitterType == config.ClassToken {
		classLikeTokenType := TokenTypeClassLike
		return &classLikeTokenType
	} else {
		v := string(emitterType)

		allowed := []string{
			string(TokenTypeClassLike),
			string(TokenTypeFunction),
			string(TokenTypeFile),
		}

		for _, allowedValue := range allowed {
			if v == allowedValue {
				t := TokenType(v)
				return &t
			}
		}

		return nil
	}
}
