package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(home))

	log.Println("Service listening at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it has started")
}
