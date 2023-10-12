package collector

import (
	"encoding/json"
	"io"
	"net/http"

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

	serializedData, err := kafkaProducer.Serialize(eventName, result)
	if err != nil {
		kafkaProducer.Produce("dlt.event", data)
	} else {
		kafkaProducer.Produce(eventName, serializedData)
	}
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", collectHandler)

	return mux
}
