package register

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) Register(register *Admin) error {
	result := repo.DB.Create(&register)
	return result.Error
}