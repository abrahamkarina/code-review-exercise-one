package web

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
