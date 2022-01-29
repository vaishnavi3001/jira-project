package main
import (
	"net/http"
	"embed"
	"io/fs"
	"log"
	"fmt"
)


//go:embed assets/*
var static embed.FS


func main() {
	host :=  "0.0.0.0"
	port := "8080"
	contentStatic, _ := fs.Sub(static, "assets/jira-frontend")
	http.Handle("/", http.FileServer(http.FS(contentStatic)))
	log.Fatal(http.ListenAndServe(host+":"+port,nil))
	fmt.Printf("Server is running at %s:%s",host, port)
}