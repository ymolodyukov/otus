package db

import (
	"context"
	"github.com/ymolodyukov/otus/internal/db/dto"
)

type DB interface {
	// RegisterUser регистрация нового пользователя
	RegisterUser(ctx context.Context, data *dto.UserData, passwordHash string) (string, error)

	// Login логин пользователя в системе
	Login(ctx context.Context, userID string, password string) (string, error)

	// GetUserByID получение пользователя по ID
	GetUserByID(ctx context.Context, userID string) (*dto.UserData, error)

	// SearchUsers поиск пользователей по префексу имени и фамилии
	SearchUsers(ctx context.Context, firstNamePrefix string, lastNamePrefix string) ([]*dto.UserData, error)
}
