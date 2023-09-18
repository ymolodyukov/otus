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
	ID         string `json:"id,omitempty"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Age        int    `json:"age"`
	Sex        string `json:"sex"`
	Biography  string `json:"biography"`
	City       string `json:"city"`
	Password   string `json:"password,omitempty"`
}

func UserDataFromDto(dtoData *dto.UserData) *UserData {
	userData := &UserData{
		ID:         dtoData.ID,
		FirstName:  dtoData.FirstName,
		SecondName: dtoData.SecondName,
		Age:        dtoData.Age,
		Sex:        dtoData.Sex,
		Biography:  dtoData.Biography,
		City:       dtoData.City,
	}

	return userData
}

func UserDataToDto(userData *UserData) *dto.UserData {
	dbData := &dto.UserData{
		ID:         userData.ID,
		FirstName:  userData.FirstName,
		SecondName: userData.SecondName,
		Age:        userData.Age,
		Sex:        userData.Sex,
		Biography:  userData.Biography,
		City:       userData.City,
	}

	return dbData
}
