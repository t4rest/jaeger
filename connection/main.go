package main

import (
	"connection/api"
	"connection/util"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
)

// Module base module interface
type Module interface {
	Start()
	Stop()
	Title() string
}

func main() {

	util.InitJaeger("connection")

	apiModel := api.New()

	// Run modules
	RunModules(apiModel)
}

// RunModules runs each of the modules in a separate goroutine.
func RunModules(modules ...Module) {
	defer func() {
		for _, m := range modules {
			logrus.Infof("Stopping module %s", m.Title())
			m.Stop()
		}
		logrus.Infof("Stopped all modules")
	}()

	for _, m := range modules {
		logrus.Infof("Starting module %s", m.Title())
		go m.Start()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
