package main

import (
	"fmt"
	handler "go-api/api"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Handler)

	err := http.ListenAndServe(":6969", mux)

	if err != nil {
		fmt.Println(err)
	}
}
