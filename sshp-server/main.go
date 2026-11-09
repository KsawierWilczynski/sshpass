package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)


func main() {
	registerEndpoints()

	log.Fatal(http.ListenAndServe(":1337", nil))
}

func registerEndpoints() {
	http.HandleFunc("GET /auth", helloWorld)
	http.HandleFunc("GET /refresh", helloWorld)
	http.HandleFunc("POST /execute/{user}/{session}", executeCommand)
	http.HandleFunc("POST /createSession/{user}", helloWorld)
	http.HandleFunc("GET /history/{user}/{session}", helloWorld)
	http.HandleFunc("POST /stopSession/{user}/{session}", helloWorld)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

type Command struct {
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

func executeCommand(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var command Command
	err := decoder.Decode(&command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := exec.Command(command.Cmd, command.Args...).CombinedOutput()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", string(output))
}
