package main

import (
	"context"
	"github.com/todolist/handlers"
	"github.com/todolist/util"
	"os/signal"

	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create logger
	l := log.New(os.Stdout, "todo-api ", log.LstdFlags)

	// createDatabase connection
	err := util.CreateDBConnection(l, "./config")
	if err != nil {
		l.Printf("Error creating database connection\n%s", err)
		return
	}

	// Set up handlers
	itemsHandler := handlers.NewItems(l)
	r := mux.NewRouter()

	getRouters := r.Methods(http.MethodGet).Subrouter()
	getRouters.HandleFunc("/", itemsHandler.GetAllItems)
	getRouters.HandleFunc("/{id}", itemsHandler.GetSingleItems)

	postRouters := r.Methods(http.MethodPost).Subrouter()
	postRouters.HandleFunc("/", itemsHandler.AddItems)

	putRouters := r.Methods(http.MethodPut).Subrouter()
	putRouters.HandleFunc("/{id}", itemsHandler.UpdateItems)

	deleteRouters := r.Methods(http.MethodDelete).Subrouter()
	deleteRouters.HandleFunc("/{id}", itemsHandler.DeleteItems)

	// Starts the server
	server := createServer(r)
	go func() {
		l.Println("Starting server")
		err := server.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: \n", err)
			os.Exit(1)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	// blocks until a signal is received
	sig := <-ch
	l.Printf("Got signal: ", sig)
	var ctx, cancelFunc = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	if nil != server.Shutdown(ctx) {
		l.Println("Error Shutting down server gracefully")
	}
}

func createServer(r *mux.Router) http.Server {
	return http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      r,                 // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
}
