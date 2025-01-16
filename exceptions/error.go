package exceptions

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

type GlobalError interface {
	Error() string
	GetCode() int
}

type HttpBadRequestError struct {
	Msg  string
	Code int
}

func NewHttpBadRequestError(msg string) HttpBadRequestError {
	return HttpBadRequestError{Msg: msg, Code: http.StatusBadRequest}
}

func (e HttpBadRequestError) Error() string {
	return e.Msg
}

func (e HttpBadRequestError) GetCode() int {
	return e.Code
}

type HttpNotFoundError struct {
	Msg  string
	Code int
}

func NewHttpNotFoundError(msg string) HttpNotFoundError {
	return HttpNotFoundError{Msg: msg, Code: http.StatusNotFound}
}

func (e HttpNotFoundError) Error() string {
	return e.Msg
}

func (e HttpNotFoundError) GetCode() int {
	return e.Code
}

type HttpForbiddenError struct {
	Msg  string
	Code int
}

func NewHttpForbiddenError() HttpForbiddenError {
	return HttpForbiddenError{Msg: "Forbidden", Code: http.StatusForbidden}
}

func (e HttpForbiddenError) Error() string {
	return e.Msg
}

func (e HttpForbiddenError) GetCode() int {
	return e.Code
}

type UnauthorizedError struct {
	Msg  string
	Code int
}

func NewUnauthorizedError(msg string) UnauthorizedError {
	return UnauthorizedError{Msg: msg, Code: http.StatusUnauthorized}
}

func (e UnauthorizedError) Error() string {
	return e.Msg
}

func (e UnauthorizedError) GetCode() int {
	return e.Code
}

type HttpInternalServerError struct {
	Msg  string
	Code int
}

func NewHttpInternalServerError() HttpInternalServerError {
	return HttpInternalServerError{Msg: "Internal Server Error", Code: http.StatusInternalServerError}
}

func (e HttpInternalServerError) Error() string {
	return e.Msg
}

func (e HttpInternalServerError) GetCode() int {
	return e.Code
}

type FailedValidation struct {
	Msg  string
	Code int
	Errors map[string]interface{}
}

func NewFailedValidation(obj interface{}, err validator.ValidationErrors) FailedValidation {
	objRef := reflect.TypeOf(obj)
	errors := make(map[string]interface{})

	for _, fieldError := range err {
		fieldName, _ := objRef.FieldByName(fieldError.Field())
		field := fieldName.Tag.Get("json")
		errors[field] = handleValidationErrorMessage(fieldError.Tag(), fieldError.Param(), field)
	}

	return FailedValidation{Msg: "Failed Validation", Code: http.StatusUnprocessableEntity, Errors: errors}
}

func (e FailedValidation) Error() string {
	return e.Msg
}

func (e FailedValidation) GetCode() int {
	return e.Code
}

func handleValidationErrorMessage(tag string, param string,field string) string {
	var msg string
	field = strings.Replace(field, "_", " ", -1)
	
	switch tag {
	case "required":
		msg = fmt.Sprintf("The %s field is required", field)
	case "email":
		msg = "This is not valid email"
	case "min": 
		msg = fmt.Sprintf("The %s field min value %s", tag, param)
	case "max":
		msg = fmt.Sprintf("The %s field max value %s", tag, param)
	}

	return msg
} 