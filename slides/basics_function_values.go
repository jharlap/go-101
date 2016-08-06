package main

import (
	"fmt"
	"net/http"
)

func main() {
	addr := ":80"

	// handler is a variable of type func(ResponseWriter, *Request)
	handler := func(w http.ResponseWriter, r *http.Request) {

		// note the func closed over addr
		fmt.Fprintf("Hello! This server listens on %s", addr)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}
