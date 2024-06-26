package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const PORT string = ":4000"

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/404", NotFoundHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/pong", pongHandler)

	log.Println("Server started listening on port ", PORT)
	http.ListenAndServe(PORT, nil)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"hello": "world"}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}
func pingHandler(w http.ResponseWriter, _ *http.Request) {
	response := map[string]string{
		"Message": "Pong",
	}
	jsonResponse, err := json.MarshalIndent(response, "", "2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func pongHandler(w http.ResponseWriter, _ *http.Request) {
	response := map[string]string{
		"Message": "Ping",
	}
	jsonResponse, err := json.MarshalIndent(response, "", "2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"error": "404- Not found"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to marshal json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(jsonResponse)
}
