package server

import (
	"net/http"
	"os"
	"strings"
)

func serverFileExists(path string, file string) bool {
	_, err := os.Stat(path + "/" + file)
	return err == nil
}

func getContentMimeType(fileName string) ([]byte, string) {
	fileContents, _ := os.ReadFile(fileName)
	contentType := ""

	if strings.HasSuffix(fileName, ".awp") {
		contentType = "text/html"
	} else {
		contentType = http.DetectContentType(fileContents)
	}

	return fileContents, contentType
}
