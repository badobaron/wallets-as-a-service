package common

import (
	"testing"
)

//ToDo: Split this function into seperate encrypt and decrypt functions
func TestEncryptAndDescrypt(t *testing.T){
	// Encrypt
	input, expected := "0e6372d9fc09b0b345ed4a8f9477d0b12c6c5b1ff7f352c4a53cf79ee3d10f42", "0e6372d9fc09b0b345ed4a8f9477d0b12c6c5b1ff7f352c4a53cf79ee3d10f42"
	result, _ := Encrypt([]byte("1234567891234567"), []byte(input))

	// Decrypt
	result, _ = Decrypt([]byte("1234567891234567"), result)
	resultString := string(result[:])


	if resultString != expected {
		t.Errorf("expected: " + expected + " is not given " + resultString)
	}
}
