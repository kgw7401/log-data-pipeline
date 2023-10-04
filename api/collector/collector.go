package collector

import (
	"io"
	"net/http"
)

func collectHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	print(string(body))
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", collectHandler)

	return mux
}
