package utils

import (
	"fmt"
	"os"
)

// CheckFileExists checks if a file exists at the given path.
func CheckFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// CheckError prints an error message and exits if err is not nil.
func CheckError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
