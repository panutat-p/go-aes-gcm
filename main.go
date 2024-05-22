package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/panutat-p/go-aes-gcm/aes_gcm"
)

var (
	encArg   string
	decArg   string
	helpText = `
Usage:
  key                Display the encryption key.
  enc <text>         Encrypt the provided text.
  dec <ciphertext>   Decrypt the provided ciphertext.
  help               Display this help message.
`
)

func main() {
	encCommand := flag.NewFlagSet("enc", flag.ExitOnError)
	decCommand := flag.NewFlagSet("dec", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'key', 'enc' or 'dec' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help":
		fmt.Print(helpText)
	case "key":
		fmt.Println("🗝️", os.Getenv("ENCRYPTION_KEY"))
	case "enc":
		err := encCommand.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Invalid argument, err:", err)
			return
		}
		if len(encCommand.Args()) > 0 {
			encArg = encCommand.Args()[0]
		} else {
			fmt.Println("No text to Encrypt")
			return
		}
		key, err := hex.DecodeString(os.Getenv("ENCRYPTION_KEY"))
		if err != nil {
			fmt.Println("No ENCRYPTION_KEY, err:", err)
			return
		}
		ciphertext, err := aes_gcm.Encrypt([]byte(encArg), key)
		if err != nil {
			fmt.Println("Error encrypting, err:", err)
			return
		}
		fmt.Printf("🔒 %x\n", ciphertext)
	case "dec":
		err := decCommand.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Invalid argument, err:", err)
			return
		}
		if len(decCommand.Args()) > 0 {
			decArg = decCommand.Args()[0]
		} else {
			fmt.Println("No text to Decrypt")
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
		decryptedText, err := aes_gcm.Decrypt(ciphertext, key)
		if err != nil {
			fmt.Println("Error decrypting:", err)
			return
		}
		fmt.Printf("📄 %s\n", decryptedText)
	default:
		fmt.Println("expected 'key', 'enc' or 'dec' subcommands")
		os.Exit(1)
	}
}
