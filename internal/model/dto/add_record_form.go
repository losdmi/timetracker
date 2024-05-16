package dto

import "strings"

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
		Task:        strings.TrimSpace(task),
		Description: strings.TrimSpace(description),
	}

	return form
}

func (f AddRecordForm) IsValid() bool {
	return f.Errors.isEmpty()
}
