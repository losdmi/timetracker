package dto

import (
	"github.com/losdmi/timetracker/internal/model"
	"github.com/losdmi/timetracker/internal/model/dto"
)

type IndexData struct {
	Records model.Records
	Form    dto.AddRecordForm
}
