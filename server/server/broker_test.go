package server

import (
	"fmt"
	"strings"
	"testing"
)

func TestTopicPush(t *testing.T) {
	topic := Topic{}
	topic.Push("hello")
	if len(topic.queue) != 1 {
		t.Error("Length of queue is not expected value!")
	}
}

func TestTopicPopSingleElement(t *testing.T) {
	topic := Topic{}
	expected := "hello"

	topic.Push(expected)

	data := topic.Pop()

	if data != expected {
		t.Errorf("%s != %s", data, expected)
	}
}

func TestParseMessage(t *testing.T) {
	// data := []byte("DECLARE PUBLISHER rabbits PUBLISH SEND \"data\"")

	// fmt.Printf("%q", data)

	// fmt.Println(strings.Index(string(data), "\""))

	// messages := []string{}
	// for _, message := range strings.Split(string(data), " ") {
	// 	messages = append(messages, message)
	// }

	// fmt.Println(len(messages))
	// fmt.Println(messages)

	message := "PUBLISH SEND \"{\"firstName\": \"Toby\"}\""

	split := strings.FieldsFunc(message, func(c rune) bool { return c == '"' })
	fmt.Println(split)
	fmt.Println(len(split))

	// indices := []int{}

	// part := message
	// for {
	// 	idx := strings.Index(part, "\"")
	// 	if idx == -1 {
	// 		break
	// 	}
	// 	indices = append(indices, idx)
	// 	part = part[idx+1:]
	// 	fmt.Println(idx)
	// 	fmt.Println(part)
	// }

	// fmt.Println(indices)
}
