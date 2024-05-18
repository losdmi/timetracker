package service

import (
	"context"
	"fmt"

	"github.com/losdmi/timetracker/internal/model"
	"github.com/losdmi/timetracker/internal/model/dto"
)

func (s *Service) AddRecord(ctx context.Context, form dto.AddRecordForm) (*model.Record, dto.AddRecordForm) {
	record, form := model.NewRecord(form)
	if record == nil {
		return nil, form
	}

	err := s.repository.CreateRecord(ctx, *record)
	if err != nil {
		fmt.Printf("Service.AddRecord ошибка: %s", err)
	}

	return record, form
}
