package model

import (
	"context"
	"strings"
)

func (m *Model) SearchUsers(ctx context.Context, firstNamePrefix string, lastNamePrefix string) ([]*UserData, error) {
	dbData, err := m.db.SearchUsers(ctx, strings.TrimSpace(firstNamePrefix), strings.TrimSpace(lastNamePrefix))
	if err != nil {
		return nil, err
	}
	users := make([]*UserData, 0, len(dbData))

	for _, userData := range dbData {
		user := UserDataFromDto(userData)
		users = append(users, user)
	}

	return users, nil
}
