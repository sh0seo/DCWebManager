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

	// favicon
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/icon/favicon.ico")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%7s] /", r.Method)

		tmplIndex := "templates/index.html"
		tmplHeader := "templates/include/header.html"
		tmplFooter := "templates/include/footer.html"
		// tmplPath2 := "templates/login.html"

		tmpl := template.Must(template.ParseFiles(tmplIndex, tmplHeader, tmplFooter))
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

		// fmt.Fprintf(w, `{"result":"fail","code":1000}`)
		fmt.Fprintf(w, `{"result":"success","code":1000}`)
	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%7s] /logout", r.Method)
		fmt.Fprintf(w, `{result":"success","code":200}`)
	})

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
