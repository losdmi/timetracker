package service

import (
	"context"
	"fmt"
	"time"

	"github.com/losdmi/timetracker/internal/model"
	modelDTO "github.com/losdmi/timetracker/internal/model/dto"
	"github.com/losdmi/timetracker/internal/service/dto"
)

func (s *Service) AddRecord(ctx context.Context, form modelDTO.AddRecordForm) dto.AddRecordResponse {
	record, form := model.NewRecord(form)
	if record == nil {
		return dto.AddRecordResponse{
			Form: form,
		}
	}

	err := s.repository.CreateRecord(ctx, *record)
	if err != nil {
		fmt.Printf("Service.AddRecord ошибка при сохранении записи: %s", err)
	}

	records, err := s.repository.ReadDayRecords(ctx, time.Now())
	if err != nil {
		fmt.Printf("Service.AddRecord ошибка при чтении записей: %s", err)
	}

	return dto.AddRecordResponse{
		Records: records,
		Form:    form,
	}
}
