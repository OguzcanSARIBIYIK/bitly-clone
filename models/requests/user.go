package requests

type UserRegister struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
