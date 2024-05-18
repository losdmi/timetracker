package service

import (
	"context"
	"fmt"

	"github.com/losdmi/timetracker/internal/model"
	"github.com/losdmi/timetracker/internal/model/dto"
)

func (s *Service) AddRecord(ctx context.Context, form dto.AddRecordForm) (*model.Record, dto.AddRecordForm) {
	record, form := model.NewRecord(form)

	if record != nil {
		err := s.repository.CreateRecord(ctx, *record)
		if err != nil {
			fmt.Println(err)
		}
	}

	return record, form
}
