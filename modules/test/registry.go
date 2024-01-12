package test

import "gorm.io/gorm"

func TestRegistry(db *gorm.DB) *service {
	repository := NewRepository(db)
	returnService := NewService(repository)
	return returnService
}
