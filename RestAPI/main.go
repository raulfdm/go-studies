package main

import (
	"fmt"
	"net/http"

	"studies.com/rest-api/api"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/api/posts", api.Get)

	err := http.ListenAndServe(":3000", server)

	if err != nil {
		fmt.Println("Something went wrong while trying to boot the server")
	}

}
