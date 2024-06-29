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
	conn.Write([]byte("DECLARE PUBLISHER rabbits\n"))
	conn.Write([]byte("END\n"))

	// conn.Write([]byte("START\n"))
	// conn.Write([]byte(fmt.Sprintf("TOPIC CREATE %s\n", topic)))
	// conn.Write([]byte("END\n"))
	conn.Write([]byte("START\n"))
	conn.Write([]byte("PUBLISH SEND \n"))
	conn.Write([]byte("\"data\"\n"))
	conn.Write([]byte("END\n"))

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buffer[:n]))
}
