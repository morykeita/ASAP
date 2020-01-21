package api

import (
	database2 "github.com/morykeita/ASAP/auth-service/internal/database"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/morykeita/ASAP/auth-service/internal/api/v1"
)

//NewRouter provide a handler API service.
func NewRouter(db database2.Database) (http.Handler, error) {
	router := mux.NewRouter()
	router.HandleFunc("/version", v1.VersionHandler)
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	userApi := &v1.UserAPI{
		DB:db,
	}
	apiRouter.HandleFunc("/users",userApi.Create).Methods(http.MethodPost)
	//apiRouter.HandleFunc("/users",userApi.Create).Methods(http.MethodGet) // get user
	//apiRouter.HandleFunc("/users",userApi.Create).Methods(http.MethodPut) // update user
	//apiRouter.HandleFunc("/users",userApi.Create).Methods(http.MethodDelete) // delete user
	return router, nil
}
