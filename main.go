package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/SirPeace/the-path/api"
)

func main() {
	http.HandleFunc("/example", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			api.GetExample(res, req)
		default:
			res.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	port, _ := strconv.ParseInt(os.Getenv("PORT"), 10, 64)
	if port == 0 {
		port = 8080
	}

	fmt.Printf("🚀 Listening on 0.0.0.0:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
