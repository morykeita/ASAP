package main

import (
	"github.com/morykeita/ASAP/auth-service/internal/api"
	"github.com/morykeita/ASAP/auth-service/internal/config"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.WithField("version" , config.Version).Debug("starting server.")
	router, err := api.NewRouter()
	
	if err != nil {
		log.WithError(err).Fatal("Error writing version")
	}
	const addr = "0.0.0.0:8088"
	server := http.Server{
		Addr : addr,
	}

	if err := server.ListenAndServe();  err != nil && !http.ErrServerClosed {
		log.WithError(err).Error("Server failed to start.")
	}

}
