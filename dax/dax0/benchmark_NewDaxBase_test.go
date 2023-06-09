package dax0_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/dax0"
	"testing"
)

func Benchmark_NewDaxBase(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base := sabi.NewDaxBase()
		_ = base
	}
	b.StopTimer()
}
