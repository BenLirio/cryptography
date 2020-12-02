package main

import (
	"crypto/rsa"
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("PRIVATE KEY\n")
	fmt.Println("Private Exponent: ", privateKey.D)
	for _, prime := range privateKey.Primes {
		fmt.Println("Private Prime: ", prime)
	}
	publicKey := privateKey.PublicKey
	fmt.Printf("\n\n\nPUBLIC KEY\n")
	fmt.Println("Public Exponent: ", publicKey.E)
	fmt.Println("Public Key Modulus: ", publicKey.N)
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		[]byte("Super secret message"),
		nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("\n\n\nENCRYPTION\n")
	fmt.Println("Encrypted Bytes: ", encryptedBytes)
	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	fmt.Printf("\n\n\nDECRYPTION\n")
	fmt.Println("Decrypted Bytes: ", decryptedBytes)
	fmt.Println("Decrypted String: ", string(decryptedBytes))
}
