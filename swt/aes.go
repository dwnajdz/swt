package swt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func encrypt(plainText, encodeKey []byte) (encoded string, err error) {
	block, err := aes.NewCipher(encodeKey)
	if err != nil {
		return
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))

	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.URLEncoding.EncodeToString(cipherText), err
}

func decrypt(secure string, encodeKey []byte) (decoded string, err error) {
	cipherText, err := base64.URLEncoding.DecodeString(secure)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(encodeKey)
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), err
}
