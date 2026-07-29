package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "🚀 Go Docker App Running")
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Healthy")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/health", health)

	fmt.Println("Listening on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
