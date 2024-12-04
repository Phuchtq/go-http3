package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"http3-integrate/constants"
	"http3-integrate/dtos/response"
	"log"
	"net/http"
)

const (
	contentKey   string = "Content-Type"
	contentValue string = "application/json"
)

func ProcessResponse(data response.ApiResponseModel) {
	if data.ErrMsg != nil {
		processFailResponse(data.ErrMsg, data.W)
		return
	}

	if data.Type != constants.NonType {
		processSuccessResponse(data)
		return
	}

	processSuccessResponseData(data.Data, data.W)
}

func processFailResponse(err error, w http.ResponseWriter) {
	var errCode int

	switch err.Error() {
	case constants.InternalErrMsg:
		errCode = http.StatusInternalServerError
	case constants.GenericsInvalidDataErrMsg:
		errCode = http.StatusBadRequest

	// Other cases
	default:
		errCode = http.StatusBadRequest
	}

	processJson(w, generateMapData("message", err.Error()), errCode)
}

func processSuccessResponse(data response.ApiResponseModel) {
	switch data.Type {
	case constants.RedirectType:
		processRedirectResponse(data)
	case constants.InformType:
		processInformResponse(fmt.Sprint(data.Data), data.W)
	default:
		processJson(data.W, generateMapData("message", "success"), http.StatusOK)
	}
}

func processSuccessResponseData(data interface{}, w http.ResponseWriter) {
	processJson(w, generateMapData("data", data), http.StatusOK)
}

func processRedirectResponse(data response.ApiResponseModel) {
	http.Redirect(data.W, data.R, fmt.Sprint(data.Data), http.StatusPermanentRedirect)
}

func processInformResponse(msg string, w http.ResponseWriter) {
	processJson(w, generateMapData("message", msg), http.StatusAccepted)
}

func generateMapData(keyMsg string, value interface{}) map[string]interface{} {
	return map[string]interface{}{keyMsg: value}
}

func processJson(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set(contentKey, contentValue)
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf(constants.ProcessJsonErrMsg+"%v", data)
		http.Error(w, constants.InternalErrMsg, http.StatusInternalServerError)
	}
}

func GenerateInvalidReqBody(w http.ResponseWriter) response.ApiResponseModel {
	return response.ApiResponseModel{
		ErrMsg: errors.New(constants.GenericsInvalidDataErrMsg),
		W:      w,
	}
}
