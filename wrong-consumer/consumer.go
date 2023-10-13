package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/slack-go/slack"
)

func slackChatPost(message []byte) {
	api := slack.New("xoxb-6029376210659-6045646293281-5v763nuI0gTsRJOCrxweGbx2")
	attachment := slack.Attachment{
		Pretext: string(message),
	}
	channelID, timestamp, err := api.PostMessage(
		"event-alert",
		slack.MsgOptionText("", false),
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "eventConsumerGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	if err != nil {
		fmt.Printf("Failed to create schema registry client: %s\n", err)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Failed to create deserializer: %s\n", err)
		os.Exit(1)
	}

	c.SubscribeTopics([]string{"dlt.event"}, nil)

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Println(*e.TopicPartition.Topic)
				if err != nil {
					fmt.Printf("Failed to deserialize payload: %s\n", err)
				} else {
					if err != nil {
						fmt.Printf("Failed to unmarshal payload: %s\n", err)
					}
					slackChatPost(e.Value)
					fmt.Printf("%% Message on %s:\n%+v\n", e.TopicPartition)
				}
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	c.Close()
}
