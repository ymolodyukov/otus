package postgres

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/ymolodyukov/otus/internal/db/dto"
)

func (p *Person) RegisterUser(ctx context.Context, userData *dto.UserData, passwordHash string) (string, error) {
	userUUID, err := uuid.NewUUID()
	if err != nil {
		return "", errors.New("create user uuid")
	}

	userID := userUUID.String()
	tx := p.DB.MustBegin()

	query := "INSERT INTO person (id, firstname, lastname, age, sex, biography, city) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	tx.MustExec(query, userID, userData.FirstName, userData.LastName, userData.Age, userData.Sex, userData.Biography, userData.City)

	credID, err := uuid.NewUUID()
	if err != nil {
		return "", errors.New("create credentials uuid")
	}

	query = "INSERT INTO credentials (id, user_id, password) VALUES ($1, $2, $3)"
	tx.MustExec(query, credID, userID, passwordHash)

	tx.Commit()

	return userID, nil
}
