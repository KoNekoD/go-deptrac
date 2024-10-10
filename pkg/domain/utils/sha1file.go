package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
)

func Sha1File(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := sha1.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil)), nil
}
