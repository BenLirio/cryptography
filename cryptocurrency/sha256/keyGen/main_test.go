package main

import (
	"crypto/sha256"
	"testing"
	"os"
)

func TestWriteKey(t *testing.T) {
	writeKey("key")

	keyFile, err := os.Open("key")
	check(err)
	privateBlock := make([]byte, 32)
	publicBlock := make([]byte, 32)
	var i int64
	for i = 0; i < (256 * 2); i++ {
		h := sha256.New()
		_, err = keyFile.ReadAt(privateBlock, 32 * i)
		check(err)
		_, err = keyFile.ReadAt(publicBlock, (i * 32) + 32 * 256 * 2)
		check(err)
		h.Write(privateBlock)
		for blockByteIndex := 0; blockByteIndex < 32; blockByteIndex++ {
			if h.Sum(nil)[blockByteIndex] != publicBlock[blockByteIndex] {
				t.Errorf("Block %d has not been hashed correctly", i)
			}
		}
	}
}
