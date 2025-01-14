package repository

type UserRepository interface {
	FindByEmail()
	Save()
}

type UserRepositoryImpl struct {}

func (u UserRepositoryImpl) FindByEmail() {

}

func (u UserRepositoryImpl) Save() {
	
}