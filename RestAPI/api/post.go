package api

import (
	"encoding/json"
	"net/http"

	"studies.com/template/data"
)

// Not fully working
func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var post data.BlogPost

		err := json.NewDecoder(r.Body).Decode(&post)

		if err != nil {
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
