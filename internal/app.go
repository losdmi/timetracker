package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/losdmi/timetracker/internal/handler"
)

type App struct {
	server *http.Server
}

func (a *App) Run() {
	fmt.Println("Starting server on port :8080")
	log.Fatal(a.server.ListenAndServe())
}

func BuildApp() *App {
	handler := handler.NewHandler()
	router := buildRoutes(handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return &App{
		server: server,
	}
}
