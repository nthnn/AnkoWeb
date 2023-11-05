package ankovm

import (
	"net/http"

	"github.com/mattn/anko/env"
	"github.com/nthnn/AnkoWeb/ankovm/std"
	"github.com/nthnn/AnkoWeb/logger"
)

func installDefinition(vmEnv *env.Env, symbol string, value interface{}) {
	if err := vmEnv.Define(symbol, value); err != nil {
		logger.Error("Error: " + err.Error())
	}
}

func installWebFunc(vmEnv *env.Env, path string, w http.ResponseWriter, r *http.Request) {
	installDefinition(vmEnv, "echo", std.EchoFn(w))
	installDefinition(vmEnv, "include", std.IncludeFn(vmEnv, path))

	installDefinition(vmEnv, "httpHeaders", std.HttpHeaderFn(r))
	installDefinition(vmEnv, "httpRemote", std.HttpRemoteFn(r))
}

func installJsonFunc(vmEnv *env.Env, path string, w http.ResponseWriter, r *http.Request) {
	installDefinition(vmEnv, "mapToJson", std.MapToJsonFn)
	installDefinition(vmEnv, "jsonToMap", std.JsonToMapFn)
	installDefinition(vmEnv, "jsonPrettify", std.JsonPrettifyFn)
	installDefinition(vmEnv, "lastJsonError", std.LastJsonErrorFn)
}

func installDefaults(vmEnv *env.Env, path string, w http.ResponseWriter, r *http.Request) {
	installWebFunc(vmEnv, path, w, r)
	installJsonFunc(vmEnv, path, w, r)
}
