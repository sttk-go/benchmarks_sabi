package v0_4_0_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_4_0/supp"
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
		base.SetUpLocalDaxSrc("cliargs", supp.FooDaxSrc{})
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
		base.SetUpLocalDaxSrc("cliargs", supp.FooDaxSrc{})
		base.SetUpLocalDaxSrc("database", supp.FooDaxSrc{})
		base.SetUpLocalDaxSrc("pubsub", supp.FooDaxSrc{})
		base.SetUpLocalDaxSrc("json", supp.FooDaxSrc{})
		base.SetUpLocalDaxSrc("env", supp.FooDaxSrc{})

		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}
