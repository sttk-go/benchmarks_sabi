package v0_4_0_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_4_0/supp"
)

func Benchmark_StartUpGlobalDaxSrcs_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", supp.FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
	}
	b.StopTimer()
}

func Benchmark_StartUpGlobalDaxSrcs_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", supp.FooDaxSrc{})

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

	sabi.AddGlobalDaxSrc("cliargs", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("database", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("pubsub", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("json", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("env", supp.FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.StartUpGlobalDaxSrcs()
	}
	b.StopTimer()
}
