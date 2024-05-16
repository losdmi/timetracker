package service

import "github.com/losdmi/timetracker/internal/model"

func (s *Service) AddRecord(form model.AddRecordForm) (*model.Record, model.AddRecordForm) {
	record, form := model.NewRecord(form)

	return record, form
}
