package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Hello World"}`))
	})

	err := http.ListenAndServe(":3000", server)

	if err != nil {
		fmt.Println("Something went wrong while trying to boot the server")
	}

}
