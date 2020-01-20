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
	return router, nil
}
