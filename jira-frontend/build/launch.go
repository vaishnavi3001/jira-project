package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed assets/*
var static embed.FS

func main() {
	host := "0.0.0.0"
	port := "5000"
	contentStatic, _ := fs.Sub(static, "assets/jira-frontend")
	http.Handle("/", http.FileServer(http.FS(contentStatic)))
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
	fmt.Printf("Server is running at %s:%s", host, port)
}
