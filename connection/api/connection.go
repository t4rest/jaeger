package api

import (
	"net/http"
)

// Connection .
type Connection struct {
	ConnectionID string `json:"connection_id"`
}

// ConnectionHandler
func (api *API) ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	con := Connection{
		ConnectionID: "ConnectionID",
	}

	JSON(w, con)
}
