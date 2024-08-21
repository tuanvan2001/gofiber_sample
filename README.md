## Directory Structure:

```json
{
  "Configs": "chứa các Cấu hình như mysql, redis, loadEnv, cors, ...",
  "Controllers": "xử lý các request http từ client.",
  "Middlewares": "chứa các hàm xử lý request http trước khi đẩy xuống Controller",
  "Models": "này chứa các model ánh xạ với các table trong mysql",
  "Routes": "Cấu hình router",
  "Services": "xử lý logic hoăc tương tác với các dịch vụ bên ngoài.",
  "main.go": "File khởi tạo server",
  "Types": {
    "Http": "Cấu hình các struct HTTP",
    "Messages": "Cấu hình các message trả về cho client",
    "Requests": "Defines the body & validation rules for POST HTTP Request."
  }
}
```
```json
{
  "Configs": "contains configurations such as mysql, redis, loadEnv, cors, etc.",
  "Controllers": "handles HTTP requests from clients.",
  "Middlewares": "contains functions that process HTTP requests before passing them to Controllers.",
  "Models": "contains models that map to tables in MySQL.",
  "Routes": "configures routing.",
  "Services": "handles logic or interacts with external services.",
  "main.go": "The file that initializes the server.",
  "Types": {
    "Http": "Defines HTTP structs.",
    "Messages": "Defines messages returned to clients.",
    "Requests": "Defines the body & validation rules for POST HTTP Requests."
  }
}
```
## LifeCycle:

```
Client --(send request)--> Router --> Middleware --> Controller --> Service
```

## Example create a resource user:

#### 1. Create file Services/User.go

```go
package Services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"goSample/Configs"
	"goSample/Models"
	"goSample/Types/Messages"
	"goSample/Types/Requests"
)

var ctx = context.Background()

func CreateUser(createUserDto *Requests.CreateUser) (int64, error) {
	// Check if the user already exists
	if _, err := FindUserByUsername(createUserDto.Username); err == nil {
		return 0, fmt.Errorf(UserMessages.Vi["Exist"])
	}
	// Create a new user
	newUser := &Models.User{
		UUID:        uuid.New(),
		Username:    createUserDto.Username,
		Password:    createUserDto.Password,
		PhoneNumber: createUserDto.PhoneNumber,
		FullName:    createUserDto.FullName,
		Email:       createUserDto.Email,
		Age:         createUserDto.Age,
		Birthday:    createUserDto.Birthday,
	}
	// Save the new user to the database
	result := Configs.MySQL.Create(newUser)
	if result.Error != nil {
		return 0, fmt.Errorf(UserMessages.Vi["CreateFail"])
	}

	// Serialize the new user to JSON
	userJSON, err := json.Marshal(newUser)
	if err != nil {
		return int64(newUser.ID), fmt.Errorf(UserMessages.Vi["CacheFail"])
	}

	// Cache the new user in Redis
	err = Configs.Redis.Set(ctx, fmt.Sprintf("user:%d", newUser.ID), userJSON, 0).Err()
	if err != nil {
		return int64(newUser.ID), fmt.Errorf(UserMessages.Vi["CacheFail"])
	}
	return int64(newUser.ID), nil
}

func FindUser(userId int) (Models.User, error) {
	var user Models.User
	result := Configs.MySQL.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		return Models.User{}, result.Error
	}
	return user, nil
}

func FindUserByUsername(username string) (Models.User, error) {
	var user Models.User
	result := Configs.MySQL.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return Models.User{}, result.Error
	}
	return user, nil
}
```

#### 2. Create file Types/Messages/User.go

```go
package UserMessages

var Vi = map[string]string{
	"Created":           "Tạo user thành công.",
	"CreateFail":        "Tạo user thất bại.",
	"NotFound":          "Không tìm thấy user.",
	"Exist":             "User đã tồn tại.",
	"CacheFail":         "Cache user that bai.",
	"CacheRetrieveFail": "Truy xuat cache that bai.",
}
```

#### 3. Create file Types/Requests/CreateUser.go

```go
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

```

#### 4. Create file Controllers/User.go

```go
package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"goSample/Middlewares"
	"goSample/Services"
	"goSample/Types/Http"
	"goSample/Types/Messages"
	"goSample/Types/Requests"
)

func CreateUser(ctx *fiber.Ctx) error {
	var createUserDto *Requests.CreateUser
	// check body & validate body
	if err := ctx.BodyParser(&createUserDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if errMsgs, hasErrors := Middlewares.Validator.Validate(createUserDto, Requests.CreateUserMessage); hasErrors {
		return Http.CreateHttpErrorValidate(errMsgs)
	}
	//call service create user
	id, errCreateUser := Services.CreateUser(createUserDto)
	if errCreateUser != nil {
		return Http.CreateHttpError(fiber.StatusBadRequest, errCreateUser.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": UserMessages.Vi["Created"],
		"data": fiber.Map{
			"userID": id,
		},
	})
}
```