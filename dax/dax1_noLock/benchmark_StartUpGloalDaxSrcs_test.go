package dax1_noLock_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/dax1_noLock"
	"testing"
)

func Benchmark_StartUpGlobalDaxSrcs_oneDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("foo", FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.StartUpGlobalDaxSrcs()
	}
	b.StopTimer()
}

func Benchmark_StartUpGlobalDaxSrcs_fiveDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("foo", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("bar", BarDaxSrc{})
	sabi.AddGlobalDaxSrc("baz", BazDaxSrc{})
	sabi.AddGlobalDaxSrc("qux", QuxDaxSrc{})
	sabi.AddGlobalDaxSrc("corge", CorgeDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.StartUpGlobalDaxSrcs()
	}
	b.StopTimer()
}
