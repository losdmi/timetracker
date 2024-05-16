package handler

import (
	"net/http"

	"github.com/losdmi/timetracker/internal/model"
)

func (h *Handler) PostRecords(w http.ResponseWriter, r *http.Request) {
	_, form := h.service.AddRecord(
		model.NewAddRecordForm(
			r.FormValue("task"),
			r.FormValue("description"),
		),
	)

	err := h.templates.ExecuteTemplate(w, "form", form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// в форме сразу валидируем в конструкторе
	//
	// в конструктор модели передаем форму. если форма невалидна то возвращаем нил
	//
	// в сервисе если конатруктор модели вернул не нил - сохраняем в репозиторий. иначе просто возвращаем как есть модель и форму
}
