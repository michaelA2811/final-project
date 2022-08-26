package entity

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Message *string     `json:"message,omitempty"`
	Error   *string     `json:"error,omitempty"`
	Data    interface{} `json:"data",omitempty`
}

func CreatedSuccess(msg *string, data interface{}) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: msg,
		Data:    data,
	}
}

func FindSuccess(msg *string, data interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: msg,
		Data:    data,
	}
}

func NotFound(msg, err *string, data interface{}) *Response {
	return &Response{
		Status: http.StatusNotFound,
		Error:  err,
		Data:   data,
	}
}

func InternalServerError(msg string, err *string, data interface{}) *Response {
	return &Response{
		Status: http.StatusInternalServerError,
		Error:  err,
		Data:   data,
	}
}
