package auth

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
