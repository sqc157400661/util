package util

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	a := NewEncryptor()
	aa := "sqctest"
	fmt.Println(a.Encrypt(aa))

	bb := "T15+2TXdFynaKaRXGrk8vA=="

	bb1, _ := a.Decrypt(bb)

	fmt.Println(bb1)
	//fmt.Println(a.Decrypt(string(decodedBytes)))
}
