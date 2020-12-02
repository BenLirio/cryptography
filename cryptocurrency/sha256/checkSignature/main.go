package main

import (
	"os"
	"fmt"
	"crypto/sha256"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var BYTES_PER_BLOCK int64 = 32
var BITS_PER_BYTE int64 = 8
var NUM_BLOCKS int64 = 256
var NUM_KEYS int64 = 2
var FIRST_KEY_OFFSET int64 = NUM_BLOCKS * BYTES_PER_BLOCK * NUM_KEYS
var SECOND_KEY_OFFSET int64 = NUM_BLOCKS * BYTES_PER_BLOCK * (NUM_KEYS + 1)

func checkSignature(keyFileName string, signatureFileName string) {
	keyFile, err := os.Open(keyFileName)
	check(err)
	signatureFile, err := os.Open(signatureFileName)
	check(err)
	sigBuffer := make([]byte, BYTES_PER_BLOCK)
	sigHashBuffer := make([]byte, BYTES_PER_BLOCK)
	firstKeyBuffer := make([]byte, BYTES_PER_BLOCK)
	secondKeyBuffer := make([]byte, BYTES_PER_BLOCK)
	message := make([]byte, BYTES_PER_BLOCK)
	var blockIdx int64
	for blockIdx = 0; blockIdx < NUM_BLOCKS; blockIdx++ {
		_, err = signatureFile.ReadAt(sigBuffer, blockIdx * BYTES_PER_BLOCK)
		check(err)
		_, err = keyFile.ReadAt(firstKeyBuffer, FIRST_KEY_OFFSET + blockIdx * BYTES_PER_BLOCK)
		check(err)
		_, err = keyFile.ReadAt(secondKeyBuffer, SECOND_KEY_OFFSET + blockIdx * BYTES_PER_BLOCK)
		check(err)
		h := sha256.New()
		h.Write(sigBuffer)
		sigHashBuffer = h.Sum(nil)
		var sigEqualToFirstKey bool = true
		var sigEqualToSecondKey bool = true
		var byteIdx int64
		for byteIdx = 0; byteIdx < BYTES_PER_BLOCK; byteIdx++ {
			if sigHashBuffer[byteIdx] != firstKeyBuffer[byteIdx] {
				sigEqualToFirstKey = false
			}
			if sigHashBuffer[byteIdx] != secondKeyBuffer[byteIdx] {
				sigEqualToSecondKey = false
			}
		}
		if sigEqualToFirstKey == sigEqualToSecondKey {
			fmt.Println("The signature equivalence should be disjoint")
		}
		if sigEqualToFirstKey {
			message[blockIdx/BITS_PER_BYTE] = message[blockIdx/BITS_PER_BYTE] | (byte(1) << (blockIdx%BITS_PER_BYTE))
			fmt.Println(message[blockIdx/BITS_PER_BYTE])
		}
	}
	fmt.Println(message)
}

func main() {
	fmt.Println("Check Signature")
	if len(os.Args) != 3 {
		fmt.Println("[ key-file-name ] [ sig-file-name ]")
	}
	checkSignature(os.Args[1], os.Args[2])
}
