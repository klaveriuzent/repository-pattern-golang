package auth

import (
	"repositoryPattern/domain"
	"repositoryPattern/helper"

	"gorm.io/gorm"
)

type UserRepository interface {
	SignUp(user domain.User) error
	GetByEmail(email string) Result
	GetByUsername(username string) Result
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) SignUp(user domain.User) error {
	// Generate ID
	userID, _ := helper.GenerateUserId(3)

	userEntity := domain.User{
		Id:       userID,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}
	result := r.db.Create(&userEntity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *userRepository) GetByEmail(email string) Result {
	var user domain.User
	result := r.db.Where("email = ?", email).First(&user)
	return Result{user, result.Error}
}

func (r *userRepository) GetByUsername(username string) Result {
	var user domain.User
	result := r.db.Where("username = ?", username).First(&user)
	return Result{user, result.Error}
}

type Result struct {
	Value interface{}
	Error error
}
