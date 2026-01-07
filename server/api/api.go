package api

import (
	"database/sql"

	"github.com/princetheprogrammerbtw/TheBiCycleApp/server/config"
)

// API holds the dependencies for our application.
type API struct {
	DB  *sql.DB
	Cfg *config.Config
}

// NewAPI creates a new API structure
func NewAPI(db *sql.DB, cfg *config.Config) *API {
	api := &API{DB: db, Cfg: cfg}
	api.setupGoogleOauth()
	return api
}
