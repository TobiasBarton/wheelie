package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"wheelie/server/server"

	"github.com/google/uuid"
)

type Message struct {
	content string
	open    bool
}

func handle_connection(client server.Client, b *server.Broker) {
	defer client.Close()

	addr := conn.RemoteAddr().String()

	log.Println("Connected to client on:", addr)
	log.Println("Assigned client id:", id)

	var message Message
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err == io.EOF {
			log.Println("Disconnected from client on:", addr)
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
			if l == "START" {
				message = Message{content: "", open: true}
			} else if l == "END" {
				message.open = false
			} else if message.open {
				message.content += l
			}

			if !message.open {
				log.Println("Message received from: " + addr)
				log.Printf("Message: %s\n", message.content)
				args := strings.Split(message.content, " ")
				if args[0] == "DECLARE" {
					err := client.Declare(b, args[1], args[2])
					if err != nil {

					}
				}

				if client.Type == "NONE" {
					continue
				}

				if args[0] == "TOPIC" && args[1] == "CREATE" {
					b.CreateTopic(args[2])
				}

				if args[0] == "PUBLISH" && args[1] == "SEND" && client.Type == "publisher" {
					log.Println("Publishing message:" + args[3])
					err := b.Publish(client.Topic, args[3])
					if err != nil {
						conn.Write([]byte(err.Error()))
					}
				}

				log.Println(strings.Split(message.content, " "))
			}
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
	b := server.Broker{}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		client := server.Client{Id: uuid.NewString(), Type: "NONE"}

		// clients = append(clients, id)

		go handle_connection(client, &b)
	}
}
