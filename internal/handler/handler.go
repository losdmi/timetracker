package handler

import "html/template"

type Handler struct {
	templates *template.Template
}

func NewHandler() *Handler {
	return &Handler{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
}
