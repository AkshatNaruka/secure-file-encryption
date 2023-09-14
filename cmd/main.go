package main

import (
	"flag"
	"fmt"
	"log"
	"secure-file-encryption/cli"
)

func main() {
	inputFile := flag.String("input", "", "Input file path")
	outputFile := flag.String("output", "", "Output file path")
	encrypt := flag.Bool("encrypt", false, "Encrypt the input file")
	decrypt := flag.Bool("decrypt", false, "Decrypt the input file")

	flag.Parse()

	if *encrypt == *decrypt || (*encrypt == false && *decrypt == false) {
		log.Fatal("Specify either -encrypt or -decrypt")
	}

	if (*encrypt && *decrypt) || (*encrypt && *inputFile == "") || (*decrypt && *inputFile == "") || *outputFile == "" {
		log.Fatal("Invalid flags or missing input/output file paths")
	}

	key, err := cli.GenerateAESKey()
	if err != nil {
		log.Fatalf("Key generation failed: %v", err)
	}

	if *encrypt {
		err := cli.EncryptFile(*inputFile, *outputFile, key)
		if err != nil {
			log.Fatalf("Encryption failed: %v", err)
		}
		fmt.Println("File encrypted successfully.")
	} else if *decrypt {
		err := cli.DecryptFile(*inputFile, *outputFile, key)
		if err != nil {
			log.Fatalf("Decryption failed: %v", err)
		}
		fmt.Println("File decrypted successfully.")
	}
}
