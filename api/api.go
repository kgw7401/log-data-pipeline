package main

import (
	"net/http"

	"github.com/log-data-pipeline/api/collector"
)

func main() {
	http.ListenAndServe(":3000", collector.NewHttpHandler())
}
