package main

import (
	"fmt"
	"log"
	"net/http"
)

func conct(a, b string) string {
	return a + b
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful\n")

	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", email)
	// if r.URL.Path != "/form" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }

	// if r.Method != "POST" {
	// 	http.Error(w, "Method is not supported.", http.StatusNotFound)
	// 	return
	// } else {

	// }
}
func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	} else {
		fmt.Fprintf(w, "Hello, world!")
	}

}
func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)

	fmt.Printf("Server is running on port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
