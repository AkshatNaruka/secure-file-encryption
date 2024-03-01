package main

import (
	"fmt"
	"log"
	"secure-file-encryption/cli"
	"github.com/manifoldco/promptui"
	"io/ioutil"
)

func main() {
	prompt := promptui.Select{
		Label: "Select Action",
		Items: []string{"Encrypt", "Decrypt"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	promptInput := promptui.Prompt{
		Label: "Input file path",
	}

	inputFile, err := promptInput.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	promptOutput := promptui.Prompt{
		Label: "Output file path",
	}

	outputFile, err := promptOutput.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	var key []byte

	switch result {
	case "Encrypt":
		key, err = cli.GenerateAESKey()
		if err != nil {
			log.Fatalf("Key generation failed: %v", err)
		}
		// Save the key to a file
		err = ioutil.WriteFile("key.txt", key, 0644)
		if err != nil {
			log.Fatalf("Failed to save key: %v", err)
		}
		err := cli.EncryptFile(inputFile, outputFile, key)
		if err != nil {
			log.Fatalf("Encryption failed: %v", err)
		}
		fmt.Println("File encrypted successfully.")
	case "Decrypt":
		// Read the key from the file
		key, err = ioutil.ReadFile("key.txt")
		if err != nil {
			log.Fatalf("Failed to read key: %v", err)
		}
		err := cli.DecryptFile(inputFile, outputFile, key)
		if err != nil {
			log.Fatalf("Decryption failed: %v", err)
		}
		fmt.Println("File decrypted successfully.")
	}
}