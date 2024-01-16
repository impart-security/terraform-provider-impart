package provider

import (
	"crypto/sha256"
	"encoding/hex"
)

func calculateSha256(input string) (string, error) {
	h := sha256.New()

	_, err := h.Write([]byte(input))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
