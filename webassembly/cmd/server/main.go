package main

import (
	"log"
	"net/http"
)

var PORT = "8080"
var FILE_DIR = "../../assets"

func main() {
	log.Println("Server started at port: " + PORT)
	http.ListenAndServe(":"+PORT, http.FileServer(http.Dir(FILE_DIR)))
}
