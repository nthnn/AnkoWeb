package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

func parseAwpFile(fileName string, fileContent string) (string, error) {
	insideAnko := false
	hasPrevLt := false
	hasPrevPr := false

	parsed := "echo(\""
	idx := 0

	for idx < len(fileContent) {
		current := rune(fileContent[idx])

		if hasPrevLt && current == '%' {
			if insideAnko {
				return "", errors.New("already inside Anko context")
			}

			hasPrevLt = false
			insideAnko = true

			parsed += "\")"
			idx += 1
			continue
		} else if hasPrevPr && current == '>' {
			if !insideAnko {
				return "", errors.New("enclosing Anko outside context")
			}

			hasPrevPr = false
			insideAnko = false

			parsed += "echo(\""
			idx += 1
			continue
		} else if (hasPrevLt && current != '%') ||
			(hasPrevPr && current != '>') {
			if hasPrevLt {
				parsed += "<"
			} else if hasPrevPr {
				parsed += "%"
			}

			hasPrevLt = false
			hasPrevPr = false
		} else {
			if current == '<' {
				hasPrevLt = true

				idx += 1
				continue
			} else if current == '%' {
				hasPrevPr = true

				idx += 1
				continue
			}
		}

		if current == '"' && !insideAnko {
			parsed += "\""
		} else {
			if current == '\n' && !insideAnko {
				parsed += "\\n"
			} else {
				parsed += string(current)
			}
		}

		idx += 1
	}

	if insideAnko {
		return "", errors.New("unenclosed Anko context, encountered end-of-file")
	}

	parsed += "\")"
	return parsed, nil
}

func echoFn(w http.ResponseWriter) func(message string) {
	return func(message string) {
		w.Write([]byte(message))
	}
}

func runScript(fileName string, fileContents string, w http.ResponseWriter) {
	parsed, err := parseAwpFile("", fileContents)
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}

	vmEnv := env.NewEnv()
	if err = vmEnv.Define("echo", echoFn(w)); err != nil {
		log.Fatal("Error: ", err)
		return
	}

	_, err = vm.Execute(vmEnv, nil, parsed)
	if err != nil {
		log.Fatal("Execution error: ", err)
		return
	}
}
