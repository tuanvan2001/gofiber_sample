package Requests

type CreateUser struct {
	Username    string `json:"username" validate:"required,min=5,max=20"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	FullName    string `json:"full_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Age         uint8  `json:"age" validate:"required,min=1,max=150"`
	Birthday    string `json:"birthday" validate:"required"`
}

var CreateUserMessage = map[string]string{
	"Username.required":    "Username không được để trống.",
	"Username.min":         "Username quá ngắn.",
	"Username.max":         "Username quá dài.",
	"Password.required":    "Mật khẩu không được để trống.",
	"PhoneNumber.required": "Số điện thoại không được để trống.",
	"FullName.required":    "Họ và tên không được để trống.",
	"Email.required":       "Email không được để trống.",
	"Email.email":          "Email không hợp lệ.",
	"Age.required":         "Tuổi không được để trống.",
	"Age.min":              "Tuổi phải lớn hơn 0.",
	"Age.max":              "Tuổi phải nhỏ hơn 150.",
	"Birthday.required":    "Ngày sinh không được để trống.",
	"Birthday.date":        "Ngày sinh không hợp lệ.",
}
