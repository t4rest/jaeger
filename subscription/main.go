package main

import (
	"os"
	"os/signal"
	"subscription/grpc_server"
	"subscription/util"

	"github.com/sirupsen/logrus"
)

// Module base module interface
type Module interface {
	Start()
	Stop()
	Title() string
}

func main() {

	util.InitJaeger("subscription")

	grpcModule := grpc_server.New()

	// Run modules
	RunModules(grpcModule)
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
