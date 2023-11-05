package ankovm

import (
	"net/http"

	"github.com/mattn/anko/env"
	"github.com/nthnn/AnkoWeb/logger"
)

func installDefinition(vmEnv *env.Env, symbol string, value interface{}) {
	if err := vmEnv.Define(symbol, value); err != nil {
		logger.Error("Error: " + err.Error())
	}
}

func installDefaults(vmEnv *env.Env, path string, w http.ResponseWriter, r *http.Request) {
	installDefinition(vmEnv, "echo", echoFn(w))
	installDefinition(vmEnv, "include", includeFn(vmEnv, path))

	installDefinition(vmEnv, "httpHeaders", httpHeaderFn(r))
	installDefinition(vmEnv, "httpRemote", httpRemoteFn(r))
}
