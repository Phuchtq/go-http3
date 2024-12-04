package response

import (
	"http3-integrate/dtos/request"
	"http3-integrate/utils"
)

type Pet struct {
	ID     string  `json:"id"`
	Gender string  `json:"gender"`
	Age    float32 `json:"age"`
	Weight float32 `json:"weight"`
	Amount int64   `json:"amount"`
	Type   string  `json:"type"`
}

const (
	cat   string = "cat"
	dog   string = "dog"
	bound int    = 30
)

var pets = []Pet{
	{
		ID:     "1",
		Gender: "Male",
		Age:    0.7,
		Weight: 2.5,
		Amount: 20,
		Type:   cat,
	},

	{
		ID:     "2",
		Gender: "Female",
		Age:    1.2,
		Weight: 4,
		Amount: 35,
		Type:   dog,
	},

	{
		ID:     "2",
		Gender: "Male",
		Age:    3.4,
		Weight: 4.4,
		Amount: 22,
		Type:   dog,
	},
}

func GetPets() []Pet {
	return pets
}

func GetPet(id string) Pet {
	for _, pet := range pets {
		if pet.ID == id {
			return pet
		}
	}

	return Pet{}
}

func SetPet(req request.EditPetReq) bool {
	if !utils.IsNumberValid[int64](req.Amount) {
		return false
	}

	for i, pet := range pets {
		if pet.ID == req.ID {
			pets[i].Amount = req.Amount
			return true
		}
	}

	return false
}

func AddPet(pet Pet) bool {
	if isAgeExceedBound(pet.Age) {
		return false
	}

	pets = append(pets, pet)
	return true
}

func RemovePet(id string) bool {
	var tmpStorage []Pet

	var res bool = false

	for _, pet := range pets {
		if pet.ID != id {
			tmpStorage = append(tmpStorage, pet)
		} else {
			res = true
		}
	}

	pets = tmpStorage
	return res
}

func isAgeExceedBound(age float32) bool {
	return age > float32(bound)
}
