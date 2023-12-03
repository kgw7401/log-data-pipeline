package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-faker/faker/v4"
)

// send_data
func sendMessage(eventName string, data []byte) {
	url := "http://localhost:3000/?event_name=" + eventName
	_, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func generateEvent(event interface{}) {
	err := faker.FakeData(&event)
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func GenerateEventRandomly(events []interface{}) {
	for {
		for _, event := range events {
			go generateEvent(event)
		}
		time.Sleep(5 * time.Second)
	}
}

func main() {
	events := []interface{}{HomeView{}, LoginDone{}}
	GenerateEventRandomly(events)
}
