package service

import (
	"context"
	"fmt"
	"time"

	modelDTO "github.com/losdmi/timetracker/internal/model/dto"
	serviceDTO "github.com/losdmi/timetracker/internal/service/dto"
)

func (s *Service) GetIndexData(ctx context.Context) *serviceDTO.IndexData {
	records, err := s.repository.ReadDayRecords(ctx, time.Now().Local())
	if err != nil {
		fmt.Printf("Service.GetIndexData ошибка: %s", err)
	}

	return &serviceDTO.IndexData{
		Records: records,
		// model.Records{
		// {Task: "FSP-1111", Description: "qweasdzxc", TimeStart: time.Date(2024, 5, 11, 21, 33, 0, 0, time.UTC), TimeEnd: &timeEnd},
		// },
		Form: modelDTO.NewAddRecordForm("", ""),
	}
}
