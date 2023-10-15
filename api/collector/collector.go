package collector

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/log-data-pipeline/api/kafka"
)

var kafkaProducer *kafka.KafkaProducer

func init() {
	var err error
	kafkaProducer, err = kafka.NewKafkaProducer()
	if err != nil {
		panic(err)
	}
}

func collectHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	q := r.URL.Query()
	eventName := q.Get("event_name")

	result := make(map[string]interface{})
	_ = json.Unmarshal(data, &result)

	serializedData, serErr := kafkaProducer.Serialize(eventName, result)

	if serErr != nil {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
		token := os.Getenv("SLACK_TOKEN")
		kafka.SendSlack(token, "event-alert", serErr.Error(), data)
	} else {
		kafkaProducer.Produce(eventName, serializedData)
	}
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", collectHandler)

	return mux
}
