package sabi_err_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/err/sabi_err"
)

func unused(v any) {}

func returnOkErr() sabi.Err {
	return sabi.Ok()
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

func Benchmark_ok_isOk(b *testing.B) {
	var err sabi.Err
	e := returnOkErr()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if e.IsOk() {
			err = e
		}
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_ok_typeSwitch(b *testing.B) {
	var err sabi.Err
	e := returnOkErr()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		switch e.Reason().(type) {
		case nil:
			err = e
		default:
			b.Fail()
		}
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_ok_ErrorString(b *testing.B) {
	var str string
	e := returnOkErr()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s := e.Error()
		str = s
	}
	b.StopTimer()
	unused(str)
}

type EmptyReason struct {
}

func returnEmptyReasonedErr() sabi.Err {
	return sabi.NewErr(EmptyReason{})
}
func Benchmark_empty(b *testing.B) {
	var err sabi.Err
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnEmptyReasonedErr()
		err = e
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_empty_isNotOk(b *testing.B) {
	var err sabi.Err
	e := returnEmptyReasonedErr()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if !e.IsOk() {
			err = e
		}
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_empty_typeSwitch(b *testing.B) {
	var err sabi.Err
	e := returnEmptyReasonedErr()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		switch e.Reason().(type) {
		case EmptyReason:
			err = e
		default:
			b.Fail()
		}
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_empty_ErrorString(b *testing.B) {
	var str string
	e := returnEmptyReasonedErr()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s := e.Error()
		str = s
	}
	b.StopTimer()
	unused(str)
	unused(e)
}

type OneFieldReason struct {
	FieldA string
}

func returnOneFieldReasonedErr() sabi.Err {
	return sabi.NewErr(OneFieldReason{FieldA: "abc"})
}
func Benchmark_oneField(b *testing.B) {
	var err sabi.Err
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnOneFieldReasonedErr()
		err = e
	}
	b.StopTimer()
	unused(err)
}

func returnOneFieldReasonedPtrErr() sabi.Err {
	return sabi.NewErr(&OneFieldReason{FieldA: "abc"})
}
func Benchmark_oneFieldPtr(b *testing.B) {
	var err sabi.Err
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnOneFieldReasonedPtrErr()
		err = e
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_oneField_isNotOk(b *testing.B) {
	var err sabi.Err
	e := returnOneFieldReasonedErr()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if !e.IsOk() {
			err = e
		}
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_oneField_typeSwitch(b *testing.B) {
	var err sabi.Err
	e := returnOneFieldReasonedErr()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		switch e.Reason().(type) {
		case OneFieldReason:
			err = e
		default:
			b.Fail()
		}
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_oneField_ErrorString(b *testing.B) {
	var str string
	e := returnOneFieldReasonedErr()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s := e.Error()
		str = s
	}
	b.StopTimer()
	unused(str)
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
func Benchmark_fiveField(b *testing.B) {
	var err sabi.Err
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnFiveFieldReasonedErr()
		err = e
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_fiveField_isNotOk(b *testing.B) {
	var err sabi.Err
	e := returnFiveFieldReasonedErr()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if !e.IsOk() {
			err = e
		}
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_fiveField_typeSwitch(b *testing.B) {
	var err sabi.Err
	e := returnFiveFieldReasonedErr()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		switch e.Reason().(type) {
		case FiveFieldReason:
			err = e
		default:
			b.Fail()
		}
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_fiveField_ErrorString(b *testing.B) {
	var str string
	e := returnFiveFieldReasonedErr()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s := e.Error()
		str = s
	}
	b.StopTimer()
	unused(str)
}

type HavingCauseReason struct {
}

type EmptyError struct {
}

func (e EmptyError) Error() string {
	return "EmptyError"
}

func returnHavingCauseReasonedErr() sabi.Err {
	return sabi.NewErr(HavingCauseReason{}, EmptyError{})
}

func Benchmark_havingCause(b *testing.B) {
	var err sabi.Err
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnHavingCauseReasonedErr()
		err = e
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_havingCause_ErrorString(b *testing.B) {
	var str string
	e := returnHavingCauseReasonedErr()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s := e.Error()
		str = s
	}
	b.StopTimer()
	unused(str)
}
