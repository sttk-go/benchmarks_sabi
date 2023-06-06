package go_error

import (
  "strconv"
)

func returnNilError() error {
  return nil
}

type EmptyError struct {
}

func (e EmptyError) Error() string {
  return "EmptyError"
}

type OneFieldError struct {
  FieldA string
}

func (e OneFieldError) Error() string {
  return "OneFieldError{FieldA:" + e.FieldA + "}"
}

type FiveFieldError struct {
	FieldA string
	FieldB int
	FieldC bool
	FieldD string
	FieldE string
}

func (e FiveFieldError) Error() string {
	return "FiveFieldError{FieldA:" + e.FieldA +
		",FieldB:" + strconv.Itoa(e.FieldB) +
		",FieldC:" + strconv.FormatBool(e.FieldC) +
		",FieldD:" + e.FieldD + ",FieldE:" + e.FieldE +
		"}"
}

type HavingCauseError struct {
	Cause error
}

func (e HavingCauseError) Error() string {
	return "HavingCauseError{cause:" + e.Cause.Error() + "}"
}

func (e HavingCauseError) Unwrap() error {
	return e.Cause
}
