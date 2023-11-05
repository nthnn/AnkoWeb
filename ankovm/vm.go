package ankovm

import (
	"net/http"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"github.com/nthnn/AnkoWeb/logger"
	"github.com/nthnn/AnkoWeb/util"
)

func RunScript(path string, fileName string, fileContents string, w http.ResponseWriter, r *http.Request) {
	parsed, err := util.ParseAwpFile(fileName, fileContents)
	if err != nil {
		logger.Error("Error: " + err.Error())
		return
	}

	vmEnv := env.NewEnv()
	installDefaults(vmEnv, path, w, r)

	_, err = parser.ParseSrc(parsed)
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
