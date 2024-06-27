package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	const (
		host  = "127.0.0.1"
		port  = 8080
		topic = "rabbits"
	)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Println("Connected to:", conn.RemoteAddr().String())

	conn.Write([]byte("START\n"))
	conn.Write([]byte(fmt.Sprintf("TOPIC CREATE %s \n", topic)))
	conn.Write([]byte("END\n"))
	// conn.Write([]byte(fmt.Sprintf("TOPIC CREATE %s", topic)))
	conn.Write([]byte("START\n"))
	conn.Write([]byte(fmt.Sprintf("PUBLISH SEND %s \n", topic)))
	conn.Write([]byte("\"data\"\n"))
	conn.Write([]byte("END\n"))
}
