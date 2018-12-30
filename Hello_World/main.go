package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World you've requested %s \n", r.URL.Path)
	})

	fmt.Println("run server on 9090")

	http.ListenAndServe(":9090", nil)
}
