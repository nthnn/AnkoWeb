package ankovm

import (
	"context"
	"net/http"
	"os"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
	"github.com/nthnn/AnkoWeb/logger"
)

func echoFn(w http.ResponseWriter) func(message string) {
	return func(message string) {
		w.Write([]byte(message))
	}
}

func includeFn(ctx *context.Context, vmEnv *env.Env, path string) func(fileName string) {
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

		_, err = vm.ExecuteContext(*ctx, vmEnv, nil, parsed)
		if err != nil {
			logger.Error("Execution error: " + err.Error())
			return
		}
	}
}
