package main

import (
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
  enc <text>         encryptAesGcm the provided text.
  dec <ciphertext>   decryptAesGcm the provided ciphertext.
  help               Display this help message.
`
)

func main() {
	encCommand := flag.NewFlagSet("enc", flag.ExitOnError)
	decCommand := flag.NewFlagSet("dec", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Print(helpText)
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
		if len(encCommand.Args()) == 0 {
			fmt.Println("No argument provided")
			os.Exit(1)
		}
		encArg = encCommand.Args()[0]
		key := os.Getenv("ENCRYPTION_KEY")
		if key == "" {
			fmt.Println("No ENCRYPTION_KEY provided")
			os.Exit(1)
		}
		ciphertext, err := aes_gcm.Encrypt(key, encArg)
		if err != nil {
			fmt.Println("Failed to Encrypt, err:", err)
			return
		}
		fmt.Printf("🔒 %s\n", ciphertext)
	case "dec":
		err := decCommand.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Invalid argument, err:", err)
			os.Exit(1)
		}
		if len(decCommand.Args()) == 0 {
			fmt.Println("No argument provided")
			os.Exit(1)
		}
		decArg = decCommand.Args()[0]
		key := os.Getenv("ENCRYPTION_KEY")
		if key == "" {
			fmt.Println("No ENCRYPTION_KEY provided")
			os.Exit(1)
		}
		decryptedText, err := aes_gcm.Decrypt(key, decArg)
		if err != nil {
			fmt.Println("Failed to Decrypt:", err)
			os.Exit(1)
		}
		fmt.Printf("📄 %s\n", decryptedText)
	default:
		fmt.Print(helpText)
		os.Exit(1)
	}
}
