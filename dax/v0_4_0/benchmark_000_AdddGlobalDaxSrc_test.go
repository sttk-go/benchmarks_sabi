package v0_4_0_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_4_0/supp"
)

func Benchmark_AddGlobalDaxSrc_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.ResetGlobals()
	}
	b.StopTimer()
}

func Benchmark_AddGlobalDaxSrc_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddGlobalDaxSrc("cliargs", supp.FooDaxSrc{})
		sabi.ResetGlobals()
	}
	b.StopTimer()
}

func Benchmark_AddGlobalDaxSrc_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddGlobalDaxSrc("cliargs", supp.FooDaxSrc{})
		sabi.AddGlobalDaxSrc("database", supp.FooDaxSrc{})
		sabi.AddGlobalDaxSrc("pubsub", supp.FooDaxSrc{})
		sabi.AddGlobalDaxSrc("json", supp.FooDaxSrc{})
		sabi.AddGlobalDaxSrc("env", supp.FooDaxSrc{})

		sabi.ResetGlobals()
	}
	b.StopTimer()
}
