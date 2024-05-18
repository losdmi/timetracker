package handler

import (
	"net/http"

	modelDTO "github.com/losdmi/timetracker/internal/model/dto"
)

func (h *Handler) PostRecords(w http.ResponseWriter, r *http.Request) {
	_, form := h.service.AddRecord(
		r.Context(),
		modelDTO.NewAddRecordForm(
			r.FormValue("task"),
			r.FormValue("description"),
		),
	)

	err := h.templates.ExecuteTemplate(w, "form", form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
