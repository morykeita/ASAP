package main

import (
	"fmt"
	"github.com/morykeita/ASAP/auth-service/internal/api"
	"github.com/morykeita/ASAP/auth-service/internal/config"
	database "github.com/morykeita/ASAP/auth-service/internal/database"
	"github.com/namsral/flag"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	flag.Parse()
	fmt.Println("")
	log.SetLevel(log.DebugLevel)
	log.WithField("version" , config.Version).Debug("starting server.")

	// creating a new databases

	db , err := database.New()
	if err != nil {
		log.WithError(err).Fatal("")
	}

	// creating a new router
	router, err := api.NewRouter(db)
	if err != nil {
		log.WithError(err).Fatal("Error writing version")
	}
	// move to yml config file or centralized config server
	const addr = "http://localhost:8080"
	server := http.Server{
		Handler:router,
		Addr : addr,
	}
	// starting server
	if err := server.ListenAndServe(); err != nil && err !=http.ErrServerClosed {
		log.WithError(err).Error("Server failed to start")
	}
}
