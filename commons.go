package hrstypes

import (
	"strconv"
	"strings"
)

// DTOObject interface
type DTOObject interface {
	GetObjectInfo() string
}

/* Status Definition */

// Status DTO
type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

// GetObjectInfo - Interface DTOObject Implementation
func (r *Status) GetObjectInfo() string {
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
	Status  Status    `json:"status"`
	RespObj DTOObject `json:"respObj"`
	Error   HRSError  `json:"error"`
}

// SetError function
func (fe *HRAResponse) SetError(err HRSError) {
	fe.Error = err
}

// GetError function
func (fe *HRAResponse) GetError() HRSError {
	return fe.Error
}
