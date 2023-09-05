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

func errorResponse(c *gin.Context, status int, errorMessage string) {
	c.JSON(status, domain.ErrorResponse{
		Status: status,
		Error:  errorMessage,
	})
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
	result := h.authService.GetUserByEmail(signUpRequest.Email)
	if result.Error == nil {
		existingUser := result.Value.(domain.User)
		if existingUser.Id != "" {
			errorResponse(c, http.StatusBadRequest, "Email already exists")
			return
		}
	}

	// Check if username already exists
	result = h.authService.GetUserByUsername(signUpRequest.Username)
	if result.Error == nil {
		existingUser := result.Value.(domain.User)
		if existingUser.Id != "" {
			errorResponse(c, http.StatusBadRequest, "Username already exists")
			return
		}
	}

	// Validate password length
	if len(signUpRequest.Password) < 8 {
		errorResponse(c, http.StatusBadRequest, "Password must be at least 8 characters long")
		return
	}

	user := domain.User{
		Email:    signUpRequest.Email,
		Username: signUpRequest.Username,
		Password: signUpRequest.Password,
		Role:     signUpRequest.Role,
	}

	err = h.authService.SignUp(user)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Failed to sign up")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
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
		errorResponse(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	user := result.Value.(domain.User)

	// Check password
	err = h.authService.CheckPassword(user.Password, signInRequest.Password)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// Generate and sign JWT token
	token, err := middlewares.GenerateToken(user.Id)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	c.JSON(http.StatusOK, domain.SignInResponse{
		Status:  http.StatusOK,
		Message: "Sign in successful",
		Data: domain.AuthResponse{
			Token: token,
		},
	})
}
