package service

import (
	"time"

	"github.com/losdmi/timetracker/internal/model"
	modelDTO "github.com/losdmi/timetracker/internal/model/dto"
	serviceDTO "github.com/losdmi/timetracker/internal/service/dto"
)

func (s *Service) GetIndexData() *serviceDTO.IndexData {
	timeEnd := time.Date(2024, 5, 11, 21, 47, 0, 0, time.UTC)
	return &serviceDTO.IndexData{
		Records: model.Records{
			{Task: "FSP-1111", Description: "qweasdzxc", TimeStart: time.Date(2024, 5, 11, 21, 33, 0, 0, time.UTC), TimeEnd: &timeEnd},
		},
		Form: modelDTO.NewAddRecordForm("", ""),
	}
}
