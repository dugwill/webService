package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var tmpls, _ = template.ParseFiles(
	"templates/index.html",
	"templates/firstPage.html",
	"templates/secondPage.html")

func main() {

	fmt.Println("Hello World")

	server := http.Server{
		Addr: ":9000",
	}

	http.HandleFunc("/", Index)
	http.HandleFunc("/firstPage", firstPage)
	http.HandleFunc("/secondPage", secondPage)

	log.Fatalln(server.ListenAndServe())

}

// Index serves up the index files
func Index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title  string
		Header string
	}{
		Title:  "Index Page",
		Header: "Hello, World!",
	}

	if err := tmpls.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Index serves up the first page
func firstPage(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title  string
		Header string
	}{
		Title:  "first Page",
		Header: "This is the first page",
	}

	if err := tmpls.ExecuteTemplate(w, "firstPage.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.(http.Flusher).Flush()

	fmt.Printf("Going to sleep at %v\n", time.Now())

	time.Sleep(5 * time.Second)

	fmt.Printf("Waking up at %v\n", time.Now())

	data.Title="Second Page"
	data.Header="This is the second page"

	if err := tmpls.ExecuteTemplate(w, "secondPage.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return

}

// Index serves up the second page
func secondPage(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title  string
		Header string
	}{
		Title:  "Second Page",
		Header: "This is the second page",
	}

	if err := tmpls.ExecuteTemplate(w, "secondPage.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Going to sleep at %v\n", time.Now())

	w.(http.Flusher).Flush()

	time.Sleep(5 * time.Second)

	fmt.Printf("Waking up at %v\n", time.Now())

}
