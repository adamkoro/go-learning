// Based on https://tutorialedge.net/golang/go-file-upload-tutorial/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var PORT = "8080"
var TEMP_DIR = "./tmp"
var UPLOAD_DIR = "./upload"

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading file\n")
	err := os.Mkdir(TEMP_DIR, 0755)
	if err != nil {
		log.Println(err)
	}

	os.Mkdir(UPLOAD_DIR, 0755)
	if err != nil {
		log.Println(err)
	}

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println("Error file from from-data")
		log.Println(err)
	}
	defer file.Close()

	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("tmp", "upload-*")
	if err != nil {
		log.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	tempFile.Write(fileBytes)
	err = os.Rename("./"+tempFile.Name(), UPLOAD_DIR+"/"+handler.Filename)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "Upload successfully\n")
}

func routes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":"+PORT, nil)
}

func main() {
	fmt.Println("File uploader")
	routes()
}
