package v0_6_0_test

import (
	"testing"

	prev "github.com/sttk/benchmarks_sabi/dax/v0_5_0"
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_6_0"
)

func BenchmarkDax_____Sabi_NewDaxBase(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base := sabi.NewDaxBase()
		_ = base
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_NewDaxBase(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base := prev.NewDaxBase()
		_ = base
	}
	b.StopTimer()
}
