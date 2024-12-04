package handlers

import (
	"http3-integrate/constants"
	"http3-integrate/dtos/request"
	"http3-integrate/dtos/response"
	business_logics "http3-integrate/usecases/business_logics"
	"http3-integrate/utils/api"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllPets(w http.ResponseWriter, r *http.Request) {
	if !api.IsMethodValid(http.MethodGet, r.Method) { // Invalid request method
		api.ProcessResponse(api.GenerateInvalidReqBody(w))
		return
	}

	api.ProcessResponse(response.ApiResponseModel{
		Data: business_logics.GenerateService().GetAllPets(),
		Type: constants.NonType,
	})
}

func GetPetsByKeyword(w http.ResponseWriter, r *http.Request) {
	if !api.IsMethodValid(http.MethodGet, r.Method) { // Invalid request method
		api.ProcessResponse(api.GenerateInvalidReqBody(w))
		return
	}

	api.ProcessResponse(response.ApiResponseModel{
		Data: business_logics.GenerateService().GetPetsByKw(mux.Vars(r)["keyword"]),
		Type: constants.NonType,
	})
}

func GetPetById(w http.ResponseWriter, r *http.Request) {
	if !api.IsMethodValid(http.MethodGet, r.Method) { // Invalid request method
		api.ProcessResponse(api.GenerateInvalidReqBody(w))
		return
	}

	api.ProcessResponse(response.ApiResponseModel{
		Data: business_logics.GenerateService().GetPetById(mux.Vars(r)["id"]),
		Type: constants.NonType,
	})
}

func CreatePet(w http.ResponseWriter, r *http.Request) {
	if !api.IsMethodValid(http.MethodPost, r.Method) { // Invalid request method
		api.ProcessResponse(api.GenerateInvalidReqBody(w))
		return
	}

	var data request.CreatePetReq
	if !api.IsRequestValid(r, data) {
		api.ProcessResponse(api.GenerateInvalidReqBody(w))
		return
	}

	api.ProcessResponse(response.ApiResponseModel{
		ErrMsg: business_logics.GenerateService().CreatePet(data),
	})
}

func EditPet(w http.ResponseWriter, r *http.Request) {
	if !api.IsMethodValid(http.MethodPut, r.Method) {
		api.ProcessResponse(api.GenerateInvalidReqBody(w))
		return
	}

	var data request.EditPetReq
	if !api.IsRequestValid(r, data) {
		api.ProcessResponse(api.GenerateInvalidReqBody(w))
		return
	}

	api.ProcessResponse(response.ApiResponseModel{
		ErrMsg: business_logics.GenerateService().EditPet(data),
	})
}

func RemovePet(w http.ResponseWriter, r *http.Request) {
	if !api.IsMethodValid(http.MethodDelete, r.Method) {
		api.ProcessResponse(api.GenerateInvalidReqBody(w))
		return
	}

	api.ProcessResponse(response.ApiResponseModel{
		ErrMsg: business_logics.GenerateService().RemovePet(mux.Vars(r)["id"]),
	})
}
