package api

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// RequestLogger is used for standard logging
func RequestLogger(next http.HandlerFunc, name string) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()

		next(response, request)

		logrus.WithFields(logrus.Fields{
			"method":   request.Method,
			"uri":      request.RequestURI,
			"name":     name,
			"duration": time.Since(start),
		}).Info()
	}
}
