package main

import (
	"html/template"
	"net/http"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/works", works)
	http.HandleFunc("/choral", choral)
	http.HandleFunc("/chamber", chamber)
	http.HandleFunc("/reviews", reviews)
	http.HandleFunc("/contact", contact)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/pub/", http.StripPrefix("/pub", http.FileServer(http.Dir("./pub"))))
	http.ListenAndServe(":8080", nil)
}

func index (w http.ResponseWriter, _ *http.Request) {

	err := tpl.ExecuteTemplate(w, "index.html", index)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal serverrrrrr error", http.StatusInternalServerError)
		return
	}
}

func about (w http.ResponseWriter, _ *http.Request) {

	err := tpl.ExecuteTemplate(w, "bio.html", about)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal serverrrrrr error", http.StatusInternalServerError)
		return
	}

}

func works (w http.ResponseWriter, _ *http.Request) {

	err := tpl.ExecuteTemplate(w, "works.html", works)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal serverrrrrr error", http.StatusInternalServerError)
		return
	}

}

func choral (w http.ResponseWriter, _ *http.Request) {

	err := tpl.ExecuteTemplate(w, "choral-works.html", choral)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal serverrrrrr error", http.StatusInternalServerError)
		return
	}

}

func chamber (w http.ResponseWriter, _ *http.Request) {

	err := tpl.ExecuteTemplate(w, "chamber-works.html", chamber)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal serverrrrrr error", http.StatusInternalServerError)
		return
	}

}

func reviews (w http.ResponseWriter, _ *http.Request) {

	err := tpl.ExecuteTemplate(w, "reviews.html", reviews)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal serverrrrrr error", http.StatusInternalServerError)
		return
	}

}

func contact (w http.ResponseWriter, _ *http.Request) {

	err := tpl.ExecuteTemplate(w, "contact.html", contact)
	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal serverrrrrr error", http.StatusInternalServerError)
		return
	}
}

