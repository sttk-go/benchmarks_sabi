package supp

import (
	sabi "github.com/sttk/benchmarks_sabi/err/v0_4_0"
)

func ReturnOkErr() sabi.Err {
	return sabi.Ok()
}

type EmptyReason struct{}

func ReturnEmptyReasonedErr() sabi.Err {
	return sabi.NewErr(EmptyReason{})
}

type OneFieldReason struct {
	FieldA string
}

func ReturnOneFieldReasonedErr() sabi.Err {
	return sabi.NewErr(OneFieldReason{FieldA: "abc"})
}

func ReturnOneFieldReasonedPtrErr() sabi.Err {
	return sabi.NewErr(&OneFieldReason{FieldA: "abc"})
}

type FiveFieldReason struct {
	FieldA string
	FieldB int
	FieldC bool
	FieldD string
	FieldE string
}

func ReturnFiveFieldReasonedErr() sabi.Err {
	return sabi.NewErr(FiveFieldReason{
		FieldA: "abc", FieldB: 123, FieldC: true, FieldD: "def", FieldE: "ghi",
	})
}

type HavingCauseReason struct {
}

type EmptyError struct {
}

func (e EmptyError) Error() string {
	return "EmptyError"
}

func ReturnHavingCauseReasonedErr() sabi.Err {
	return sabi.NewErr(HavingCauseReason{}, EmptyError{})
}
