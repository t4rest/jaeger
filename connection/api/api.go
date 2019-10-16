package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/plugin/ochttp/propagation/b3"
)

// API serves the end users requests.
type API struct {
}

// New return new API instance
func New() *API {
	return &API{}
}

// Title returns the title.
func (api *API) Title() string {
	return "API"
}

// Start starts the http server and binds the handlers.
func (api *API) Start() {

	mux := http.NewServeMux()
	mux.Handle("/connections", RequestLogger(api.ConnectionHandler, "connection.ConnectionHandler"))
	mux.Handle("/connections/:id", RequestLogger(api.ConnectionHandler, "connection.ConnectionHandler"))

	log.Fatal(http.ListenAndServe(":8087", &ochttp.Handler{
		Handler:     mux,
		Propagation: &b3.HTTPFormat{},
	}))
}

// Stop stops server
func (api *API) Stop() {

}

// JSON writes to ResponseWriter a single JSON-object
func JSON(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		logrus.Error(err)
	}
}
