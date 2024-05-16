package model

import (
	"fmt"
	"strings"
)

type AddRecordFormErrors struct {
	Task        error
	Description error
}

func (fe AddRecordFormErrors) isEmpty() bool {
	return fe.Task == nil || fe.Description == nil
}

type AddRecordForm struct {
	Task        string
	Description string

	Errors AddRecordFormErrors
}

func NewAddRecordForm(task string, description string) AddRecordForm {
	form := AddRecordForm{
		Task:        task,
		Description: description,
	}

	return form
}

func (f AddRecordForm) validate() AddRecordForm {
	if strings.TrimSpace(f.Task) == "" {
		f.Errors.Task = fmt.Errorf("task must be non-empty")
	}
	if strings.TrimSpace(f.Description) == "" {
		f.Errors.Description = fmt.Errorf("description must be non-empty")
	}

	return f
}

func (f AddRecordForm) IsValid() bool {
	return f.Errors.isEmpty()
}
