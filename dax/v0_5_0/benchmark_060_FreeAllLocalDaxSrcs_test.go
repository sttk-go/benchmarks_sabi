package v0_5_0_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_5_0"
)

func Benchmark_FreeAllLocalDaxSrcs_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}

func Benchmark_FreeAllLocalDaxSrcs_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.AddLocalDaxSrc("cliargs", FooDaxSrc{})
		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}

func Benchmark_FreeAllLocalDaxSrcs_fiveDs(b *testing.B) {
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

		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}
