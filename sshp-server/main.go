package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	registerEndpoints()

	log.Fatal(http.ListenAndServe(":1337", nil))
}

func registerEndpoints() {
	http.HandleFunc("GET /auth", helloWorld)
	http.HandleFunc("GET /refresh", helloWorld)
	http.HandleFunc("POST /execute/{user}/{session}", helloWorld)
	http.HandleFunc("POST /createSession/{user}", helloWorld)
	http.HandleFunc("GET /history/{user}/{session}", helloWorld)
	http.HandleFunc("POST /stopSession/{user}/{session}", helloWorld)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
