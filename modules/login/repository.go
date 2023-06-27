package login

import (

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) Login(username, password string)  (*Admin, error) {
	var admin Admin
	result := repo.DB.Where("username", username).Where("password", password).First(&admin)
	return &admin, result.Error
}

func (repo Repository) GetAdminById(id int) (*Admin, error) {
	var admin *Admin
	result := repo.DB.First(&admin, id)
	return admin, result.Error
}