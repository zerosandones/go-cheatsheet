package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	Subject string
	Body    string
	time    time.Time //this value will not be converted into json because it is not exported
	From    string
	To      map[string]string
}

func main() {
	msg := Message{
		Subject: "Hello",
		Body:    "This is a test message",
		time:    time.Now(),
		From:    "dave@example.com",
		To:      map[string]string{"john@example.com": "John", "jane@example.com": "Jane"},
	}

	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	println(string(jsonBytes))

	var msg2 Message
	err = json.Unmarshal(jsonBytes, &msg2)
	if err != nil {
		panic(err)
	}
	fmt.Println(msg2.Subject, msg2.Body, msg2.time, msg2.From, msg2.To) //the time value is the default time value as the time valueis missing from the json

}
