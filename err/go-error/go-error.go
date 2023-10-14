package go_error

import (
	"strconv"
)

func ReturnNilError() error {
	return nil
}

func ReturnEmptyError() error {
	return EmptyError{}
}

type EmptyError struct {
}

func (e EmptyError) Error() string {
	return "EmptyError"
}

func ReturnOneFieldError() error {
	return OneFieldError{FieldA: "abc"}
}

func ReturnOneFieldErrorPtr() error {
	return &OneFieldError{FieldA: "abc"}
}

type OneFieldError struct {
	FieldA string
}

func (e OneFieldError) Error() string {
	return "OneFieldError{FieldA:" + e.FieldA + "}"
}

func ReturnFiveFieldError() error {
	return FiveFieldError{
		FieldA: "abc", FieldB: 123, FieldC: true, FieldD: "def", FieldE: "ghi",
	}
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

func ReturnHavingCauseError() error {
	return HavingCauseError{Cause: EmptyError{}}
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
