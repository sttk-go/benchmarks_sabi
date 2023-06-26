package orderedmap_isasync_async_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_isasync_async"
)

func Benchmark_StartUpGlobalDaxSrcs_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
	}
	b.StopTimer()
}

func Benchmark_StartUpGlobalDaxSrcs_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.StartUpGlobalDaxSrcs()
	}
	b.StopTimer()
}

func Benchmark_StartUpGlobalDaxSrcs_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("database", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("pubsub", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("json", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("env", FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.StartUpGlobalDaxSrcs()
	}
	b.StopTimer()
}
