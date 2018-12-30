package main

import (
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		_ = details

		tpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":9090", nil)
}
