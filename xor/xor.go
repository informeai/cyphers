package xor

import (
	"encoding/hex"
	"errors"
	"fmt"
)

type Xor struct{}

func (x Xor) Encrypt(msg string, key string) (string, error) {
	var hash []byte
	if msg != "" && key != "" && len(key) > 0 {

		m := []byte(msg)
		k := []byte(key)
		for i := 0; i < len(m); i++ {
			hash = append(hash, m[i]^k[i%len(k)])
		}
		text := fmt.Sprintf("%x", hash)
		if text != "" {
			return text, nil
		}
	}
	return "", errors.New("failed encrypt")
}

func (x Xor) Decrypt(hash string, key string) (string, error) {
	if hash != "" && key != "" && len(key) > 0 {
		var decHash []byte
		decoded, err := hex.DecodeString(hash)
		if err != nil {
			return "", errors.New("failed read hash hex")
		}
		k := []byte(key)
		for i := 0; i < len(decoded); i++ {
			decHash = append(decHash, decoded[i]^k[i%len(k)])
		}
		text := fmt.Sprintf("%s", decHash)
		if err != nil {
			return "", errors.New("failed decrypt")
		}
		return text, nil
	}
	return "", errors.New("hash or key is empty")
}
