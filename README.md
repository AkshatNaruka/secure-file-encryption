# Secure File Encryption Tool

This is a simple command-line tool for securely encrypting and decrypting files using the Advanced Encryption Standard (AES) with Galois/Counter Mode (AES-GCM). It provides a secure way to protect your sensitive data.

## Features

- **File Encryption:** Encrypt your files using AES-GCM encryption for secure storage or transmission.
- **File Decryption:** Decrypt encrypted files back to their original state.
- **Key Generation:** Generate a secure AES encryption key for use in the encryption and decryption process.

## Prerequisites

Before using this tool, ensure you have the following installed:

- Go (Golang): [https://golang.org/](https://golang.org/)

## Usage

### Installation

Clone this repository to your local machine:

```bash
git clone https://github.com/akshat-naruka/secure-file-encryption.git
cd secure-file-encryption
```
## Building the Project
If you haven't built the project yet, you can do so with the following command:
```bash
go build
```
## Running the Tool
### To encrypt a file:
```bash
./secure-file-encryption -encrypt -input=input.txt -output=encrypted.txt
```
### To decrypt a file:
```bash
./secure-file-encryption -decrypt -input=encrypted.txt -output=decrypted.txt
```

Replace the file paths (input.txt, encrypted.txt, decrypted.txt) and options with your own.

## Example 
```bash
./secure-file-encryption -encrypt -input=example.txt -output=example_encrypted.txt


./secure-file-encryption -decrypt -input=example_encrypted.txt -output=example_decrypted.txt
```

# License
This project is licensed under the MIT License - see the LICENSE file for details.

# Acknowledgments
This project uses the Go programming language.
Encryption is performed using the AES-GCM encryption mode.
Key generation uses a secure random number generator.

Author <b>Akshat Naruka</b>

GitHub: https://github.com/AkshatNaruka

Website: https://akshatnaruka.netlify.app/



