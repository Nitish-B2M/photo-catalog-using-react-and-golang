package requests

type UserRequest struct {
	RequestUsername string `json:"username" validate:"required,min=3,max=30"`
	RequestEmail    string `json:"email" validate:"required,email"`
	RequestPassword string `json:"password" validate:"required,min=6,max=100"`
}

type UserLoginRequest struct {
	RequestEmail    string `json:"email" validate:"required,email"`
	RequestPassword string `json:"password" validate:"required,min=6,max=100"`
}
