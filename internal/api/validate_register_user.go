package api

import (
	"errors"

	"github.com/ymolodyukov/otus/internal/model"
)

func validateRegisterUser(userData *model.UserData) error {
	if userData == nil {
		return errors.New("empty user data")
	}
	if userData.FirstName == "" {
		return errors.New("empty first name")
	}
	if userData.LastName == "" {
		return errors.New("empty second name")
	}
	if userData.Age <= 0 {
		return errors.New("age should be greater than zero")
	}
	if userData.Sex != "m" && userData.Sex != "w" {
		return errors.New("sex should be 'm' or 'w'")
	}
	if userData.Biography == "" {
		return errors.New("empty biography")
	}
	if userData.City == "" {
		return errors.New("empty city")
	}
	if userData.Password == "" {
		return errors.New("empty password")
	}

	return nil
}
