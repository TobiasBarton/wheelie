package server

import (
	"fmt"
	"net"
)

type Topic struct {
	name  string
	queue []string
}

func (t *Topic) Push(data string) {
	t.queue = append(t.queue, data)
}

func (t *Topic) Pop() string {
	idx := len(t.queue) - 1
	data := t.queue[len(t.queue)-1]
	t.queue = append(t.queue[:idx], t.queue[idx+1:]...)
	return data
}

type Client struct {
	Id    string
	Type  string
	Topic string
	conn  net.Conn
}

func (c *Client) Declare(b *Broker, t string, topic string) error {
	if !b.CheckTopicExists(topic) {
		return fmt.Errorf("Topic does not exist: %s", topic)
	}

	c.Type = t
	c.Topic = topic

	return nil
}

// func (c *Client) Write(s string) (int, error) {
// 	return conn.
// }

func (c *Client) IsValid() bool {
	return c.Type != "NONE"
}

func (c *Client) Close() {
	c.conn.Close()
}

func InitClient(conn net.Conn) (Client, error) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return Client{}, err
	}
	fmt.Println(buffer[:n])

	return Client{}, nil
}

type Broker struct {
	topics  []Topic
	clients []string
}

func (b *Broker) CreateTopic(topic string) {
	for _, t := range b.topics {
		if t.name == topic {
			return
		}
	}

	b.topics = append(b.topics, Topic{name: topic})
}

func (b *Broker) CheckTopicExists(topic string) bool {
	for _, t := range b.topics {
		if t.name == topic {
			return true
		}
	}

	return false
}

func (b *Broker) Publish(topic string, data string) error {
	for _, t := range b.topics {
		if t.name == topic {
			t.Push(data)
			return nil
		}
	}

	return fmt.Errorf("Topic does not exist: %s", topic)
}
