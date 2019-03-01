package api

import (
	"encoding/json"
	"log"
	"net/http"
	"service1/events/publish"
	"service1/subprovider"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// API serves the end users requests.
type API struct {
	pub    publish.Publisher
	sp     subprovider.SubProvider
	Router *httprouter.Router
}

// New return new API instance
func New(pub publish.Publisher, sp subprovider.SubProvider) *API {
	return &API{pub: pub, sp: sp}
}

// Title returns the title.
func (api *API) Title() string {
	return "API"
}

// Start starts the http server and binds the handlers.
func (api *API) Start() {
	api.Router = httprouter.New()

	api.Router.GET("/health", api.Health)
	api.Router.POST("/notification", RequestLogger(api.NotificationHandler, "notification"))

	logrus.Infof("Listening on port :8081")
	err := http.ListenAndServe(":8081", api.Router)
	if err == nil {
		log.Fatal("Error can't launch the server on port :8081")
	}
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
