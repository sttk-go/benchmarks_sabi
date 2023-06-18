package go_error_test

import (
	"testing"

	"github.com/sttk/benchmarks_sabi/err/go_error"
)

func unused(v any) {}

func returnNilError() error {
	return nil
}

func Benchmark_ok(b *testing.B) {
	var err error
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnNilError()
		err = e
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_ok_isOk(b *testing.B) {
	var err error
	e := returnNilError()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if e == nil {
			err = e
		}
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_ok_typeSwitch(b *testing.B) {
	var err error
	e := returnNilError()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		switch e.(type) {
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
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s := "nil"
		str = s
	}
	b.StopTimer()
	unused(str)
}

func returnEmptyError() error {
	return go_error.EmptyError{}
}

func Benchmark_empty(b *testing.B) {
	var err error
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnEmptyError()
		err = e
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_empty_isOk(b *testing.B) {
	var err error
	e := returnEmptyError()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if e != nil {
			err = e
		}
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_empty_typeSwitch(b *testing.B) {
	var err error
	e := returnEmptyError()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		switch e.(type) {
		case go_error.EmptyError:
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
	e := returnEmptyError()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s := e.Error()
		str = s
	}
	b.StopTimer()
	unused(str)
}

func returnOneFieldError() error {
	return go_error.OneFieldError{FieldA: "abc"}
}

func Benchmark_oneField(b *testing.B) {
	var err error
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnOneFieldError()
		err = e
	}
	b.StopTimer()
	unused(err)
}

func returnOneFieldErrorPtr() error {
	return &go_error.OneFieldError{FieldA: "abc"}
}

func Benchmark_oneFieldPtr(b *testing.B) {
	var err error
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnOneFieldErrorPtr()
		err = e
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_oneField_isOk(b *testing.B) {
	var err error
	e := returnOneFieldError()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if e != nil {
			err = e
		}
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_oneField_typeSwitch(b *testing.B) {
	var err error
	e := returnOneFieldError()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		switch e.(type) {
		case go_error.OneFieldError:
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
	e := returnOneFieldError()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s := e.Error()
		str = s
	}
	b.StopTimer()
	unused(str)
}

func returnFiveFieldError() error {
	return go_error.FiveFieldError{
		FieldA: "abc", FieldB: 123, FieldC: true, FieldD: "def", FieldE: "ghi",
	}
}

func Benchmark_fiveField(b *testing.B) {
	var err error
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnFiveFieldError()
		err = e
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_fiveField_isOk(b *testing.B) {
	var err error
	e := returnFiveFieldError()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if e != nil {
			err = e
		}
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_fiveField_typeSwitch(b *testing.B) {
	var err error
	e := returnFiveFieldError()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		switch e.(type) {
		case go_error.FiveFieldError:
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
	e := returnFiveFieldError()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s := e.Error()
		str = s
	}
	b.StopTimer()
	unused(str)
}

func returnHavingCauseError() error {
	return go_error.HavingCauseError{Cause: go_error.EmptyError{}}
}

func Benchmark_havingCause(b *testing.B) {
	var err error
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := returnHavingCauseError()
		err = e
	}
	b.StopTimer()
	unused(err)
}

func Benchmark_havingCause_ErrorString(b *testing.B) {
	var str string
	e := returnHavingCauseError()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s := e.Error()
		str = s
	}
	b.StopTimer()
	unused(str)
}
