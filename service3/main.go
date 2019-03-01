package main

import (
	"os"
	"os/signal"
	"service3/events/subscribe"
	"service3/util"

	"github.com/sirupsen/logrus"
)

// Module base module interface
type Module interface {
	Start()
	Stop()
	Title() string
}

func main() {

	util.InitJaeger("service3")

	eventsModule := subscribe.New()

	// Run modules
	RunModules(eventsModule)
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
