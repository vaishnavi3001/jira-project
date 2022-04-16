package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	host := "0.0.0.0"
	port := "8500"
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/jira-frontend/index.html")
	})
	fmt.Printf("Server is running at %s:%s", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, r))
}
