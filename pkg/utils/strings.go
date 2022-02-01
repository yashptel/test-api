package utils

import (
	"net/http"
	"strconv"
)

func GetIntFromURL(r *http.Request, name string, defaultVal int) int {

	val := r.URL.Query().Get(name)

	num, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return num
}
