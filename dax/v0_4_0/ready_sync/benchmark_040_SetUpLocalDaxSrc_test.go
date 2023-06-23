package ready_sync_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0/ready_sync"
)

func Benchmark_SetUpLocalDaxSrc_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.FreeLocalDaxSrcForTest(base, "cliargs")
	}
	b.StopTimer()
}

func Benchmark_SetUpLocalDaxSrc_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})
		sabi.FreeLocalDaxSrcForTest(base, "cliargs")
	}
	b.StopTimer()
}

func Benchmark_SetUpLocalDaxSrc_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})
		base.SetUpLocalDaxSrc("database", FooDaxSrc{})
		base.SetUpLocalDaxSrc("pubsub", FooDaxSrc{})
		base.SetUpLocalDaxSrc("json", FooDaxSrc{})
		base.SetUpLocalDaxSrc("env", FooDaxSrc{})

		sabi.FreeLocalDaxSrcForTest(base, "cliargs")
		sabi.FreeLocalDaxSrcForTest(base, "database")
		sabi.FreeLocalDaxSrcForTest(base, "pubsub")
		sabi.FreeLocalDaxSrcForTest(base, "json")
		sabi.FreeLocalDaxSrcForTest(base, "env")
	}
	b.StopTimer()
}
