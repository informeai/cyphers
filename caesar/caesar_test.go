package caesar

import "testing"

var msg string = "Julius Caesar"    // default -> Julius Caesar
var hot int = 5                     // default 10
var cypher string = "Ozqnzx%Hfjxfw" // default Ozqnzx%Hfjxfw

func TestMsgOrHotIsEmpty(t *testing.T) {
	if msg == "" || hot == 0 {
		t.Error("text or hot is empty")
	}
}
func TestEncrypt(t *testing.T) {
	caesar := New()
	text, err := caesar.Encrypt(msg, hot)
	if err != nil {
		t.Error("failed encrypted")
	}
	if text == "" {
		t.Error("text is empty")
	}
}
func TestDecrypt(t *testing.T) {
	caesar := New()
	text, err := caesar.Decrypt(cypher, hot)
	if err != nil {
		t.Error("failed decrypted")
	}
	if text == "" {
		t.Error("text is empty")
	}
}
