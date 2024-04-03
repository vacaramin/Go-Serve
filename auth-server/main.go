package main

import (
	"net/http"
)

func main() {
	// Start the server
	http.Handle("/", getRouter())
}
func getRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/login", loginHandler)
	return router
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the login request
}
