package orderedmap_isasync_async_test

import (
	"sync"
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_isasync_async"
)

func Benchmark_NewDaxBase(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	var wg sync.WaitGroup

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base := sabi.NewDaxBase(&wg)
		_ = base
	}
	b.StopTimer()
}
