package ag_setup_noerr_sync_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0/ag_setup_noerr_sync"
)

func Benchmark_ShutdownGlobalDaxSrcs_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", &FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
	}
	b.StopTimer()
}

func Benchmark_ShutdownGlobalDaxSrcs_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", &FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.ShutdownGlobalDaxSrcs()
	}
	b.StopTimer()
}

func Benchmark_ShutdownGlobalDaxSrcs_freeDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", &FooDaxSrc{})
	sabi.AddGlobalDaxSrc("database", &FooDaxSrc{})
	sabi.AddGlobalDaxSrc("pubsub", &FooDaxSrc{})
	sabi.AddGlobalDaxSrc("json", &FooDaxSrc{})
	sabi.AddGlobalDaxSrc("env", &FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.ShutdownGlobalDaxSrcs()
	}
	b.StopTimer()
}
