package handler

import (
	"html/template"

	"github.com/losdmi/timetracker/internal/service"
)

type Handler struct {
	templates *template.Template
	service   *service.Service
}

func NewHandler(
	templates *template.Template,
	service *service.Service,
) *Handler {
	return &Handler{
		templates: templates,
		service:   service,
	}
}
