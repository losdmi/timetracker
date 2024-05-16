package service

import (
	"github.com/losdmi/timetracker/internal/model"
	"github.com/losdmi/timetracker/internal/model/dto"
)

func (s *Service) AddRecord(form dto.AddRecordForm) (*model.Record, dto.AddRecordForm) {
	record, form := model.NewRecord(form)

	return record, form
}
