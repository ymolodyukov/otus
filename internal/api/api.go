package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ymolodyukov/otus/internal/model"
	"github.com/ymolodyukov/otus/internal/otuserr"
)

// New конструктор сервиса
func New(model *model.Model) *service {
	return &service{
		model: model,
	}
}

type service struct {
	model *model.Model
}

func (x service) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		otuserr.SendInternalServerError(w)

		return
	}

	loginData := &model.LoginData{}
	if err := json.Unmarshal(rawBody, &loginData); err != nil {
		otuserr.SendInternalServerError(w)

		return
	}

	if err := validateLogin(loginData); err != nil {
		otuserr.SendBadRequest(w, err)

		return
	}

	token, err := x.model.Login(r.Context(), loginData.ID, loginData.Password)
	if err != nil {
		if errors.Is(err, otuserr.ErrNotFound) {
			otuserr.SendNotFound(w)
		} else {
			otuserr.SendInternalServerError(w)
		}

		return
	}

	otuserr.SendSuccess(w, token)

}

func (x service) RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		otuserr.SendInternalServerError(w)

		return
	}

	userData := &model.UserData{}
	if err := json.Unmarshal(rawBody, &userData); err != nil {
		otuserr.SendInternalServerError(w)

		return
	}

	if err := validateRegisterUser(userData); err != nil {
		otuserr.SendBadRequest(w, err)

		return
	}

	userId, err := x.model.RegisterUser(r.Context(), userData)
	if err != nil {
		otuserr.SendInternalServerError(w)

		return
	}

	otuserr.SendSuccess(w, userId)

}

func (x service) GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	if err := validateGetUserById(userID); err != nil {
		otuserr.SendBadRequest(w, err)

		return
	}

	userData, err := x.model.GetUserById(r.Context(), userID)
	if err != nil {
		if errors.Is(err, otuserr.ErrNotFound) {
			otuserr.SendNotFound(w)
		} else {
			log.Println(err)
			otuserr.SendInternalServerError(w)
		}

		return
	}

	otuserr.SendSuccess(w, userData)
}

func (x service) SearchUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	firstNamePrefix := query.Get("first_name")
	lastNamePrefix := query.Get("last_name")

	if err := validateSearchUsers(firstNamePrefix, lastNamePrefix); err != nil {
		otuserr.SendBadRequest(w, err)

		return
	}

	users, err := x.model.SearchUsers(r.Context(), firstNamePrefix, lastNamePrefix)
	if err != nil {
		if errors.Is(err, otuserr.ErrNotFound) {
			otuserr.SendNotFound(w)
		} else {
			log.Println("[ERROR] " + err.Error())
			otuserr.SendInternalServerError(w)
		}

		return
	}

	otuserr.SendSuccess(w, users)
}
