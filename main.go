package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// NAME is App name
const NAME = "\n\n\n\nDCCaffe Web Manager Start\n\n\n\n"

func main() {
	log.Printf(NAME)

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%7s] /", r.Method)

		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		time := struct {
			Title string
		}{
			Title: fmt.Sprintf("%s %s", NAME, time.Now()),
		}
		tmpl.Execute(w, time)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%7s] /login", r.Method)
		if r.Method != "POST" {
			http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		fmt.Fprintf(w, `{"result":"fail","code":1000}`)
	})

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
