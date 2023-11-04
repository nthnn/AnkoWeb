package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func requestHandler(path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fileName := r.URL.Path[1:]

		if fileName == "" {
			if serverFileExists(path, "index.html") {
				fileName = "index.html"
			} else if serverFileExists(path, "index.htm") {
				fileName = "index.html"
			} else {
				fileName = "index.awp"
			}
		}

		if !serverFileExists(path, fileName) {
			handleNotFound(path, w, r)
			return
		}

		fileContents, contentType := getContentMimeType(path + "/" + fileName)
		w.Header().Set("Content-Type", contentType)
		w.Write(fileContents)
	}
}

func serverFileExists(path string, file string) bool {
	_, err := os.Stat(path + "/" + file)
	return err == nil
}

func handleNotFound(path string, w http.ResponseWriter, r *http.Request) {
	if serverFileExists(path, "404.awp") {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("404 Not Found"))
}

func getContentMimeType(fileName string) ([]byte, string) {
	fileContents, _ := ioutil.ReadFile(fileName)
	contentType := ""

	if strings.HasSuffix(fileName, ".awp") {
		contentType = "text/html"
	} else {
		contentType = http.DetectContentType(fileContents)
	}

	return fileContents, contentType
}

func awpServer(path string, host string, port int16) {
	http.HandleFunc("/", requestHandler(path))

	addr := host + ":" + strconv.Itoa(int(port))
	log.Println("Server is running on " + addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
