package controllers

import "yordanluturyali/golang-auth-rest/service"

type AuthController interface {
	Login()
	Register()
	GoogleCallback()
}

type AuthControllerImpl struct {
	UserService service.AuthService
}

func (c AuthControllerImpl) Login() {

}

func (c AuthControllerImpl) Register() {

}

func (c AuthControllerImpl) GoogleCallback() {

}