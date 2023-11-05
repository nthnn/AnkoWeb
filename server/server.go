package server

import (
	"net/http"
	"strconv"

	"github.com/nthnn/AnkoWeb/logger"
)

func AwpServer(path string, host string, port int16) {
	http.HandleFunc("/", requestHandler(path))

	addr := host + ":" + strconv.Itoa(int(port))
	logger.Plain("Server is running on " + addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		logger.Error("Error: " + err.Error())
	}
}
