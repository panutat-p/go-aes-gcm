package aes_gcm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

const (
	NOUNCE_SIZE = 12
)

func encryptAesGcm(key []byte, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, NOUNCE_SIZE)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nonce, nonce, plainText, nil)
	return ciphertext, nil
}

func decryptAesGcm(key []byte, ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < NOUNCE_SIZE {
		return nil, errors.New("ciphertext is too short")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce, ciphertext := ciphertext[:NOUNCE_SIZE], ciphertext[NOUNCE_SIZE:]
	plainText, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
