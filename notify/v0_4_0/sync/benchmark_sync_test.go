package sync_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/notify/v0_4_0/sync"
)

func returnOkErr() sabi.Err {
	return sabi.Ok()
}

func Benchmark_AddSyncErrHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sabi.AddSyncErrHandler(func(err sabi.Err, occ sabi.ErrOccasion) {})
	}
}

func Benchmark_ok(b *testing.B) {
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := returnOkErr()
		e = err
	}
	_ = e
}

type EmptyReason struct {
}

func returnEmptyReasonedErr() sabi.Err {
	return sabi.NewErr(EmptyReason{})
}

func Benchmark_emptyReason(b *testing.B) {
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := returnEmptyReasonedErr()
		e = err
	}
	_ = e
}

type OneFieldReason struct {
	FieldA string
}

func returnOneFieldReasonedErr() sabi.Err {
	return sabi.NewErr(OneFieldReason{FieldA: "abc"})
}

func Benchmark_oneFieldReason(b *testing.B) {
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := returnOneFieldReasonedErr()
		e = err
	}
	_ = e
}

type FiveFieldReason struct {
	FieldA string
	FieldB int
	FieldC bool
	FieldD string
	FieldE string
}

func returnFiveFieldReasonedErr() sabi.Err {
	return sabi.NewErr(FiveFieldReason{
		FieldA: "abc", FieldB: 123, FieldC: true, FieldD: "def", FieldE: "ghi",
	})
}

func Benchmark_fiveFieldReason(b *testing.B) {
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := returnFiveFieldReasonedErr()
		e = err
	}
	_ = e
}

type EmptyError struct {
}

func (e EmptyError) Error() string {
	return "EmptyError"
}

type HavingCauseError struct {
	Cause error
}

func (e HavingCauseError) Error() string {
	return "HavingCauseError{cause:" + e.Cause.Error() + "}"
}

func returnHavingCauseReasonedErr() sabi.Err {
	return sabi.NewErr(HavingCauseError{}, EmptyError{})
}

func Benchmark_havingCauseReason(b *testing.B) {
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := returnHavingCauseReasonedErr()
		e = err
	}
	_ = e
}
