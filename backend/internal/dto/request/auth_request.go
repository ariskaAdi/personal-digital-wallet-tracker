package request

type RegisterUserRequest struct {
	Name     string
	Email    string
	Password string
}

type LoginUserRequest struct {
	Email    string
	Password string
}