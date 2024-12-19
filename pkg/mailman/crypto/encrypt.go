package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// Encrypt шифрует текст с использованием AES
func Encrypt(plainText, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	plainBytes := []byte(plainText)
	cipherText := make([]byte, aes.BlockSize+len(plainBytes))
	iv := cipherText[:aes.BlockSize]
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainBytes)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}
