package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/ymolodyukov/otus/internal/db"
	"github.com/ymolodyukov/otus/internal/db/dto"
	"github.com/ymolodyukov/otus/internal/db/postgres"
)

// New конструктор модели сервиса
func New(db *sqlx.DB) *Model {
	pg := &postgres.Person{
		DB: db,
	}

	return &Model{
		db: pg,
	}
}

type Model struct {
	db db.DB
}

type LoginData struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type UserData struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Sex       string `json:"sex"`
	Biography string `json:"biography"`
	City      string `json:"city"`
	Password  string `json:"password,omitempty"`
}

func UserDataFromDto(dtoData *dto.UserData) *UserData {
	userData := &UserData{
		ID:        dtoData.ID,
		FirstName: dtoData.FirstName,
		LastName:  dtoData.LastName,
		Age:       dtoData.Age,
		Sex:       dtoData.Sex,
		Biography: dtoData.Biography,
		City:      dtoData.City,
	}

	return userData
}

func UserDataToDto(userData *UserData) *dto.UserData {
	dbData := &dto.UserData{
		ID:        userData.ID,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Age:       userData.Age,
		Sex:       userData.Sex,
		Biography: userData.Biography,
		City:      userData.City,
	}

	return dbData
}
