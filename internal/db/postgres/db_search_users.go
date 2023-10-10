package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ymolodyukov/otus/internal/db/dto"
	"github.com/ymolodyukov/otus/internal/db/postgres/colums"
	"github.com/ymolodyukov/otus/internal/otuserr"
	"log"
)

func (p *Person) SearchUsers(ctx context.Context, firstNamePrefix string, lastNamePrefix string) ([]*dto.UserData, error) {
	var persons []*dto.UserData
	arg := make(map[string]interface{})
	query := "SELECT id, first_name, last_name, age, sex, biography, city FROM person WHERE "
	where := ""
	if firstNamePrefix != "" {
		arg[colums.PersonFirstName] = firstNamePrefix + "%"
		where += colums.PersonFirstName + " LIKE :" + colums.PersonFirstName
	}
	if lastNamePrefix != "" {
		if where != "" {
			where += " AND "
		}
		where += colums.PersonLastName + " LIKE :" + colums.PersonLastName
		arg[colums.PersonLastName] = lastNamePrefix + "%"
	}

	query = query + where
	log.Println("[DEBUG] query: ", query)

	rows, err := p.DB.NamedQuery(query, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, otuserr.ErrNotFound
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		person := &dto.UserData{}
		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName,
			&person.Age, &person.Sex, &person.Biography, &person.City)
		if err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}
