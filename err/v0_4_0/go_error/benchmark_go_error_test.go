package go_error_test

import (
	"testing"

	"github.com/sttk/benchmarks_sabi/err/v0_4_0/go_error"
)

func returnNilError() error {
	return nil
}

func Benchmark_ok(b *testing.B) {
	var e error
	for i := 0; i < b.N; i++ {
		err := returnNilError()
		e = err
	}
	_ = e
}

func Benchmark_ok_isOk(b *testing.B) {
	b.StopTimer()
	err := returnNilError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		if err == nil {
			e = err
		}
	}
	_ = e
}

func Benchmark_ok_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := returnNilError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		switch err.(type) {
		case nil:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func Benchmark_ok_ErrorString(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		str := "nil"
		s = str
	}
	_ = s
}

func returnEmptyError() error {
	return go_error.EmptyError{}
}

func Benchmark_empty(b *testing.B) {
	var e error
	for i := 0; i < b.N; i++ {
		err := returnEmptyError()
		e = err
	}
	_ = e
}

func Benchmark_empty_isNotOk(b *testing.B) {
	b.StopTimer()
	err := returnEmptyError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		if err != nil {
			e = err
		}
	}
	_ = e
}

func Benchmark_empty_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := returnEmptyError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		switch err.(type) {
		case go_error.EmptyError:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func Benchmark_empty_ErrorString(b *testing.B) {
	b.StopTimer()
	err := returnEmptyError()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}

func returnOneFieldError() error {
	return go_error.OneFieldError{FieldA: "abc"}
}

func Benchmark_oneField(b *testing.B) {
	var e error
	for i := 0; i < b.N; i++ {
		err := returnOneFieldError()
		e = err
	}
	_ = e
}

func returnOneFieldErrorPtr() error {
	return &go_error.OneFieldError{FieldA: "abc"}
}

func Benchmark_oneFieldPtr(b *testing.B) {
	var e error
	for i := 0; i < b.N; i++ {
		err := returnOneFieldErrorPtr()
		e = err
	}
	_ = e
}

func Benchmark_oneField_isNotOk(b *testing.B) {
	b.StopTimer()
	err := returnOneFieldError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		if err != nil {
			e = err
		}
	}
	_ = e
}

func Benchmark_oneField_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := returnOneFieldError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		switch err.(type) {
		case go_error.OneFieldError:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func Benchmark_oneField_ErrorString(b *testing.B) {
	b.StopTimer()
	err := returnOneFieldError()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}

func returnFiveFieldError() error {
	return go_error.FiveFieldError{
		FieldA: "abc", FieldB: 123, FieldC: true, FieldD: "def", FieldE: "ghi",
	}
}

func Benchmark_fiveField(b *testing.B) {
	var e error
	for i := 0; i < b.N; i++ {
		err := returnFiveFieldError()
		e = err
	}
	_ = e
}

func Benchmark_fiveField_isNotOk(b *testing.B) {
	b.StopTimer()
	err := returnFiveFieldError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		if err != nil {
			e = err
		}
	}
	_ = e
}

func Benchmark_fiveField_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := returnFiveFieldError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		switch err.(type) {
		case go_error.FiveFieldError:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func Benchmark_fiveField_ErrorString(b *testing.B) {
	b.StopTimer()
	err := returnFiveFieldError()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}

func returnHavingCauseError() error {
	return go_error.HavingCauseError{Cause: go_error.EmptyError{}}
}

func Benchmark_havingCause(b *testing.B) {
	var e error
	for i := 0; i < b.N; i++ {
		err := returnHavingCauseError()
		e = err
	}
	_ = e
}

func Benchmark_havingCause_ErrorString(b *testing.B) {
	b.StopTimer()
	err := returnHavingCauseError()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}
