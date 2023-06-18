package unlock_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/sabi_0_4_0/unlock"
)

func Benchmark_SetupGlobalDaxSrcs_oneDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrcForTest("cliargs", FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.StartUpGlobalDaxSrcs()
	}
	b.StopTimer()
}

func Benchmark_SetupGlobalDaxSrcs_fiveDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrcForTest("cliargs", FooDaxSrc{})
	sabi.AddGlobalDaxSrcForTest("database", FooDaxSrc{})
	sabi.AddGlobalDaxSrcForTest("pubsub", FooDaxSrc{})
	sabi.AddGlobalDaxSrcForTest("json", FooDaxSrc{})
	sabi.AddGlobalDaxSrcForTest("env", FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.StartUpGlobalDaxSrcs()
	}
	b.StopTimer()
}
