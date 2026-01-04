package api

import (
	"encoding/json"
	"net/http"
)

func GetExample(res http.ResponseWriter, req *http.Request) {
	responseData := map[string]string{"test": "hello world!"}
	jsonBytes, _ := json.Marshal(responseData)

	res.Header().Add("Content-Type", "application/json")
	res.Write(jsonBytes)
}
