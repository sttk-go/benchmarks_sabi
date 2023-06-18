package async_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/notify/async"
)

func unused(v any) {}

func returnOkErr() sabi.Err {
	return sabi.Ok()
}

func Benchmark_AddAsyncErrHandler(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddAsyncErrHandler(func(err sabi.Err, occ sabi.ErrOccasion) {})
		sabi.FixErrCfgs()
	}
	b.StopTimer()
}

func Benchmark_ok(b *testing.B) {
	var err sabi.Err
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnOkErr()
		err = e
	}
	b.StopTimer()
	unused(err)
}

type EmptyReason struct {
}

func returnEmptyReasonedErr() sabi.Err {
	return sabi.NewErr(EmptyReason{})
}

func Benchmark_emptyReason(b *testing.B) {
	var err sabi.Err
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnEmptyReasonedErr()
		err = e
	}
	b.StopTimer()
	unused(err)
}

type OneFieldReason struct {
	FieldA string
}

func returnOneFieldReasonedErr() sabi.Err {
	return sabi.NewErr(OneFieldReason{FieldA: "abc"})
}

func Benchmark_oneFieldReason(b *testing.B) {
	var err sabi.Err
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnOneFieldReasonedErr()
		err = e
	}
	b.StopTimer()
	unused(err)
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
	var err sabi.Err
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnFiveFieldReasonedErr()
		err = e
	}
	b.StopTimer()
	unused(err)
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
	var err sabi.Err
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnHavingCauseReasonedErr()
		err = e
	}
	b.StopTimer()
	unused(err)
}
