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

var NUM_BLOCKS int64 = 256
var BYTES_PER_BLOCK int64 = 32
var BITS_PER_BYTE int64 = 8

func sign(keyFileName string, message string) ([]byte, error) {
	var signature []byte = make([]byte, NUM_BLOCKS * BYTES_PER_BLOCK)
	var buffer []byte = make([]byte, BYTES_PER_BLOCK)
	keyFile, err := os.Open(keyFileName)
	check(err)


	h := sha256.New()
	h.Write([]byte(message))
	hashedMessage := h.Sum(nil)
	for byteIdx, messageByte := range hashedMessage {
		var mask byte = 128 // [10000000] - bit representation
		var bitIdx int64
		for bitIdx = 0; bitIdx < BITS_PER_BYTE; bitIdx++ {
			var nthBit int64 = int64(byteIdx) * BITS_PER_BYTE + bitIdx
			var readOffset int64 = nthBit * BYTES_PER_BLOCK
			if messageByte & mask == mask {
				readOffset = readOffset + NUM_BLOCKS * BYTES_PER_BLOCK
			}
			keyFile.ReadAt(buffer, readOffset) // reads unit is per byte
			copy(signature[nthBit * BYTES_PER_BLOCK:nthBit * BYTES_PER_BLOCK + BYTES_PER_BLOCK], buffer)
		}
	}
	return signature, nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("[ key-file ] [ signature-file ] [ message ]")
		os.Exit(1)
	}
	signature, err := sign(os.Args[1], os.Args[3])
	check(err)
	signatureFile, err := os.Create(os.Args[2])
	defer signatureFile.Close()
	signatureFile.Write(signature)
}
