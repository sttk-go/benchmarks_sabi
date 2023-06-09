package dax3_ready_async_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/dax3_ready_async"
	"testing"
)

func Benchmark_AddGlobalDaxSrc_oneDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddGlobalDaxSrc("foo", FooDaxSrc{})
	}
	b.StopTimer()
}

func Benchmark_AddGlobalDaxSrc_fiveDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddGlobalDaxSrc("foo", FooDaxSrc{})
		sabi.AddGlobalDaxSrc("bar", BarDaxSrc{})
		sabi.AddGlobalDaxSrc("baz", BazDaxSrc{})
		sabi.AddGlobalDaxSrc("qux", QuxDaxSrc{})
		sabi.AddGlobalDaxSrc("corge", CorgeDaxSrc{})
	}
	b.StopTimer()
}
