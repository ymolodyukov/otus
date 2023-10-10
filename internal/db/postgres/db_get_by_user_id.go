package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ymolodyukov/otus/internal/db/dto"
	"github.com/ymolodyukov/otus/internal/db/postgres/colums"
	"github.com/ymolodyukov/otus/internal/otuserr"
)

func (p *Person) GetUserByID(ctx context.Context, userID string) (*dto.UserData, error) {
	person := dto.UserData{}
	query := `SELECT ` +
		colums.PersonId + `, ` + colums.PersonFirstName + `, ` + colums.PersonLastName +
		`, ` + colums.PersonAge + `, ` + colums.PersonSex + `,	` + colums.PersonBiography + `, ` + colums.PersonCity +
		` FROM person WHERE id=$1`

	err := p.DB.Get(&person, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, otuserr.ErrNotFound
		}

		return nil, err
	}

	return &person, nil
}
