package handler

import (
	"net/http"
)

func (h *Handler) GetIndex(w http.ResponseWriter, r *http.Request) {
	indexData := h.service.GetIndexData(r.Context())

	err := h.templates.ExecuteTemplate(w, "index", indexData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
