package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type Page struct {
	Items    any   `json:"items"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}

type AppError struct {
	Status  int
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Body{Code: 0, Message: "ok", Data: data})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, Body{Code: 0, Message: "ok", Data: data})
}

func Fail(c *gin.Context, err error) {
	var appErr AppError
	if errors.As(err, &appErr) {
		c.JSON(appErr.Status, Body{Code: appErr.Code, Message: appErr.Message})
		return
	}
	c.JSON(http.StatusInternalServerError, Body{Code: 500, Message: err.Error()})
}

func BadRequest(message string) AppError {
	return AppError{Status: http.StatusBadRequest, Code: 400, Message: message}
}

func Unauthorized(message string) AppError {
	return AppError{Status: http.StatusUnauthorized, Code: 401, Message: message}
}

func Forbidden(message string) AppError {
	return AppError{Status: http.StatusForbidden, Code: 403, Message: message}
}

func NotFound(message string) AppError {
	return AppError{Status: http.StatusNotFound, Code: 404, Message: message}
}
