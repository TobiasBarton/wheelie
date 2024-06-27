package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"slices"
	"strings"

	"github.com/google/uuid"
)

type Message struct {
	content string
	open    bool
}

type Broker struct {
	topics  []string
	clients []string
}

func (b *Broker) create_topic(topic string) {
	if slices.Contains(b.topics, topic) {
		return
	}

	b.topics = append(b.topics, topic)
}

func handle_connection(conn net.Conn, id string, broker *Broker) {
	defer conn.Close()

	client := conn.RemoteAddr().String()

	log.Println("Connected to client on:", client)
	log.Println("Assigned client id:", id)

	var message Message
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err == io.EOF {
			log.Println("Disconnected from client on:", client)
			return
		}
		if err != nil {
			log.Fatal(err)
		}

		if n == 0 {
			continue
		}

		lines := bytes.Split(buffer, []byte("\n"))
		for _, line := range lines {
			l := string(line)
			log.Println(l)
			if l == "START" {
				message = Message{content: "", open: true}
			} else if l == "END" {
				message.open = false
			} else if message.open {
				message.content += l
			}
		}

		if !message.open {
			log.Println("Message received from: " + client)
			log.Printf("Message: %s\n", message.content)
			args := strings.Split(message.content, " ")
			if args[0] == "TOPIC" && args[1] == "CREATE" {
				broker.create_topic(args[2])
			}

			log.Println(strings.Split(message.content, " "))
			log.Println(len(strings.Split(message.content, " ")))
		}
	}
}

func main() {
	const (
		host = "127.0.0.1"
		port = 8080
	)

	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Printf("Listening on: %s:%d\n", host, port)

	// clients := []string{}
	// topics := []string{}
	broker := Broker{}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		id := uuid.New().String()
		// clients = append(clients, id)

		go handle_connection(conn, id, &broker)
	}
}
