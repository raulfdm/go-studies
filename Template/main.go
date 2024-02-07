package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/", renderIndexHtml)

	err := http.ListenAndServe(":80", server)

	if err != nil {
		fmt.Println("Something went wrong while trying to boot the server")
	}
}

func renderIndexHtml(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/index.html.tmpl")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	temp.Execute(w, "Hello World!")

	w.WriteHeader(http.StatusOK)
}
