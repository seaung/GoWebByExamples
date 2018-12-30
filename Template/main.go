package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTtitle string
	Todos      []Todo
}

func main() {
	tpl := template.Must(template.ParseFiles("./static/layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTtitle: "Todo List",
			Todos: []Todo{
				{Title: "go language", Done: false},
				{Title: "pythonic", Done: true},
				{Title: "TypeScript", Done: false},
				{Title: "JavaScript", Done: true},
			},
		}
		tpl.Execute(w, data)
	})

	fmt.Println(" * run server on 9090 port.")

	http.ListenAndServe(":9090", nil)
}
