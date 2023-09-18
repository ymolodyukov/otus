package model

import (
	"context"
)

func (m *Model) GetUserById(ctx context.Context, userID string) (*UserData, error) {
	dbData, err := m.db.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	userData := UserDataFromDto(dbData)

	return userData, nil
}
