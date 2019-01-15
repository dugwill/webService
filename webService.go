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

type Thing struct {
	Name string
	Type string
}

func main() {

	fmt.Println("Hello World")

	server := http.Server{
		Addr: "192.168.1.31:9000",
	}

	http.HandleFunc("/firstPage/", firstPage)
	http.HandleFunc("/secondPage/", secondPage)
	http.HandleFunc("/", Index)

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

	fmt.Println("Really, what am I doing here")

	if err := tmpls.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// firstPage serves up the first page
func firstPage(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title  string
		Header string
		List   []Thing
	}{
		Title:  "first Page",
		Header: "This is the first page",
	}

	data.List = append(data.List, Thing{Name: "Thing 1", Type: "widget"})
	data.List = append(data.List, Thing{Name: "Thing 2", Type: "not a widget"})

	fmt.Printf("Length of data list: %d\n", len(data.List))

	for _, i := range data.List {
		fmt.Println(i)
	}

	if err := tmpls.ExecuteTemplate(w, "firstPage.html", data); err != nil {
		fmt.Printf("failed to server Index")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//w.(http.Flusher).Flush()

	/*
		fmt.Printf("Going to sleep at %v\n", time.Now())

		time.Sleep(5 * time.Second)

		fmt.Printf("Waking up at %v\n", time.Now())

		data.Title = "Second Page"
		data.Header = "This is the second page"

		if err := tmpls.ExecuteTemplate(w, "secondPage.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	*/
	return

}

// secondPage serves up the second page
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

	fmt.Printf("hoping I am not here")

	fmt.Printf("Going to sleep at %v\n", time.Now())

	w.(http.Flusher).Flush()

	time.Sleep(5 * time.Second)

	fmt.Printf("Waking up at %v\n", time.Now())

}
