package app

import (
	"{{cookiecutter.project_name}}/config"
	"{{cookiecutter.project_name}}/log"

	"github.com/sirupsen/logrus"
)

// Application pre run
func PreRun() {
	// Create a new logger
	l := log.New(log.Config{
		Level:        config.LogLevel(),
		Format:       config.LogFormat(),
		File:         config.LogFile(),
		Console:      config.LogConsole(),
		LogstashType: config.LogstashType(),
	})
	// Always log these fields
	l.PersistentFields(logrus.Fields{
		"version":   Version(),
		"buildTime": BuildTime(),
	})
	// Set global logger
	log.SetGlobalLogger(l)
}

// Runs the Application
func Run() error {
	return nil // Do stuff
}
