package Requests

type RegisterDto struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=5"`
}

var RegisterDtoMessage = map[string]string{
	"Username.required": "Username không được để trống.",
	"Username.min":      "Username quá ngắn.",
	"Username.max":      "Username quá dài.",
	"Password.required": "Mật khẩu không được để trống.",
	"Password.min":      "Mật khẩu không được quá ngắn.",
}
