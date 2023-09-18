package model

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ymolodyukov/otus/internal/cryptor"
	"github.com/ymolodyukov/otus/internal/otuserr"
)

func (m *Model) Login(ctx context.Context, userID string, password string) (string, error) {
	passwordHash := cryptor.GetPasswordHash(password)
	token, err := m.db.Login(ctx, userID, passwordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", otuserr.ErrNotFound
		}

		return "", err
	}

	return token, nil
}
