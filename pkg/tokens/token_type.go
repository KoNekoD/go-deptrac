package tokens

import (
	"github.com/KoNekoD/go-deptrac/pkg/emitters"
)

type TokenType string

const (
	TokenTypeClassLike TokenType = "class-like"
	TokenTypeFunction  TokenType = "function"
	TokenTypeFile      TokenType = "file_supportive"
)

func NewTokenTypeTryFromEmitterType(emitterType emitters.EmitterType) *TokenType {
	if emitterType == emitters.EmitterTypeClassToken {
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
