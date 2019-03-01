package api

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// RequestLogger is used for standard logging
func RequestLogger(next httprouter.Handle, name string) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		start := time.Now()

		next(response, request, ps)

		logrus.WithFields(logrus.Fields{
			"method":   request.Method,
			"uri":      request.RequestURI,
			"name":     name,
			"duration": time.Since(start),
		}).Info()
	}
}
