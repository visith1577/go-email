package main

import (
	"html/template"
	"log"
	"net/http"

	pat "github.com/bmizerany/pat"
)

func home(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/home.html", nil)
}

func confirmation(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/confirmation.html", nil)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	templ, err := template.ParseFiles(filename)
	if err != nil {
		log.Print(err)
		http.Error(w, "Something went wrong!!!", http.StatusInternalServerError)
	}

	if err := templ.Execute(w, data); err != nil {
		log.Print(err)
		http.Error(w, "Something went wrong!!!", http.StatusInternalServerError)
	}
}

func send(w http.ResponseWriter, r *http.Request) {
	msg := &Message{
		Email:   r.PostFormValue("email"),
		content: r.PostFormValue("content"),
	}

	if msg.Validate() == false {
		render(w, "templates/home.html", msg)
		return
	}

	if err := msg.Deliver(); err != nil {
		log.Print(err)
		http.Error(w, "Something went wrong!!!", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
}

func main() {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(home))
	mux.Post("/", http.HandlerFunc(send))
	mux.Get("/confirmations", http.HandlerFunc(confirmation))

	log.Print("Listening...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
