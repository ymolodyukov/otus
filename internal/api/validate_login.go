package api

import (
	"errors"

	"github.com/ymolodyukov/otus/internal/model"
)

func validateLogin(loginData *model.LoginData) error {
	if loginData == nil {
		return errors.New("empty login data")
	}

	if loginData.ID == "" {
		return errors.New("empty login")
	}

	if loginData.Password == "" {
		return errors.New("empty password")
	}

	return nil
}
