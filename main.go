package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func encrypt(plainText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, 12)
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

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	if len(ciphertext) < 12 {
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

	nonce, ciphertext := ciphertext[:12], ciphertext[12:]
	plainText, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

var (
	encArg string
	decArg string
)

func main() {
	encCommand := flag.NewFlagSet("enc", flag.ExitOnError)
	decCommand := flag.NewFlagSet("dec", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'key', 'enc' or 'dec' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "key":
		fmt.Println("🗝️", os.Getenv("ENCRYPTION_KEY"))
	case "enc":
		encCommand.Parse(os.Args[2:])
		if len(encCommand.Args()) > 0 {
			encArg = encCommand.Args()[0]
		} else {
			fmt.Println("No text to encrypt")
			return
		}
		key, err := hex.DecodeString(os.Getenv("ENCRYPTION_KEY"))
		if err != nil {
			fmt.Println("No ENCRYPTION_KEY, err:", err)
			return
		}
		ciphertext, err := encrypt([]byte(encArg), key)
		if err != nil {
			fmt.Println("Error encrypting:", err)
			return
		}
		fmt.Printf("🔒 %x\n", ciphertext)
	case "dec":
		decCommand.Parse(os.Args[2:])
		if len(decCommand.Args()) > 0 {
			decArg = decCommand.Args()[0]
		} else {
			fmt.Println("No text to decrypt")
			return
		}
		key, err := hex.DecodeString(os.Getenv("ENCRYPTION_KEY"))
		if err != nil {
			fmt.Println("No ENCRYPTION_KEY, err:", err)
			return
		}
		ciphertext, err := hex.DecodeString(decArg)
		if err != nil {
			fmt.Println("Error decoding ciphertext:", err)
			return
		}
		decryptedText, err := decrypt(ciphertext, key)
		if err != nil {
			fmt.Println("Error decrypting:", err)
			return
		}
		fmt.Printf("🔓 %s\n", decryptedText)
	default:
		fmt.Println("expected 'key', 'enc' or 'dec' subcommands")
		os.Exit(1)
	}
}
