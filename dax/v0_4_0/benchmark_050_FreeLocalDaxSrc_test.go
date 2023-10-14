package v0_4_0_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_4_0/supp"
)

func Benchmark_FreeLocalDaxSrc_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddLocalDaxSrcForTest(base, "cliargs", supp.FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "database", supp.FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "pubsub", supp.FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "json", supp.FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "env", supp.FooDaxSrc{})
	}
	b.StopTimer()
}

func Benchmark_FreeLocalDaxSrc_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddLocalDaxSrcForTest(base, "cliargs", supp.FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "database", supp.FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "pubsub", supp.FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "json", supp.FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "env", supp.FooDaxSrc{})

		base.FreeLocalDaxSrc("cliargs")
	}
	b.StopTimer()
}

func Benchmark_FreeLocalDaxSrc_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddLocalDaxSrcForTest(base, "cliargs", supp.FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "database", supp.FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "pubsub", supp.FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "json", supp.FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "env", supp.FooDaxSrc{})

		base.FreeLocalDaxSrc("cliargs")
		base.FreeLocalDaxSrc("database")
		base.FreeLocalDaxSrc("pubsub")
		base.FreeLocalDaxSrc("json")
		base.FreeLocalDaxSrc("env")
	}
	b.StopTimer()
}
