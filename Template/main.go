package main

import (
	"fmt"
	"net/http"
	"text/template"

	"studies.com/template/data"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/", renderIndexHtml)
	server.HandleFunc("/posts", renderPosts)

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

func renderPosts(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/posts.html.tmpl")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	temp.Execute(w, data.GetPosts())

	w.WriteHeader(http.StatusOK)
}
