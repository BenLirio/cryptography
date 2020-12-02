package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("client")
	conn, err := net.Dial("tcp", "localhost:4001")
}
