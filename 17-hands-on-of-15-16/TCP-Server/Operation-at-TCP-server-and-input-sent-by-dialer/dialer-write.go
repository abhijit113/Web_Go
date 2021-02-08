package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	connW, errW := net.Dial("tcp", "localhost:8080")

	if errW != nil {
		log.Fatalln(errW)
	}
	defer connW.Close()

	//writing here
	fmt.Fprintln(connW, "Writing to TCP server")
}
