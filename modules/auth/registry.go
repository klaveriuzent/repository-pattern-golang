package auth

import "gorm.io/gorm"

func AuthRegistry(db *gorm.DB) AuthService {
	userRepository := NewUserRepository(db)
	authService := NewAuthService(userRepository)

	return authService
}
