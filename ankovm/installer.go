package ankovm

import (
	"bytes"
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

func installWebFunc(vmEnv *env.Env, path string, buff *bytes.Buffer, w http.ResponseWriter, r *http.Request) {
	installDefinition(vmEnv, "echo", std.EchoFn(buff))
	installDefinition(vmEnv, "include", std.IncludeFn(vmEnv, path))

	installDefinition(vmEnv, "httpHeaders", std.HttpHeaderFn(r))
	installDefinition(vmEnv, "httpRemote", std.HttpRemoteFn(r))

	installDefinition(vmEnv, "setCookie", std.SetCookieFn(w))
	installDefinition(vmEnv, "getCookie", std.GetCookieFn(r))
	installDefinition(vmEnv, "deleteCookie", std.DeleteCookieFn(w))
}

func installJsonFunc(vmEnv *env.Env, path string, buff *bytes.Buffer, w http.ResponseWriter, r *http.Request) {
	installDefinition(vmEnv, "mapToJson", std.MapToJsonFn)
	installDefinition(vmEnv, "jsonToMap", std.JsonToMapFn)
	installDefinition(vmEnv, "jsonPrettify", std.JsonPrettifyFn)
	installDefinition(vmEnv, "lastJsonError", std.LastJsonErrorFn)
}

func installTimeFunc(vmEnv *env.Env, path string, buff *bytes.Buffer, w http.ResponseWriter, r *http.Request) {
	installDefinition(vmEnv, "unixTimeNow", std.UnixTimeNowFn)
}

func installDefaults(vmEnv *env.Env, path string, buff *bytes.Buffer, w http.ResponseWriter, r *http.Request) {
	installWebFunc(vmEnv, path, buff, w, r)
	installJsonFunc(vmEnv, path, buff, w, r)
	installTimeFunc(vmEnv, path, buff, w, r)
}
