package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

type EncryptionService interface {
	EncryptData(data string, keySize uint) (string, string)
	DecryptData(data string, privateKey string) string
}

type encryptionService struct {
}

func NewEncryptionService() EncryptionService {
	return &encryptionService{}
}

func (e *encryptionService) EncryptData(data string, keySize uint) (string, string) {
	var privateKey, _ = rsa.GenerateKey(rand.Reader, int(keySize))
	publicKey := privateKey.PublicKey

	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		[]byte(data),
		nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	privateKeyBytes := ExportRsaPrivateKeyAsPemStr(privateKey)

	encodedPrivateKey := b64.StdEncoding.EncodeToString([]byte(privateKeyBytes))

	return string(encryptedBytes), encodedPrivateKey
}

func (e *encryptionService) DecryptData(data string, privateKeyStr string) string {
	decodedPrivateKey, _ := b64.StdEncoding.DecodeString(privateKeyStr)
	privateKey, _ := ParseRsaPrivateKeyFromPemStr(string(decodedPrivateKey))

	decryptedBytes, _ := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, []byte(data), nil)

	return string(decryptedBytes)
}

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privkey)
	privateKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKeyBytes,
		},
	)
	return string(privateKeyPem)
}

func ParseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
	fmt.Println(privPEM)
	fmt.Println("####################################")
	block, erro := pem.Decode([]byte(privPEM))
	fmt.Println(block)
	if block == nil {
		println(string(erro))
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}
