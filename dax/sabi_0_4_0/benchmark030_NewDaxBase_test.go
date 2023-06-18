package sabi_0_4_0_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/sabi_0_4_0"
)

func Benchmark_NewDaxBase(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base := sabi.NewDaxBase()
		_ = base
	}
	b.StopTimer()
}
