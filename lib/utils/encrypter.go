package utils

import "crypto/rsa"

type Encrypter interface {
	Encrypt(message string) string
	Decrypt(cipher string) string
}

type MockEncrypter struct {
}

func (m MockEncrypter) Encrypt(message string) string {
	return message
}

func (m MockEncrypter) Decrypt(cipher string) string {
	return cipher
}

type RSAEncrypter struct {
	pubKey  *rsa.PublicKey
	privKey *rsa.PrivateKey
}

func (r RSAEncrypter) Encrypt(message string) string {
	return ""
}

func (r RSAEncrypter) Decrypt(cipher string) string {
	return ""
}
