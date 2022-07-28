package errors

import (
	"errors"
	"fmt"
)

// Not mutant error
type NotMutant struct {
	Err error
}

func (e *NotMutant) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewNotMutant() *NotMutant {
	defaultMsg := "not mutant"
	return &NotMutant{Err: errors.New(defaultMsg)}
}

// ------------------------------------------------------------

// Duplicate dna entry on database error
type DuplicateDna struct {
	Err error
}

func (e *DuplicateDna) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewDuplicateDna(isMutant bool) *DuplicateDna {
	defaultMsg := "duplicate sample dna"
	if isMutant {
		return &DuplicateDna{Err: errors.New(defaultMsg + ": mutant")}
	}
	return &DuplicateDna{Err: errors.New(defaultMsg + ": not mutant")}
}

// ------------------------------------------------------------

// Method not allowed
type MethodNotAllowed struct {
	Err error
}

func (e *MethodNotAllowed) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewMethodNotAllowed() *MethodNotAllowed {
	defaultMsg := "method not allowed"
	return &MethodNotAllowed{Err: errors.New(defaultMsg)}
}

// ------------------------------------------------------------

// Bad request
type BadRequest struct {
	Err error
}

func (e *BadRequest) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewBadRequest(msg ...string) *BadRequest {
	defaultMsg := "bad request"
	if len(msg) == 0 {
		return &BadRequest{Err: errors.New(defaultMsg)}
	}
	return &BadRequest{Err: errors.New(msg[0])}
}

// ------------------------------------------------------------
