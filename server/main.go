package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/ymolodyukov/otus/internal/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ymolodyukov/otus/internal/api"
	"github.com/ymolodyukov/otus/internal/model"

	_ "github.com/lib/pq"
)

func main() {
	// коннект к базе данных
	dbConn, err := sqlx.Connect("postgres", "user=otus password=12345678 dbname=otusdb sslmode=disable")
	if err != nil {
		log.Fatalf("db connection error %v \n", err)
	}

	// инициализация схемы базы данных
	dbConn.MustExec(db.Schema)

	mdl := model.New(dbConn)
	apiService := api.New(mdl)

	router := mux.NewRouter()
	router.HandleFunc("/login", apiService.Login).Methods("POST")
	router.HandleFunc("/user/register", apiService.RegisterUser).Methods("POST")
	router.HandleFunc("/user/get/{id:[a-z0-9-]+}", apiService.GetUserById).Methods("GET")

	http.Handle("/", router)

	port := "8080"
	log.Printf("server listening at port %v...\n", port)

	if err = http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("start http server error %v \n", err)
	}
}
