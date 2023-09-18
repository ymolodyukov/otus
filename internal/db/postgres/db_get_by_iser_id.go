package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ymolodyukov/otus/internal/db/dto"
	"github.com/ymolodyukov/otus/internal/otuserr"
)

func (p *Person) GetUserByID(ctx context.Context, userID string) (*dto.UserData, error) {
	person := dto.UserData{}
	query := `SELECT id, first_name, second_name, age, sex,	biography, city FROM person WHERE id=$1`
	err := p.DB.Get(&person, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, otuserr.ErrNotFound
		}

		return nil, err
	}

	return &person, nil
}
