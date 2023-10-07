package collector

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/xeipuuv/gojsonschema"
)

func collectHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	unmarshalData := data
	var resultMap map[string]interface{}
	json.Unmarshal(unmarshalData, &resultMap)
	jsonFile := fmt.Sprintf("file:///Users/kgw7401/log-data-pipeline/api/collector/schema/%s.json", resultMap["event_name"])

	schemaLoader := gojsonschema.NewReferenceLoader(jsonFile)
	documentLoader := gojsonschema.NewBytesLoader(data)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}
	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, err := range result.Errors() {
			fmt.Printf("- %s\n", err)
		}
	}
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", collectHandler)

	return mux
}
