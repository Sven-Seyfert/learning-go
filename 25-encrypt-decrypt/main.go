package main

import (
	"25-encrypt-decrypt/crypt"

	"fmt"
)

func main() {
	crypt.GenerateSaltString()

	const dataToEncrypt = `C:\Store\Repositories\GitHub\learning-go\25-encrypt-decrypt\main.go`
	fmt.Printf("DataToEncrypt:\n%s\n\n", dataToEncrypt)

	encrypted := crypt.Encrypt(dataToEncrypt)
	fmt.Printf("Encrypted:\n%s\n\n", encrypted)

	decrypted := crypt.Decrypt(encrypted)
	fmt.Printf("Decrypted:\n%s\n\n", decrypted)
}
