package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// Decrypt расшифровывает текст с использованием AES
func Decrypt(cipherText, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	cipherBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	plainBytes := make([]byte, len(cipherBytes[aes.BlockSize:]))
	iv := cipherBytes[:aes.BlockSize]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plainBytes, cipherBytes[aes.BlockSize:])

	return string(plainBytes), nil
}
