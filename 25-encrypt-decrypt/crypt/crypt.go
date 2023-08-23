package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

var salt string

func GenerateSaltString() {
	bytes := generate32ByteKeyForAes256()
	salt = encodeByteToString(bytes)

	fmt.Printf("Salt:\n%s\n\n", salt)
}

func generate32ByteKeyForAes256() []byte {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}

	return bytes
}

func encodeByteToString(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

func Encrypt(stringToEncrypt string) (encryptedString string) {
	key := decodeStringToBytes(salt)
	plainText := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	cipherText := aesGCM.Seal(nonce, nonce, plainText, nil)

	return fmt.Sprintf("%x", cipherText)
}

func Decrypt(encryptedString string) (decryptedString string) {
	key := decodeStringToBytes(salt)
	enc := decodeStringToBytes(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, cipherText := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s", plainText)
}

func decodeStringToBytes(data string) []byte {
	key, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}

	return key
}
