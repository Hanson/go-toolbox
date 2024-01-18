package utils

import (
	"encoding/hex"
	"github.com/hanson/go-toolbox/utils/crypto"
)

func Encrypt(key string, plainText []byte) ([]byte, error) {
	binKey, _ := hex.DecodeString(key)
	aes := crypto.NewAesCBCPKCS7()
	xor_data := crypto.XorCrypto(plainText, binKey)
	cipherText, err := aes.Encrypt(xor_data, binKey, binKey)
	return cipherText, err
}

func Decrypt(key string, cipherText []byte) ([]byte, error) {
	binKey, _ := hex.DecodeString(key)
	aes := crypto.NewAesCBCPKCS7()
	xor_data, err := aes.Decrypt(cipherText, binKey, binKey)
	return crypto.XorCrypto(xor_data, binKey), err
}
