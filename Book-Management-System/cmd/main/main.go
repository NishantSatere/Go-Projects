package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/nishantdev/book-management-system/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Passing the router to the RegisterBookStoreRoutes function
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
