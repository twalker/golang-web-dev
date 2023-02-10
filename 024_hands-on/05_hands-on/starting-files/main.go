package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	// fs := http.FileServer(http.Dir("public"))
	// http.Handle("/pics/", fs)
	// http.Handle("/pics/", http.StripPrefix("public", http.FileServer(http.Dir("public/"))))
	http.Handle("/pics/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", root)

	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, nil)

}
