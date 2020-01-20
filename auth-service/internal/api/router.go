package api

import {
	"net/http"
	"github.com/gorilla/mux"
}
//NewRouter provide a handler API service.
func NewRouter(http.Handle,error)  {
	router := mux.NewRouter()
	router.HandleFunc("/version",v1.VersionHandler)
	return router, nil
}