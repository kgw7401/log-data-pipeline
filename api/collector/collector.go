package collector

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/log-data-pipeline/api/kafka"
)

func collectHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	q := r.URL.Query()
	eventName := q.Get("event_name")

	result := make(map[string]interface{})
	_ = json.Unmarshal(data, &result)

	kafka.ProduceData(eventName, result)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", collectHandler)

	return mux
}
