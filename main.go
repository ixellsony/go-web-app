package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/view/", serveView)

	http.ListenAndServe(":8080", nil)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func serveView(w http.ResponseWriter, r *http.Request) {
	viewName := r.URL.Path[len("/view/"):]

	if viewName == "" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("views/" + viewName + ".html")
	if err != nil {
		http.Error(w, "View not found", http.StatusNotFound)
		return
	}

	params := r.URL.Query()

	tmpl.Execute(w, params)
}
