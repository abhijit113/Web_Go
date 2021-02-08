package main

import (
	"bufio"
	"fmt"
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
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		//print every line
		fmt.Println(ln)
	}
	defer conn.Close()
}
