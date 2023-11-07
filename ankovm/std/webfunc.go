package std

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"github.com/nthnn/AnkoWeb/logger"
	"github.com/nthnn/AnkoWeb/util"
)

func EchoFn(buff *bytes.Buffer) func(value interface{}) {
	return func(value interface{}) {
		str := fmt.Sprintf("%v", value)
		buff.WriteString(str)
	}
}

func IncludeFn(vmEnv *env.Env, path string) func(fileName string) {
	return func(fileName string) {
		fileContents, err := os.ReadFile(path + "/" + fileName)
		if err != nil {
			logger.Error("Error: " + err.Error())
			return
		}

		parsed, err := util.ParseAwpFile("", string(fileContents))
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

func HeaderFn(w http.ResponseWriter) func(name, value string) {
	return func(name, value string) {
		w.Header().Add(name, value)
	}
}

func HttpHeaderFn(r *http.Request) func() map[string][]string {
	return func() map[string][]string {
		return r.Header
	}
}

func HttpRemoteFn(r *http.Request) func() map[string]interface{} {
	return func() map[string]interface{} {
		dict := make(map[string]interface{})
		body, err := io.ReadAll(r.Body)

		dict["method"] = r.Method
		dict["host"] = r.Host
		dict["length"] = r.ContentLength
		dict["remote_addr"] = r.RemoteAddr
		dict["request_uri"] = r.RequestURI
		dict["encoding"] = r.TransferEncoding
		dict["url"] = r.URL.String()
		dict["form"] = r.Form
		dict["cookies"] = r.Cookies()

		if err == nil {
			body := string(body)
			dict["body"] = body
		}

		if r.Method == "POST" && dict["body"] != nil {
			entries := strings.Split(string(body), "&")
			post := make(map[string]string)

			for i := 0; i < len(entries); i++ {
				pair := strings.Split(entries[i], "=")
				post[util.UnescapeString(pair[0])] = util.UnescapeString(pair[1])
			}

			dict["post_data"] = post
		} else if r.Method == "GET" {
			content := r.RequestURI
			content = content[strings.Index(content, "?")+1:]

			entries := strings.Split(content, "&")
			get := make(map[string]string)

			for i := 0; i < len(entries); i++ {
				pair := strings.Split(entries[i], "=")
				get[util.UnescapeString(pair[0])] = util.UnescapeString(pair[1])
			}

			dict["get_data"] = get
		}

		return dict
	}
}

func SetCookieFn(w http.ResponseWriter) func(string, string, int, string) {
	return func(name, value string, age int, path string) {
		cookie := http.Cookie{}
		cookie.Name = name
		cookie.Value = value
		cookie.Path = path
		cookie.MaxAge = age

		http.SetCookie(w, &cookie)
	}
}

func GetCookieFn(r *http.Request) func(string) string {
	return func(name string) string {
		value := ""
		cookies := r.Cookies()

		for i := 0; i < len(cookies); i++ {
			if name == cookies[i].Name {
				value = cookies[i].Value
				break
			}
		}

		return value
	}
}

func DeleteCookieFn(w http.ResponseWriter) func(string, string) {
	return func(name string, path string) {
		cookie := http.Cookie{}
		cookie.Name = name
		cookie.Value = ""
		cookie.Path = path
		cookie.Expires = time.Unix(0, 0)

		http.SetCookie(w, &cookie)
	}
}
