package model

import (
	"fmt"
	"time"

	"github.com/losdmi/timetracker/internal/model/dto"
)

type Record struct {
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
	}, form
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

func (r *Record) Duration() time.Duration {
	timeEnd := time.Now().UTC()
	if r.TimeEnd != nil {
		timeEnd = *r.TimeEnd
	}

	return timeEnd.Sub(r.TimeStart).Round(time.Minute)
}
