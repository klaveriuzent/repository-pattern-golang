package domain

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type SignInResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    AuthResponse `json:"data"`
}

type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
