package auth

import (
	"fmt"
	"net/http"
	"repositoryPattern/domain"
	"repositoryPattern/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authHandler struct {
	authService AuthService
}

func NewAuthHandler(v1 *gin.RouterGroup, authService AuthService) {
	handler := &authHandler{authService}

	auth := v1.Group("auth")

	auth.POST("sign_up", handler.SignUp)
	auth.POST("sign_in", handler.SignIn)
}

func (h *authHandler) SignUp(c *gin.Context) {
	var signUpRequest domain.User

	err := c.ShouldBindJSON(&signUpRequest)
	if err != nil {

		errorMessages := []string{}

		for _, v := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", v.Field(), v.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	// Check if email already exists
	var existingUser domain.User
	result := h.authService.GetUserByEmail(signUpRequest.Email)
	if result.Error == nil {
		existingUser = result.Value.(domain.User)
		if existingUser.Id != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Email already exists",
			})
			return
		}
	}

	// Check if username already exists
	result = h.authService.GetUserByUsername(signUpRequest.Username)
	if result.Error == nil {
		existingUser = result.Value.(domain.User)
		if existingUser.Id != "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Username already exists",
			})
			return
		}
	}

	// Validate password length
	if len(signUpRequest.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password must be at least 8 characters long",
		})
		return
	}

	user := domain.User{
		Email:    signUpRequest.Email,
		Username: signUpRequest.Username,
		Password: signUpRequest.Password,
	}

	err = user.HashPassword()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	err = h.authService.SignUp(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to sign up",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sign up successful",
	})
}

func (h *authHandler) SignIn(c *gin.Context) {
	var signInRequest domain.AuthRequest

	err := c.ShouldBindJSON(&signInRequest)
	if err != nil {
		errorMessages := []string{}

		for _, v := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", v.Field(), v.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	// Get user by username
	result := h.authService.GetUserByUsername(signInRequest.Username)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	user := result.Value.(domain.User)

	// Check password
	err = user.CheckPassword(signInRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Generate and sign JWT token
	token, err := middlewares.GenerateToken(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
