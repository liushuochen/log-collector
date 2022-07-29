// Package encrypt used to AES encrypt for data.
package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"identify/exception"
)

const (
	// DatabaseUserKey used for the pwd of database user
	DatabaseUserKey = "kfAxW3XgA6cr4zxpZa913SZbRmGlSghK"
)

// Encrypt returns a cipher and an error.
func Encrypt(password, key string) (string, error) {
	keyBytes, err := checkKey(key)
	if err != nil {
		return "", err
	}

	return aesEncrypt([]byte(password), keyBytes)
}

// UnEncrypt resource a source string and an error.
func UnEncrypt(ciphertext, key string) (string, error) {
	keyBytes, err := checkKey(key)
	if err != nil {
		return "", exception.NewInvalidKeyLengthError(key)
	}

	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	data, err := aesDcrypt(ciphertextBytes, keyBytes)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func checkKey(key string) ([]byte, error) {
	if len(key) % 16 != 0 {
		return nil, exception.NewInvalidKeyLengthError(key)
	}
	return []byte(key), nil
}

func aesEncrypt(data, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()

	data = pKCS7Padding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func aesDcrypt(crypted, keyBytes []byte) ([]byte, error) {
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyBytes[:blockSize])
	data := make([]byte, len(crypted))
	blockMode.CryptBlocks(data, crypted)
	return pKCS7UnPadding(data), nil
}

func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func pKCS7UnPadding(originalData []byte) []byte {
	length := len(originalData)
	unPadding := int(originalData[length-1])
	return originalData[:length - unPadding]
}
