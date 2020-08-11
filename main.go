// main.go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// NAME is App name
const NAME = "DCCaffe Web Manager Start"

// PORT is Server port
const PORT = ":1323"

// Params is data
type Params struct {
	PageName string
	UserName string
}

func main() {
	log.Printf("%s[%s]", NAME, PORT)

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// favicon
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/icon/favicon.ico")
	})

	http.HandleFunc("/user_regist", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%7s] /", r.Method)
		w.Header().Set("Content-Type", "text/html")

		tmplName := "templates/user_manage.html"
		tmpl := getTemplate(tmplName)

		p := Params{
			PageName: "user_regist",
			UserName: "Tester",
		}

		if err := tmpl.ExecuteTemplate(w, "layout", p); err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Print(err)
		}
	})

	http.HandleFunc("/user_manage", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%7s] /", r.Method)
		w.Header().Set("Content-Type", "text/html")

		tmplName := "templates/user_manage.html"
		tmpl := getTemplate(tmplName)

		p := Params{
			PageName: "user_manage",
			UserName: "Tester",
		}

		if err := tmpl.ExecuteTemplate(w, "layout", p); err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Print(err)
		}
	})

	http.HandleFunc("/menu_manage", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%7s] /", r.Method)
		w.Header().Set("Content-Type", "text/html")

		tmplName := "templates/menu_manage.html"
		tmpl := getTemplate(tmplName)

		p := Params{
			PageName: "user_manage",
			UserName: "Tester",
		}

		if err := tmpl.ExecuteTemplate(w, "layout", p); err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Print(err)
		}
	})

	http.HandleFunc("/cancel_order_manage", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%7s] /", r.Method)
		w.Header().Set("Content-Type", "text/html")

		tmplName := "templates/cancel_order_manage.html"
		tmpl := getTemplate(tmplName)

		p := Params{
			PageName: "cancel_order_manage",
			UserName: "Tester",
		}

		if err := tmpl.ExecuteTemplate(w, "layout", p); err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Print(err)
		}
	})

	http.HandleFunc("/bill", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%7s] /", r.Method)
		w.Header().Set("Content-Type", "text/html")

		tmplName := "templates/bill.html"
		tmpl := getTemplate(tmplName)

		p := Params{
			PageName: "bill",
			UserName: "Tester",
		}

		if err := tmpl.ExecuteTemplate(w, "layout", p); err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Print(err)
		}
	})

	http.HandleFunc("/customer_bill", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%7s] /", r.Method)
		w.Header().Set("Content-Type", "text/html")

		tmplName := "templates/customer_bill.html"
		tmpl := getTemplate(tmplName)

		p := Params{
			PageName: "customer_bill",
			UserName: "Tester",
		}

		if err := tmpl.ExecuteTemplate(w, "layout", p); err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Print(err)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%7s] /", r.Method)
		w.Header().Set("Content-Type", "text/html")

		tmplIndex := "templates/index.html"
		tmpl := getTemplate(tmplIndex)

		p := Params{
			PageName: "index",
			UserName: "Teseter",
		}
		err := tmpl.ExecuteTemplate(w, "layout", p)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Print(err)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%7s] /login", r.Method)
		w.Header().Set("Content-Type", "text/html")

		if r.Method != "POST" {
			http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// fmt.Fprintf(w, `{"result":"fail","code":1000}`)
		fmt.Fprintf(w, `{"result":"success","code":1000}`)
	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%7s] /logout", r.Method)
		w.Header().Set("Content-Type", "text/html")

		fmt.Fprintf(w, `{result":"success","code":200}`)
	})

	err := http.ListenAndServe(PORT, nil)
	log.Fatal(err)
}

func getTemplate(name string) *template.Template {
	tmplLayout := "templates/include/layout.html"
	tmplHeader := "templates/include/header.html"
	tmplFooter := "templates/include/footer.html"
	tmplSide := "templates/include/side.html"

	return template.Must(template.ParseFiles(name, tmplLayout, tmplHeader, tmplFooter, tmplSide))
}
