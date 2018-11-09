package hrstypes

import (
	"fmt"
)

// HRSError interface
type HRSError interface {
	ShowError() string
}

/** FatalError Error **/

// FatalError Error
type FatalError struct {
	Err string
}

// ShowError - Shows
func (fte *FatalError) ShowError() string {
	return fmt.Sprintf("A fatal error occured: %s", fte.Err)
}

// SetError function
func (fte *FatalError) SetError(err string) {
	fte.Err = err
}

// GetError function
func (fte *FatalError) GetError() string {
	return fte.Err
}

/** FunctionalError Error **/

// FunctionalError Error
type FunctionalError struct {
	Err string
}

// ShowError - Shows
func (fe *FunctionalError) ShowError() string {
	return fmt.Sprintf("A functional error occured: %s", fe.Err)
}

// SetError function
func (fe *FunctionalError) SetError(err string) {
	fe.Err = err
}

// GetError function
func (fe *FunctionalError) GetError() string {
	return fe.Err
}

/** TechnicalError Error **/

// TechnicalError - Struct
type TechnicalError struct {
	Err string
}

// ShowError - Shows
func (te *TechnicalError) ShowError() string {
	return fmt.Sprintf("A technical error occured: %s", te.Err)
}

// SetError function
func (te *TechnicalError) SetError(err string) {
	te.Err = err
}

// GetError function
func (te *TechnicalError) GetError() string {
	return te.Err
}
