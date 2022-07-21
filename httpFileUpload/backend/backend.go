// Based on https://tutorialedge.net/golang/go-file-upload-tutorial/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

var PORT = "8081"
var TEMP_DIR = "./tmp"
var UPLOAD_DIR = "./upload"
var wg sync.WaitGroup

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
	mReader(w, r)
}
func mReader(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		//if part.FileName() is empty, skip this iteration.
		if part.FileName() == "" {
			continue
		}
		dst, err := os.Create("./upload/" + part.FileName())
		defer dst.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

/*for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		log.Printf("Uploaded File: %+v\n", part.FileName())
		tempFile, err := ioutil.TempFile("tmp", "upload-*")
		if err != nil {
			log.Println(err)
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(part)
		if err != nil {
			log.Println(err)
		}
		tempFile.Write(fileBytes)
		err = os.Rename("./"+tempFile.Name(), UPLOAD_DIR+"/"+part.FileName())
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, "Upload successfully\n")
	}
}*/

/*
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
*/

func routes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":"+PORT, nil)
}

func main() {
	fmt.Println("File uploader backend started at port: " + PORT)
	routes()
}
