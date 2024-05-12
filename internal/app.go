package internal

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/losdmi/timetracker/internal/handler"
	"github.com/losdmi/timetracker/internal/service"
)

type App struct {
	server *http.Server
}

func (a *App) Run() {
	fmt.Println("Starting server on port :8080")
	log.Fatal(a.server.ListenAndServe())
}

func BuildApp() *App {
	templates := template.Must(template.ParseGlob("view/*.html"))
	service := service.NewService()

	handler := handler.NewHandler(templates, service)
	router := buildRoutes(handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return &App{
		server: server,
	}
}
