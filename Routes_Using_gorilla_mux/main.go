package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		title := params["title"]
		page := params["page"]
		fmt.Fprintf(w, "You've requested the book : %s on page %s\n", title, page)
	})

	fmt.Println("[(*)] run server on 9090 port.")

	http.ListenAndServe(":9090", r)
}
