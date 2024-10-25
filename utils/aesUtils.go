package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

var key = []byte("examplekey123456")

// AES加密
func Encrypt(str string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	strByte := []byte(str)
	// 使用 CBC 模式
	plainText := pkcs7Padding(strByte, block.BlockSize())
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return hex.EncodeToString(cipherText), nil
}

// AES解密
func Decrypt(decryptStr string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	cipherText, _ := hex.DecodeString(decryptStr)

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)
	padding, err := pkcs7UnPadding(cipherText)
	return string(padding), err
}

// PKCS7 填充
func pkcs7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

// PKCS7 去填充
func pkcs7UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])
	if unpadding > length {
		return nil, fmt.Errorf("padding size error")
	}
	return src[:(length - unpadding)], nil
}
