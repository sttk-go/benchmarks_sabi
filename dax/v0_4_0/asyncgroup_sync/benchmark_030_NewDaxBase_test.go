package asyncgroup_sync_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0/asyncgroup_sync"
)

func Benchmark_NewDaxBase(b *testing.B) {
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
