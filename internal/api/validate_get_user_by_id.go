package api

import (
	"errors"

	"github.com/google/uuid"
)

func validateGetUserById(userID string) error {
	_, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("invalid user id")
	}

	return nil
}
