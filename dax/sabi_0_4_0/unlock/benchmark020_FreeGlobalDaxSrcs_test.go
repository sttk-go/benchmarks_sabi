package unlock_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/unlock"
	"testing"
)

func Benchmark_FreeGlobalDaxSrcs_oneDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddGlobalDaxSrcForTest("cliargs", FooDaxSrc{})
		sabi.ShutdownGlobalDaxSrcs()
	}
	b.StopTimer()
}

func Benchmark_FreeGlobalDaxSrcs_fiveDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddGlobalDaxSrcForTest("cliargs", FooDaxSrc{})
		sabi.AddGlobalDaxSrcForTest("database", FooDaxSrc{})
		sabi.AddGlobalDaxSrcForTest("pubsub", FooDaxSrc{})
		sabi.AddGlobalDaxSrcForTest("json", FooDaxSrc{})
		sabi.AddGlobalDaxSrcForTest("env", FooDaxSrc{})

		sabi.ShutdownGlobalDaxSrcs()
	}
	b.StopTimer()
}
