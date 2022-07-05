package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

var PORT = ":8000"

func currentTrack(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	cmd := exec.Command("mpc status | head -n 2")
	err := cmd.Run()
	stdout := []byte("Current track not found")
	if err == nil {
		stdout, err = cmd.Output()
	}
	fmt.Fprintf(w, string(stdout))
}

func skipTrack(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("mpc next")
	err := cmd.Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("Unable to skip track")
	}
}

func handleRequests() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/current", currentTrack)
	http.HandleFunc("/skip", skipTrack)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func main() {
	handleRequests()
}
