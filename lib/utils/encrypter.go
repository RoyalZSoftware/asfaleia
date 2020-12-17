// https://gist.github.com/miguelmota/3ea9286bd1d3c2a985b67cac4ba2130a
package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"log"
)

// PrivateKeyToBytes private key to bytes
func PrivateKeyToBytes(priv *rsa.PrivateKey) []byte {
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	return privBytes
}

// PublicKeyToBytes public key to bytes
func PublicKeyToBytes(pub *rsa.PublicKey) (error, []byte) {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return err, nil
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return nil, pubBytes
}

// BytesToPrivateKey bytes to private key
func BytesToPrivateKey(priv []byte) (error, *rsa.PrivateKey) {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return err, nil
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return err, nil
	}
	return nil, key
}

func BytesToPublicKey(pub []byte) (error, *rsa.PublicKey) {
	pemCert := "-----BEGIN PUBLIC KEY-----\n" +
		string(pub) + "\n-----END PUBLIC KEY-----"
	block, _ := pem.Decode([]byte(pemCert))
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes

	var err error
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return err, nil
		}
	}

	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return err, nil
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return errors.New("not ok"), nil
	}

	return nil, key
}

func EncodeInBase64(message []byte) string {
	return base64.StdEncoding.EncodeToString(message)
}

type RSAEncrypter struct {
	pubKey  *rsa.PublicKey
	privKey *rsa.PrivateKey
}

func (r RSAEncrypter) Encrypt(message []byte,
	pubKey *rsa.PublicKey) (error, []byte) {
	hash := sha1.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pubKey, message, nil)
	if err != nil {
		return err, nil
	}
	return nil, ciphertext
}

func (r RSAEncrypter) Decrypt(
	cipher []byte,
	privKey *rsa.PrivateKey,
) (error, []byte) {
	hash := sha1.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privKey, cipher, nil)
	if err != nil {
		return err, nil
	}
	return nil, plaintext
}
