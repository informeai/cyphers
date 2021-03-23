package xor

import (
	"testing"
)

var msg string = "informeai"           // default -> informeai
var key string = "secret"              // default -> secret
var hash string = "1a0b051d171916040a" // default 1a0b051d171916040a
func TestMsgOrKeyEmpty(t *testing.T) {
	if msg == "" {
		t.Error("error: msg is empty")
	}
	if key == "" {
		t.Error("error: key is empty")
	}
}

func TestEncrypt(t *testing.T) {
	x := Xor{}
	hash, err := x.Encrypt(msg, key)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	if hash == "" {
		t.Errorf("error: hash is empty.")
	}
	if hash != "" {
		t.Log(hash)
	}

}

func TestDecrypt(t *testing.T) {
	x := Xor{}
	hash, err := x.Decrypt(hash, key)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	if hash == "" {
		t.Errorf("error: hash is empty.")
	}
}
