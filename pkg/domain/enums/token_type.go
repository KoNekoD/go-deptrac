package enums

import "github.com/pkg/errors"

type TokenType string

const (
	TokenTypeClassLike TokenType = "class-like"
	TokenTypeFunction  TokenType = "function"
	TokenTypeFile      TokenType = "file_supportive"
)

func NewTokenType(s string) (TokenType, error) {
	switch s {
	case string(TokenTypeClassLike):
		return TokenTypeClassLike, nil
	case string(TokenTypeFunction):
		return TokenTypeFunction, nil
	case string(TokenTypeFile):
		return TokenTypeFile, nil
	default:
		return "", errors.New("invalid token type string: " + s)
	}
}

func NewTokenTypeTryFromEmitterType(emitterType EmitterType) *TokenType {
	if emitterType == EmitterTypeClassToken {
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
