package aes_gcm

import (
	"encoding/base64"
	"errors"
)

func CheckKey(keyB64 string) error {
	if keyB64 == "" {
		return errors.New("key is empty")
	}
	key, err := base64.StdEncoding.DecodeString(keyB64)
	if err != nil {
		return err
	}
	if len(key) != 32 {
		return errors.New("key length is not 32 bytes")
	}
	return nil
}

func Encrypt(keyB64 string, plainText string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(keyB64)
	if err != nil {
		return "", err
	}
	if len(key) != 32 {
		return "", errors.New("key length is not 32 bytes")
	}
	ciphertext, err := encryptAesGcm(key, []byte(plainText))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(keyB64 string, cipherTextB64 string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(keyB64)
	if err != nil {
		return "", err
	}
	if len(key) != 32 {
		return "", errors.New("key length is not 32 bytes")
	}
	cipherText, err := base64.StdEncoding.DecodeString(cipherTextB64)
	if err != nil {
		return "", err
	}
	plainText, err := decryptAesGcm(key, cipherText)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}
