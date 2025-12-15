package request

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

type UpdateUserRequest struct {
	Name     string
	Email    string
	Password string
}