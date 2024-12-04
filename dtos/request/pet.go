package request

type EditPetReq struct {
	ID     string `json:"id"`
	Amount int64  `json:"amount"`
}

type CreatePetReq struct {
	Gender string  `json:"gender"`
	Age    float32 `json:"age"`
	Weight float32 `json:"weight"`
	Amount int64   `json:"amount"`
	Type   string  `json:"type"`
}
