package main

import (
	"io"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I am the root")
}
func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I am dog")
}
func me(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "I am me")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
