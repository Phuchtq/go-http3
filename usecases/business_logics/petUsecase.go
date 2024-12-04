package businesslogics

import (
	"errors"
	"http3-integrate/constants"
	"http3-integrate/dtos/request"
	"http3-integrate/dtos/response"
	"http3-integrate/utils"
	"strings"
)

type IPetUsecase interface {
	GetAllPets() []response.Pet
	GetPetsByKw(kw string) *[]response.Pet
	GetPetById(id string) *response.Pet
	EditPet(req request.EditPetReq) error
	CreatePet(req request.CreatePetReq) error
	RemovePet(id string) error
}

type petUsecase struct{}

func GenerateService() IPetUsecase {
	return &petUsecase{}
}

func (p *petUsecase) GetAllPets() []response.Pet {
	return response.GetPets()
}

func (p *petUsecase) GetPetsByKw(kw string) *[]response.Pet {
	var pets = response.GetPets()
	var keyword string = utils.NormalizeString(kw)

	if kw == "" {
		return &pets
	}

	var res []response.Pet

	for _, pet := range pets {
		if strings.Contains(pet.Type, keyword) {
			res = append(res, pet)
		}
	}

	return &res
}

func (p *petUsecase) GetPetById(id string) *response.Pet {
	var res = response.GetPet(id)

	if res.ID == "" {
		return nil
	}

	return &res
}

func (p *petUsecase) EditPet(req request.EditPetReq) error {
	if !response.SetPet(req) {
		return errors.New("Invalid data")
	}

	return nil
}

func (p *petUsecase) CreatePet(req request.CreatePetReq) error {
	var errMsg error = errors.New(constants.GenericsInvalidDataErrMsg)

	if isPetExist(req.Gender, req.Type, response.GetPets()) {
		return errMsg
	}

	if !response.AddPet(response.Pet{
		ID:     utils.GenerateId(),
		Gender: req.Gender,
		Age:    req.Age,
		Amount: req.Amount,
		Weight: req.Weight,
		Type:   req.Type,
	}) {
		return errMsg
	}

	return nil
}

func (p *petUsecase) RemovePet(id string) error {
	if !response.RemovePet(id) {
		return errors.New(constants.GenericsInvalidDataErrMsg)
	}

	return nil
}

func isPetExist(gender, kind string, pets []response.Pet) bool {
	var lowerGend string = utils.NormalizeString(gender)
	var lowerKind string = utils.NormalizeString(kind)

	for _, pet := range pets {
		if pet.Gender == lowerGend && pet.Type == lowerKind {
			return true
		}
	}

	return false
}
