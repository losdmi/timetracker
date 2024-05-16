package internal

import (
	"net/http"

	"github.com/losdmi/timetracker/internal/handler"
)

func buildRoutes(h *handler.Handler) http.Handler {
	router := http.NewServeMux()

	router.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("GET /", h.GetIndex)
	router.HandleFunc("POST /records", h.PostRecords)

	return router
}
