package ankovm

import (
	"context"
	"net/http"

	"github.com/mattn/anko/env"
	"github.com/nthnn/AnkoWeb/logger"
)

func installDefaults(ctx *context.Context, vmEnv *env.Env, path string, w http.ResponseWriter, r *http.Request) {
	if err := vmEnv.Define("echo", echoFn(w)); err != nil {
		logger.Error("Error: " + err.Error())
		return
	}

	if err := vmEnv.Define("include", includeFn(ctx, vmEnv, path)); err != nil {
		logger.Error("Error: " + err.Error())
		return
	}

	if err := vmEnv.Define("httpHeaders", httpHeaderFn(r)); err != nil {
		logger.Error("Error: " + err.Error())
		return
	}
}
