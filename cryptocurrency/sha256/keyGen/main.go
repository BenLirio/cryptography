package main

import (
	"fmt"
	"crypto/sha256"
	"crypto/rand"
	"os"
)

var NUM_BLOCKS int64 = 256
var BYTES_PER_BLOCK int64 = 32
var NUM_KEYS int64 = 2

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter only filename")
		os.Exit(1)
	}
	writeKey(os.Args[1])
}

func writeKey(fileName string) {
	secretKeyFile, err := os.Create(fileName)
	check(err)
	defer secretKeyFile.Close()
	secretKeyBytes := make([]byte, BYTES_PER_BLOCK)
	var nthBlock int64
	for nthBlock = 0; nthBlock < (NUM_BLOCKS * NUM_KEYS); nthBlock++ {
		h := sha256.New()
		_, err = rand.Read(secretKeyBytes)
		var byteOffset int64 = nthBlock * BYTES_PER_BLOCK
		_, err = secretKeyFile.WriteAt(secretKeyBytes, byteOffset)
		check(err)
		h.Write(secretKeyBytes)
		byteOffset = BYTES_PER_BLOCK * NUM_BLOCKS * NUM_KEYS + byteOffset
		_, err = secretKeyFile.WriteAt(h.Sum(nil), byteOffset)
		check(err)
	}
}
