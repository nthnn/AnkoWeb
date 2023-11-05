package ankovm

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"github.com/nthnn/AnkoWeb/logger"
)

func echoFn(w http.ResponseWriter) func(value interface{}) {
	return func(value interface{}) {
		str := fmt.Sprintf("%v", value)
		w.Write([]byte(str))
	}
}

func includeFn(vmEnv *env.Env, path string) func(fileName string) {
	return func(fileName string) {
		fileContents, err := os.ReadFile(path + "/" + fileName)
		if err != nil {
			logger.Error("Error: " + err.Error())
			return
		}

		parsed, err := parseAwpFile("", string(fileContents))
		if err != nil {
			logger.Error("Error: " + err.Error())
			return
		}

		_, err = parser.ParseSrc(parsed)
		scanner := parser.Scanner{}
		scanner.Scan()

		if err != nil {
			logger.Error("Parser error: " + err.Error() + " (" + fileName + ")")
			return
		}

		_, err = vm.Execute(vmEnv, nil, parsed)
		if err != nil {
			logger.Error("Execution error: " + err.Error() + " (" + fileName + ")")
			return
		}
	}
}

func httpHeaderFn(r *http.Request) func() map[string][]string {
	return func() map[string][]string {
		return r.Header
	}
}

func httpRemoteFn(r *http.Request) func() map[string]interface{} {
	return func() map[string]interface{} {
		dict := make(map[string]interface{})

		dict["method"] = r.Method
		dict["host"] = r.Host
		dict["length"] = r.ContentLength
		dict["remote_addr"] = r.RemoteAddr
		dict["request_uri"] = r.RequestURI
		dict["encoding"] = r.TransferEncoding
		dict["url"] = r.URL.String()
		dict["form"] = r.Form
		dict["post_data"] = r.PostForm
		dict["cookies"] = r.Cookies()

		return dict
	}
}
