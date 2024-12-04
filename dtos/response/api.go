package response

import "net/http"

type ApiResponseModel struct {
	Data   interface{}
	ErrMsg error
	Type   string
	W      http.ResponseWriter
	R      *http.Request
}
