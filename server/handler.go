package server

import (
	"net/http"
	"os"
	"strings"

	"github.com/nthnn/AnkoWeb/ankovm"
	"github.com/nthnn/AnkoWeb/logger"
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
			logger.Error(fileName + ": error")
			return
		}

		fileContents, contentType := getContentMimeType(path + "/" + fileName)
		w.Header().Set("Content-Type", contentType)

		if strings.HasSuffix(fileName, ".awp") {
			ankovm.RunScript(path, fileName, string(fileContents), w, r)
		} else {
			w.Write(fileContents)
		}

		logger.Info(fileName + ": ok")
	}
}

func handleNotFound(path string, w http.ResponseWriter, r *http.Request) {
	if serverFileExists(path, "404.awp") {
		fileContents, _ := os.ReadFile(path + "/" + "404.awp")
		ankovm.RunScript(path, "404.awp", string(fileContents), w, r)

		return
	}

	http.NotFound(w, r)
}
