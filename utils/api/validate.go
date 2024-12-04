package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func IsRequestValid(req *http.Request, data interface{}) bool {
	body, err := io.ReadAll(req.Body)

	if err != nil {
		return false
	}

	if json.Unmarshal(body, &data) != nil {
		return false
	}

	return true
}

func IsMethodValid(expected, actual string) bool {
	return expected == actual
}
