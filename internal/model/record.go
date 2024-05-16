package model

import "time"

type Record struct {
	Task        string
	Description string
	TimeStart   time.Time
	TimeEnd     *time.Time
}

type Records []Record

func NewRecord(form AddRecordForm) (*Record, AddRecordForm) {
	form = form.validate()
	if form.IsValid() {
		return &Record{
			Task:        form.Task,
			Description: form.Description,
		}, form
	}

	return nil, form
}

func (r *Record) Duration() time.Duration {
	timeEnd := time.Now().UTC()
	if r.TimeEnd != nil {
		timeEnd = *r.TimeEnd
	}

	return timeEnd.Sub(r.TimeStart).Round(time.Minute)
}
