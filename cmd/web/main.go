package main

import (
	"log"
	"net/http"
)

func main() {
	mux := routes()

	err := http.ListenAndServe(":7007", mux)

	if err != nil {
		log.Fatal(err)
	}

}
