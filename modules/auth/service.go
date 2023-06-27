package auth

import (
	"errors"
	"repositoryPattern/domain"
)

type AuthService interface {
	SignUp(user domain.User) error
	GetUserByEmail(email string) Result
	GetUserByUsername(username string) Result
}

type authService struct {
	userRepository UserRepository
}

func NewAuthService(userRepository UserRepository) AuthService {
	return &authService{userRepository}
}

func (s *authService) SignUp(user domain.User) error {
	err := s.userRepository.SignUp(user)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.email" {
			return errors.New("email already exists")
		} else if err.Error() == "UNIQUE constraint failed: users.username" {
			return errors.New("username already exists")
		} else {
			return errors.New("failed to sign up")
		}
	}
	return nil
}

func (s *authService) GetUserByEmail(email string) Result {
	var user domain.User
	result := s.userRepository.GetByEmail(email)
	if result.Error == nil {
		user = result.Value.(domain.User)
	}
	return Result{user, result.Error}
}

func (s *authService) GetUserByUsername(username string) Result {
	var user domain.User
	result := s.userRepository.GetByUsername(username)
	if result.Error == nil {
		user = result.Value.(domain.User)
	}
	return Result{user, result.Error}
}
