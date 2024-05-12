package service

import (
	"time"

	"github.com/losdmi/timetracker/internal/model"
	"github.com/losdmi/timetracker/internal/service/dto"
)

func (s *Service) GetIndexData() *dto.IndexData {
	return &dto.IndexData{
		Records: model.Records{
			{Task: "FSP-1111", Description: "qweasdzxc", TimeStart: time.Date(2024, 5, 11, 21, 33, 0, 0, time.UTC), TimeEnd: time.Date(2024, 5, 11, 21, 47, 0, 0, time.UTC)},
		},
	}
}
