package util

import (
	"net/url"

	"github.com/nthnn/AnkoWeb/logger"
)

func UnescapeString(origin string) string {
	str, err := url.QueryUnescape(origin)
	if err != nil {
		logger.Error("Error: " + err.Error())
	}

	return str
}
