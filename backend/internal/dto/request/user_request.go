package request

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

type UpdateUserRequest struct {
	Id       int
	Name     string
	Email    string
	Password string
}