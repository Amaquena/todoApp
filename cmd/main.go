package main

import (
	"github.com/sirupsen/logrus"
	"github.com/todoApp/pkg/api/todo"
	"github.com/todoApp/pkg/config"
	"github.com/todoApp/pkg/server"
	"github.com/todoApp/pkg/storage"
	"os"
	"os/signal"
	"time"
)

var log = logrus.WithField("ctx", "main")

const httpShutdownTimeout = 10 * time.Second

func main() {
	// Load config
	conf, err := config.NewConfig()
	if err != nil {
		log.WithError(err).Fatalf("error loading configs")
	}

	// Configure logger
	err = config.ConfigureLogger(conf.Application.LogLevel)
	if err != nil {
		log.WithError(err).Fatalf("error configuring logger")
	}

	// Establish Database connection
	conn, err := storage.NewDBConnection(conf)
	if err != nil {
		log.WithError(err).Fatalf("Error establishing database connection")
	}

	// Setup Database
	db, err := storage.NewSQLStore(conn)
	if err != nil {
		log.WithError(err).Fatal("failed to setup database tables")
	}
	log.Info("MySQL database connected successfully established")

	// Initialise the Todo Service Api
	todoService := todo.NewTodoService(db, conf.Application.RequestTimeOut)

	// Starts the server
	todoServer := server.NewServer(todoService)
	todoServer.Serve(&conf.Server)
	log.WithFields(logrus.Fields{
		"host": conf.Server.Host,
		"port": conf.Server.Port,
	}).Info("ToDo API service successfully started")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	//blocks until a signal is received
	sig := <-ch
	log.Infof("got signal: %s", sig.String())
	//var ctx, cancelFunc = context.WithTimeout(context.Background(), 30*time.Second)
	//defer cancelFunc()
	db.Close()
	todoServer.Shutdown()
	log.Info("Gracefully shutting down Todo API Service")
}
