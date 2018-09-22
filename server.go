package hrstypes

import (
	"fmt"
	"strconv"
	"strings"
)

/* Status Definition */

// Status DTO
type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

// Interface MetadataObject Implementation
func (r *Status) getObjectInfo() string {
	info := []string{
		strconv.Itoa(r.Code),
		r.Description,
	}
	resp := strings.Join(info, ": ")
	return resp
}

/* HRAResponse Definition */

// HRAResponse DTO
type HRAResponse struct {
	Status  Status         `json:"status"`
	RespObj MetadataObject `json:"respObj"`
	Error   HRSError       `json:"error"`
}

// SetError function
func (fe *HRAResponse) SetError(err HRSError) {
	fe.Error = err
}

// GetError function
func (fe *HRAResponse) GetError() HRSError {
	return fe.Error
}

// HRSError interface
type HRSError interface {
	ShowError() string
}

/** FatalError Error **/

// FatalError Error
type FatalError struct {
	Err error
}

// ShowError - Shows
func (fte *FatalError) ShowError() string {
	return fmt.Sprintf("[ERROR] A fatal error occured: %v", fte.Err)
}

// SetError function
func (fte *FatalError) SetError(err error) {
	fte.Err = err
}

// GetError function
func (fte *FatalError) GetError() error {
	return fte.Err
}

/** FunctionalError Error **/

// FunctionalError Error
type FunctionalError struct {
	Err error
}

// ShowError - Shows
func (fe *FunctionalError) ShowError() string {
	return fmt.Sprintf("[ERROR] A functional error occured: %v", fe.Err)
}

// SetError function
func (fe *FunctionalError) SetError(err error) {
	fe.Err = err
}

// GetError function
func (fe *FunctionalError) GetError() error {
	return fe.Err
}

/** TechnicalError Error **/

// TechnicalError - Struct
type TechnicalError struct {
	Err error
}

// ShowError - Shows
func (te *TechnicalError) ShowError() string {
	return fmt.Sprintf("[ERROR] A functional error occured: %v", te.Err)
}

// SetError function
func (te *TechnicalError) SetError(err error) {
	te.Err = err
}

// GetError function
func (te *TechnicalError) GetError() error {
	return te.Err
}

/** General DTOs **/

// MetadataObject interface
type MetadataObject interface {
	getObjectInfo() string
}
