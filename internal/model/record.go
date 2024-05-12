package model

import "time"

type Record struct {
	Task        string
	Description string
	TimeStart   time.Time
	TimeEnd     time.Time
}

func (r *Record) Duration() time.Duration {
	return r.TimeEnd.Sub(r.TimeStart).Round(time.Minute)
}

type Records []Record
