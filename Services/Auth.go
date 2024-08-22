package Services

import (
	"fmt"
	"goSample/Types/Messages"
	"goSample/Types/Requests"
	"log"
)

func Login(loginDto *Requests.LoginDto) (any, error) {
	// Check if the user already exists
	user, err := FindUserByUsername(loginDto.Username)
	if err != nil {
		return 0, fmt.Errorf(Messages.User["Exist"])
	}
	log.Println(user)
	log.Println(err)
	return user, nil
}
