package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func getKeyBytes(key []byte) []byte {
	l := len(key)
	keyBytes := make([]byte, l)
	copy(keyBytes, key)
	switch {
	case l < aes.BlockSize:
		keyBytes = append(keyBytes, make([]byte, aes.BlockSize-l)...)
	case l > aes.BlockSize:
		keyBytes = key[:aes.BlockSize]
	}
	return keyBytes
}

func AESCBCEncrypt(key []byte, encodeStr string) ([]byte, error) {
	plaintext := []byte(encodeStr)
	plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	//根据key 生成密文
	keyBytes := getKeyBytes(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// AESCBCDecrypt from base64 to decrypted string
func AESCBCDecrypt(key []byte, decodeBytes []byte) (string, error) {
	keyBytes := getKeyBytes(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	iv := decodeBytes[:blockSize]
	ciphertext := decodeBytes[blockSize:]
	origData := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(origData, ciphertext)
	origData = PKCS5UnPadding(origData)
	return string(origData), nil
}

func AESCBCEncryptZeroIV(plaintext []byte) ([]byte, error) {
	//plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	//根据key 生成密文
	key := bytes.Repeat([]byte{0x0}, aes.BlockSize)
	keyBytes := getKeyBytes(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	for idx, _ := range iv {
		iv[idx] = 0x00
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}
