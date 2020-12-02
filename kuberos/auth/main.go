package main

import (
	"fmt"
	"net"
	"crypto/rsa"
	"crypto/rand"
	"crypto"
)

func check(e error) {
	if e != nil {
		panic(e)	
	}
}

func handleConnection(conn net.Conn, publicKey rsa.PublicKey, privateKey *rsa.PrivateKey) {
	var err error
	req := make([]byte, 256)
	_, err = conn.Read(req)
	check(err)
	if string(req) == "GET public key" {
		_, err = conn.Write([]byte("E=" + string(privateKey.E) + "N=" + privateKey.N.String()))
		check(err)
	} else {
		decryptedReq, err := privateKey.Decrypt(nil, req, &rsa.OAEPOptions{Hash: crypto.SHA256})
		check(err)	
		fmt.Println(decryptedReq)
	}
}

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	check(err)
	publicKey := privateKey.PublicKey

	listener, err := net.Listen("tcp", ":4001")
	check(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConnection(conn, publicKey, privateKey)
	}
	
}
