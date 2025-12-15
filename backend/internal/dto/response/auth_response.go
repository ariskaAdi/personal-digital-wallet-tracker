package response

type AuthResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type LoginResponse struct {
	Token string `json:"token"`
}