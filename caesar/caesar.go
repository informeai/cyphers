package caesar

import (
	"errors"
	"fmt"
)

type Caesar struct{}

func (c Caesar) Encrypt(msg string, hot int) (string, error) {
	if msg != "" && hot != 0 {

		var textBytes []byte
		m := []byte(msg)
		h := byte(hot)

		for i := 0; i < len(m); i++ {
			textBytes = append(textBytes, m[i]+h)

		}
		if len(textBytes) > 0 {
			return fmt.Sprintf("%s", textBytes), nil
		}
		return "", errors.New("failed encrypted")
	}
	return "", errors.New("msg or hot is empty")
}
func (c Caesar) Decrypt(cypher string, hot int) (string, error) {
	if cypher != "" && hot != 0 {

		var cypherBytes []byte
		m := []byte(cypher)
		h := byte(hot)

		for i := 0; i < len(m); i++ {
			cypherBytes = append(cypherBytes, m[i]-h)

		}
		if len(cypherBytes) > 0 {
			return fmt.Sprintf("%s", cypherBytes), nil
		}
		return "", errors.New("failed encrypted")
	}
	return "", errors.New("cypher or hot is empty")
}
