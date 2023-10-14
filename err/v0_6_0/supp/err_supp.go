package supp

import (
	errs "github.com/sttk/benchmarks_sabi/err/v0_6_0"
)

func ReturnOkErr() errs.Err {
	return errs.Ok()
}

type EmptyReason struct{}

func ReturnEmptyReasonedErr() errs.Err {
	return errs.New(EmptyReason{})
}

type OneFieldReason struct {
	FieldA string
}

func ReturnOneFieldReasonedErr() errs.Err {
	return errs.New(OneFieldReason{FieldA: "abc"})
}

func ReturnOneFieldReasonedPtrErr() errs.Err {
	return errs.New(&OneFieldReason{FieldA: "abc"})
}

type FiveFieldReason struct {
	FieldA string
	FieldB int
	FieldC bool
	FieldD string
	FieldE string
}

func ReturnFiveFieldReasonedErr() errs.Err {
	return errs.New(FiveFieldReason{
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

func ReturnHavingCauseReasonedErr() errs.Err {
	return errs.New(HavingCauseReason{}, EmptyError{})
}
