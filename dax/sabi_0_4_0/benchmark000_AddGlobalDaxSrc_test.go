package sabi_0_4_0_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/sabi_0_4_0"
)

func Benchmark_AddGlobalDaxSrc_oneDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})
	}
	b.StopTimer()
}

func Benchmark_AddGlobalDaxSrc_fiveDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})
		sabi.AddGlobalDaxSrc("database", FooDaxSrc{})
		sabi.AddGlobalDaxSrc("pubsub", FooDaxSrc{})
		sabi.AddGlobalDaxSrc("json", FooDaxSrc{})
		sabi.AddGlobalDaxSrc("env", FooDaxSrc{})
	}
	b.StopTimer()
}
