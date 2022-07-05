package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
    "strconv"
    "strings"
)

var PORT = ":8080"

func connListeners(w http.ResponseWriter, r *http.Request) {
    cmd := "/usr/bin/ss | grep 8000 | wc -l"
	out := exec.Command("bash", "-c", cmd)
    stdout, err := out.Output()
    if err != nil  {
        stdout = []byte("Current track not found")
    }
    x := strings.ReplaceAll(string(stdout), "\n", "")
    conn, err := strconv.Atoi(x)
    if err != nil {
        fmt.Println(err)
    }
    conn = conn / 3
    fmt.Fprintf(w, "Connected listeners: %d", conn)
}

func currentTrack(w http.ResponseWriter, r *http.Request) {
    cmd := "/usr/bin/mpc status | head -n 1"
	out := exec.Command("bash", "-c", cmd)
    stdout, err := out.Output()
    if err != nil  {
        stdout = []byte("Current track not found")
    }
	fmt.Fprintf(w, strings.ReplaceAll(string(stdout), "\n", "</br>"))
}

func skipTrack(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("/usr/bin/mpc", "next")
	err := cmd.Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print("Unable to skip track")
	}
}

func handleRequests() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/listeners", connListeners)
	http.HandleFunc("/current", currentTrack)
	http.HandleFunc("/skip", skipTrack)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func main() {
	handleRequests()
}
