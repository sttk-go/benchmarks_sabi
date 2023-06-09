package dax1_noLock_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/dax1_noLock"
	"testing"
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
