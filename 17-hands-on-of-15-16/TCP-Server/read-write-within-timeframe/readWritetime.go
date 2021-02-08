package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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

//reading and writing by the server
func handle(conn net.Conn) {
	//setting up the timer

	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalln("CONN TIMEOUT")
	}

	//reading starts
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		//print every line
		fmt.Println(ln)
		//writing starts
		fmt.Fprintf(conn, "I heard what you have said, %s\n", ln)
	}
	defer conn.Close()
}
