package gateway

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type CryptoGw interface {
	GenerateHash(bytes []byte) string
	Encode(b []byte) string
	Decode(s string) ([]byte, error)
	Encrypt(text, MySecret string) (string, error)
	Decrypt(text, MySecret string) (string, error)
}

type CryptoGateway struct {
	secretString string
	cryptoBytes  []byte
}

func NewCryptoGateway(secretString string) *CryptoGateway {
	return &CryptoGateway{
		secretString: secretString,
		cryptoBytes:  []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05},
	}
}

func (gw *CryptoGateway) GenerateHash(bytes []byte) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", bytes)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (gw *CryptoGateway) Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (gw *CryptoGateway) Decode(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Encrypt method is to encrypt or hide any classified text
func (gw *CryptoGateway) Encrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, gw.cryptoBytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return gw.Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func (gw *CryptoGateway) Decrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText, err := gw.Decode(text)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, gw.cryptoBytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}
