package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darkphotonKN/go-im/internal/handlers"
)

func main() {
	mux := routes()

	// start goroutine for websocket channel listener
	fmt.Println("Started listening to websocket channel.")
	go handlers.ListenForWSChannel()

	err := http.ListenAndServe(":7007", mux)

	if err != nil {
		log.Fatal(err)
	}
}
