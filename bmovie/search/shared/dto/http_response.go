package dto

type (
	HttpRespBody struct {
		Code    string      `json:"response_code"`
		Message string      `json:"response_message"`
		Data    interface{} `json:"data"`
	}
)
