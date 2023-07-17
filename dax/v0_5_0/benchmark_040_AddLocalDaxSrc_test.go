package v0_5_0_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_5_0"
)

func Benchmark_AddLocalDaxSrc_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.FreeAllLocalDaxSrcsForTest(base)
	}
	b.StopTimer()
}

func Benchmark_AddLocalDaxSrc_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.AddLocalDaxSrc("cliargs", FooDaxSrc{})
		sabi.FreeAllLocalDaxSrcsForTest(base)
	}
	b.StopTimer()
}

func Benchmark_AddLocalDaxSrc_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.AddLocalDaxSrc("cliargs", FooDaxSrc{})
		base.AddLocalDaxSrc("database", FooDaxSrc{})
		base.AddLocalDaxSrc("pubsub", FooDaxSrc{})
		base.AddLocalDaxSrc("json", FooDaxSrc{})
		base.AddLocalDaxSrc("env", FooDaxSrc{})

		sabi.FreeAllLocalDaxSrcsForTest(base)
	}
	b.StopTimer()
}
