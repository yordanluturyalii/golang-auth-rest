package service

import "yordanluturyali/golang-auth-rest/repository"

type AuthService interface {
	Login()
	Register()
	GoogleCallback()
}

type AuthServiceImpl struct {
	UserRepo repository.UserRepository
}

func (a AuthServiceImpl) Login() {
	
}

func (a AuthServiceImpl) Register() {
	
}

func (a AuthServiceImpl) GoogleCallback() {
	
}