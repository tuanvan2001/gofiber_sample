package Services

import (
	"fmt"
	"goSample/Models"
	"goSample/Types/Messages"
	"goSample/Types/Requests"
)

type LoginResponse struct {
	Token string      `json:"token"`
	User  Models.User `json:"user"`
}

func Login(loginDto *Requests.LoginDto) (LoginResponse, error) {
	// Check if the user already exists
	user, err := FindUserByUsername(loginDto.Username)
	if err != nil {
		return LoginResponse{}, fmt.Errorf(Messages.User["NotFound"])
	}

	token, _ := CreateTokenJWT(int(user.ID), user.Username)

	return LoginResponse{
		Token: token,
		User:  user,
	}, nil
}
