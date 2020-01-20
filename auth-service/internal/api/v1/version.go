package v1

import (
	"encoding/json"
	"github.com/morykeita/ASAP/auth-service/internal/config"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// API for returning version
// When server starts, set version and then use it if necessary

type ServerVersion struct {
	version string "json : 'version'"
}

// marshalled json
var versionJSON []byte

func init() {
	var err error
	versionJSON, err = json.Marshal(ServerVersion{
		version: config.Version,
	})

	if err != nil {
		panic(err)
	}
}

func VersionHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	if _, err := w.Write(versionJSON); err != nil {
		log.WithError(err).Debug("Error writing version")
	}
}
