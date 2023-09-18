package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/ymolodyukov/otus/internal/otuserr"
)

func (p *Person) Login(ctx context.Context, userID string, passwordHash string) (string, error) {
	var credID string
	query := `SELECT id FROM credentials WHERE user_id=$1 AND password=$2`
	err := p.DB.Get(&credID, query, userID, passwordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", otuserr.ErrNotFound
		}

		return "", err
	}

	sessionToken, err := uuid.NewUUID()
	if err != nil {
		return "", errors.New("create session token error")
	}
	
	return sessionToken.String(), nil

}
