package test

import "gorm.io/gorm"

type Repository interface {
	GetTest(text string) string
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTest(text string) string {
	value := "Ini adalah value anda -> " + text
	return value
}
