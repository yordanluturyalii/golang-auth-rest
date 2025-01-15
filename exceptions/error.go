package exceptions

import (
	"fmt"
	"net/http"
	"reflect"
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
		fieldName := fieldError.Field()

		if f, ok := objRef.Elem().FieldByName(fieldName); ok {
			jsonTag := f.Tag.Get("json")
			if jsonTag != "" {
				fieldName = jsonTag
			}
		}

		errors[fieldName] = fmt.Sprintf("%s", fieldError.Tag())
	}

	return FailedValidation{Msg: "Failed Validation", Code: http.StatusUnprocessableEntity, Errors: errors}
}

func (e FailedValidation) Error() string {
	return e.Msg
}

func (e FailedValidation) GetCode() int {
	return e.Code
}