package collector

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/log-data-pipeline/api/kafka"
)

func collectHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	result := make(map[string]string)
	_ = json.Unmarshal(data, &result)
	eventName := result["event_name"]
	kafka.ProduceData(eventName, data)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", collectHandler)

	return mux
}
