package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, errS := net.Listen("tcp", ":8080")

	if errS != nil {
		log.Println(errS)
	}
	defer li.Close()

	for {
		connS, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleAtServer(connS)

	}
}

func handleAtServer(conn net.Conn) {

	//reading by server
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		r := rotateBy5(bs)

		//write by server
		fmt.Fprintf(conn, "%s - %s\n", ln, r)
	}
}

func rotateBy5(bs []byte) []byte {
	//the var is just declaration no assignmet , so : is not used here
	var r = make([]byte, len(bs))
	for i, v := range bs {
		if r[i] <= 117 {

			r[i] = v + 5
		} else {

			r[i] = v - 5
		}
	}
	return r
}

func dialerRead() {

}
