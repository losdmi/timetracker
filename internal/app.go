package internal

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/losdmi/timetracker/internal/handler"
	"github.com/losdmi/timetracker/internal/repository"
	"github.com/losdmi/timetracker/internal/service"
)

type App struct {
	server *http.Server
	dbpool *pgxpool.Pool
}

func (a *App) Run() {
	fmt.Println("Starting server on port :8080")
	log.Fatal(a.server.ListenAndServe())
}

func BuildApp() *App {
	ctx := context.Background()

	templates := template.Must(template.ParseGlob("view/*.html"))

	dbpool, err := pgxpool.New(ctx, os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("ошибка при соединении с БД: %s", err)
	}

	repository := repository.NewRepository(dbpool)
	service := service.NewService(repository)

	handler := handler.NewHandler(templates, service)
	router := buildRoutes(handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return &App{
		server: server,
		dbpool: dbpool,
	}
}
