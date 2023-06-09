package dax2_ready_sync_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/dax2_ready_sync"
	"testing"
)

func Benchmark_ShutdownGlobalDaxSrcs_oneDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("foo", FooDaxSrc{})
	sabi.StartUpGlobalDaxSrcs()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.ShutdownGlobalDaxSrcs()
	}
	b.StopTimer()
}

func Benchmark_ShutdownGlobalDaxSrcs_fiveDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("foo", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("bar", BarDaxSrc{})
	sabi.AddGlobalDaxSrc("baz", BazDaxSrc{})
	sabi.AddGlobalDaxSrc("qux", QuxDaxSrc{})
	sabi.AddGlobalDaxSrc("corge", CorgeDaxSrc{})
	sabi.StartUpGlobalDaxSrcs()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.ShutdownGlobalDaxSrcs()
	}
	b.StopTimer()
}
