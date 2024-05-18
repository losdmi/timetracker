package dto

import (
	"github.com/losdmi/timetracker/internal/model"
	"github.com/losdmi/timetracker/internal/model/dto"
)

type AddRecordResponse struct {
	Records model.Records
	Form    dto.AddRecordForm
}
