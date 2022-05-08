package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	tmpl, err := template.ParseFiles("page.html")
	if err != nil {
		log.Panicln("invalid template")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, r)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
