package handler

import (
	"net/http"

	modelDTO "github.com/losdmi/timetracker/internal/model/dto"
)

func (h *Handler) PostRecords(w http.ResponseWriter, r *http.Request) {
	response := h.service.AddRecord(
		r.Context(),
		modelDTO.NewAddRecordForm(
			r.FormValue("task"),
			r.FormValue("description"),
		),
	)

	err := h.templates.ExecuteTemplate(w, "form", response.Form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if len(response.Records) == 0 {
		return
	}

	err = h.templates.ExecuteTemplate(w, "records", response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
