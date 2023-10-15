package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

// func sendSlack(url string, b []byte) {

// 	attachment1 := slack.Attachment{}
// 	attachment1.AddAction(slack.Action{Type: "button", Text: "resolve", Confirm: &slack.Confirm{Title: "Are you sure?", Text: "Wouldn't you prefer a good game of chess?", OkText: "Yes", DismissText: "No"}})
// 	// attachment1.AddAction(slack.Action { Type: "button", Text: "Cancel", Url: "https://flights.example.com/abandon/r123456", Style: "danger" })

// 	attachment1.AddField(slack.Field{Title: "스키마 오류 감지", Value: string(b)})

// 	payload := slack.Payload{
// 		Text:        "*[스키마 오류]*\n",
// 		Attachments: []slack.Attachment{attachment1},
// 	}
// 	payload.SendSlack(url)
// }

func sendSlack(token string, channel string, e string, b []byte) {
	api := slack.New(token)
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	attachment1 := slack.Attachment{
		Title:   "스키마 오류",
		Pretext: "에러 메시지",
		Text:    e,
	}

	attachment2 := slack.Attachment{
		Title:      "스키마 오류",
		Fallback:   "선택 오류",
		CallbackID: "schema_error",
		Pretext:    "페이로드",
		Text:       "```" + out.String() + "```",
		// Actions: []slack.AttachmentAction{
		// 	{
		// 		Name: "Resolve",
		// 		Text: "Resolve",
		// 		Type: slack.ActionType("button"),
		// 		Confirm: &slack.ConfirmationField{
		// 			Title:       "Resolve Issue",
		// 			Text:        "해결 처리 하시겠습니까?",
		// 			OkText:      "Yes",
		// 			DismissText: "No",
		// 		},
		// 		Value: "resolve",
		// 	},
		// },
	}
	message := slack.MsgOptionAttachments(attachment1, attachment2)
	channelID, timestamp, err := api.PostMessage(channel, slack.MsgOptionText("", false), message)
	if err != nil {
		fmt.Printf("Could not send message: %v", err)
	}
	fmt.Printf("Message with buttons sucessfully sent to channel %s at %s", channelID, timestamp)
}

// func main() {
// 	c, err := kafka.NewConsumer(&kafka.ConfigMap{
// 		"bootstrap.servers": "localhost:9092",
// 		"group.id":          "eventConsumerGroup",
// 		"auto.offset.reset": "earliest",
// 	})

// 	if err != nil {
// 		panic(err)
// 	}

// 	sigchan := make(chan os.Signal, 1)
// 	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

// 	if err != nil {
// 		fmt.Printf("Failed to create schema registry client: %s\n", err)
// 		os.Exit(1)
// 	}

// 	if err != nil {
// 		fmt.Printf("Failed to create deserializer: %s\n", err)
// 		os.Exit(1)
// 	}

// 	c.SubscribeTopics([]string{"dlt.event"}, nil)

// 	err = godotenv.Load(".env")

// 	// A signal handler or similar could be used to set this to false to break the loop.
// 	run := true

// 	for run {
// 		select {
// 		case sig := <-sigchan:
// 			fmt.Printf("Caught signal %v: terminating\n", sig)
// 			run = false
// 		default:
// 			ev := c.Poll(100)
// 			if ev == nil {
// 				continue
// 			}

// 			switch e := ev.(type) {
// 			case *kafka.Message:
// 				fmt.Println(*e.TopicPartition.Topic)
// 				if err != nil {
// 					fmt.Printf("Failed to deserialize payload: %s\n", err)
// 				} else {
// 					if err != nil {
// 						fmt.Printf("Failed to unmarshal payload: %s\n", err)
// 					}
// 					token := os.Getenv("SLACK_TOKEN")
// 					sendSlack(token, "event-alert", e.Value)
// 					fmt.Printf("%% Message on %s", e.TopicPartition)
// 				}
// 				if e.Headers != nil {
// 					fmt.Printf("%% Headers: %v\n", e.Headers)
// 				}
// 			case kafka.Error:
// 				// Errors should generally be considered
// 				// informational, the client will try to
// 				// automatically recover.
// 				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
// 			default:
// 				fmt.Printf("Ignored %v\n", e)
// 			}
// 		}
// 	}

// 	c.Close()
// }

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	token := os.Getenv("SLACK_TOKEN")
	sendSlack(token, "event-alert", "에러 메시지", []byte(`{"event_name":"view_home","user_id":"james","device_id":"b2b2b2b2-b2b2-b2b2-b2b2-b2b2b2b2b2b2","platform":"ios"}`))
}
