package middlewares

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"os"
)

type CredEncryption struct {
	block cipher.Block
	err   error
}

// Ensure your AES key is the correct length (16, 24, or 32 bytes)
func CredEncrypt() *CredEncryption {
	block, err := aes.NewCipher([]byte(os.Getenv("ENCRYPTION_KEY")))

	return &CredEncryption{
		block: block,
		err:   err,
	}
}

func (cr *CredEncryption) Encrypt(data map[string]string) (string, error) {
	if cr.err != nil {
		return "", cr.err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Add padding
	padding := aes.BlockSize - len(jsonData)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	jsonData = append(jsonData, padtext...)

	ciphertext := make([]byte, aes.BlockSize+len(jsonData))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCBCEncrypter(cr.block, iv)
	stream.CryptBlocks(ciphertext[aes.BlockSize:], jsonData)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func (cr *CredEncryption) Decrypt(encryptedData string) (map[string]string, error) {
	if cr.err != nil {
		return nil, cr.err
	}

	ciphertext, err := base64.URLEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	stream := cipher.NewCBCDecrypter(cr.block, iv)
	stream.CryptBlocks(ciphertext, ciphertext)

	// Remove padding
	padding := int(ciphertext[len(ciphertext)-1])
	if padding < 1 || padding > aes.BlockSize {
		return nil, errors.New("invalid padding")
	}
	ciphertext = ciphertext[:len(ciphertext)-padding]

	var data map[string]string
	err = json.Unmarshal(ciphertext, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
