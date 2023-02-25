package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"

	"github.com/mergermarket/go-pkcs7"
)

// Decrypt decrypts an string.
func Decrypt(encrypted, key string) (string, error) {
	cipherText, _ := hex.DecodeString(encrypted)

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		log.Fatalln(err)
	}

	if len(cipherText) < aes.BlockSize {
		log.Fatalln("cipherText too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	if len(cipherText)%aes.BlockSize != 0 {
		log.Fatalln("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)
	return string(cipherText), nil
}
