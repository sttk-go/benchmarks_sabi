package v0_4_0_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/err/v0_4_0"
)

func returnOkErr() sabi.Err {
	return sabi.Ok()
}
func Benchmark_ok(b *testing.B) {
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := returnOkErr()
		e = err
	}
	_ = e
}

func Benchmark_ok_isOk(b *testing.B) {
	b.StopTimer()
	err := returnOkErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		if err.IsOk() {
			e = err
		}
	}
	_ = e
}

func Benchmark_ok_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := returnOkErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		switch err.Reason().(type) {
		case nil:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func Benchmark_ok_ErrorString(b *testing.B) {
	b.StopTimer()
	err := returnOkErr()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}

type EmptyReason struct {
}

func returnEmptyReasonedErr() sabi.Err {
	return sabi.NewErr(EmptyReason{})
}

func Benchmark_empty(b *testing.B) {
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := returnEmptyReasonedErr()
		e = err
	}
	_ = e
}

func Benchmark_empty_isNotOk(b *testing.B) {
	b.StopTimer()
	err := returnEmptyReasonedErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		if !err.IsOk() {
			e = err
		}
	}
	_ = e
}

func Benchmark_empty_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := returnEmptyReasonedErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		switch err.Reason().(type) {
		case EmptyReason:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func Benchmark_empty_ErrorString(b *testing.B) {
	b.StopTimer()
	err := returnEmptyReasonedErr()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}

type OneFieldReason struct {
	FieldA string
}

func returnOneFieldReasonedErr() sabi.Err {
	return sabi.NewErr(OneFieldReason{FieldA: "abc"})
}

func Benchmark_oneField(b *testing.B) {
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := returnOneFieldReasonedErr()
		e = err
	}
	_ = e
}

func returnOneFieldReasonedPtrErr() sabi.Err {
	return sabi.NewErr(&OneFieldReason{FieldA: "abc"})
}

func Benchmark_oneFieldPtr(b *testing.B) {
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := returnOneFieldReasonedPtrErr()
		e = err
	}
	_ = e
}

func Benchmark_oneField_isNotOk(b *testing.B) {
	b.StopTimer()
	err := returnOneFieldReasonedErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		if !err.IsOk() {
			e = err
		}
	}
	_ = e
}

func Benchmark_oneField_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := returnOneFieldReasonedErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		switch err.Reason().(type) {
		case OneFieldReason:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func Benchmark_oneField_ErrorString(b *testing.B) {
	b.StopTimer()
	err := returnOneFieldReasonedErr()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
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
	for i := 0; i < b.N; i++ {
		err := returnFiveFieldReasonedErr()
		_ = err
	}
}

func Benchmark_fiveField_isNotOk(b *testing.B) {
	b.StopTimer()
	err := returnFiveFieldReasonedErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		if !err.IsOk() {
			e = err
		}
	}
	_ = e
}

func Benchmark_fiveField_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := returnFiveFieldReasonedErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		switch err.Reason().(type) {
		case FiveFieldReason:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func Benchmark_fiveField_ErrorString(b *testing.B) {
	b.StopTimer()
	err := returnFiveFieldReasonedErr()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
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
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := returnHavingCauseReasonedErr()
		e = err
	}
	_ = e
}

func Benchmark_havingCause_ErrorString(b *testing.B) {
	b.StopTimer()
	err := returnHavingCauseReasonedErr()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}
