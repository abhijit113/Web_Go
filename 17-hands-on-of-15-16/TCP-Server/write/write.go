package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	//creating a tcp server
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		//accept wait and returns the next connection to the listener
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		io.WriteString(conn, "Welcome from the TCP server\n")
		fmt.Fprint(conn, "How is your day?\n")
		fmt.Fprintf(conn, "Hope it's good!!!!\n")

	}

}
