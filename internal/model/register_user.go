package model

import (
	"context"
	"github.com/ymolodyukov/otus/internal/cryptor"
)

func (m *Model) RegisterUser(ctx context.Context, userData *UserData) (string, error) {
	dbUserData := UserDataToDto(userData)

	passwordHash := cryptor.GetPasswordHash(userData.Password)

	userID, err := m.db.RegisterUser(ctx, dbUserData, passwordHash)
	if err != nil {
		return "", err
	}

	return userID, nil
}
