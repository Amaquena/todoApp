package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

var log = logrus.WithField("ctx", "config")

func ConfigureLogger(logLevel string) error {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			return "", filepath.Base(f.File) + ":" + strconv.Itoa(f.Line)
		},
	})

	logrus.SetReportCaller(true)
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return fmt.Errorf("unable to set log level: %s", err.Error())
	}
	log.Info("logger configured successfully")
	logrus.SetLevel(level)
	return nil
}
