package dax2_ready_sync_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/dax2_ready_sync"
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
