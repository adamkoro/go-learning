package main

import (
	"fmt"
	"net/http"
)

var PORT = "8080"
var WEBPAGE_CONTENT = "./static"

func routes() {
	http.Handle("/", http.FileServer(http.Dir(WEBPAGE_CONTENT)))
	http.ListenAndServe(":"+PORT, nil)
}

func main() {
	fmt.Println("File uploader frontend started at port: " + PORT)
	routes()
}
