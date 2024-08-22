package Services

import (
	"fmt"
	"goSample/Configs"
	"goSample/Models"
	"goSample/Types/Messages"
	"goSample/Types/Requests"
)

func CreateUser(createUserDto *Requests.CreateUser) (int64, error) {
	// Check if the user already exists
	if _, err := FindUserByUsername(createUserDto.Username); err == nil {
		return 0, fmt.Errorf(Messages.User["Exist"])
	}

	// Create a new user
	newUser := &Models.User{
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
		return 0, fmt.Errorf(Messages.User["CreateFail"])
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
