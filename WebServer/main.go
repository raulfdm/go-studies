package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.NewServeMux()

	staticServer := http.FileServer(http.Dir("./wwwroot"))

	server.Handle("/", staticServer)

	err := http.ListenAndServe(":80", server)

	if err != nil {
		fmt.Println("Something went wrong while trying to boot the server")
	}

}
