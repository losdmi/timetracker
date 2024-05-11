package handler

import (
	"net/http"
)

func (h *Handler) GetIndex(w http.ResponseWriter, r *http.Request) {
	err := h.templates.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
