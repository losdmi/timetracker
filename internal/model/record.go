package model

import (
	"fmt"
	"time"

	"github.com/losdmi/timetracker/internal/model/dto"
)

type Record struct {
	ID          uint64
	Task        string
	Description string
	TimeStart   time.Time
	TimeEnd     *time.Time
}

type Records []Record

func NewRecord(form dto.AddRecordForm) (*Record, dto.AddRecordForm) {
	form = validateAddRecordForm(form)
	if !form.IsValid() {
		return nil, form
	}

	return &Record{
		Task:        form.Task,
		Description: form.Description,
		TimeStart:   time.Now().UTC(),
	}, dto.NewAddRecordForm("", "")
}

func validateAddRecordForm(f dto.AddRecordForm) dto.AddRecordForm {
	if f.Task == "" {
		f.Errors.Task = fmt.Errorf("task must be non-empty")
	}
	if f.Description == "" {
		f.Errors.Description = fmt.Errorf("description must be non-empty")
	}

	return f
}

func (r *Record) DurationString() string {
	timeEnd := time.Now()
	if r.TimeEnd != nil {
		timeEnd = *r.TimeEnd
	}

	duration := timeEnd.Sub(r.TimeStart).Round(time.Minute)
	hours := int(duration.Minutes()) / 60
	minutes := int(duration.Minutes()) % 60

	return fmt.Sprintf("%d ч %d мин", hours, minutes)
}
