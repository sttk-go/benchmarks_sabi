package v0_6_0_test

import (
	"testing"

	prev "github.com/sttk/benchmarks_sabi/dax/v0_5_0"
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_6_0"

	prev_supp "github.com/sttk/benchmarks_sabi/dax/v0_5_0/supp"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_6_0/supp"
)

func BenchmarkDax_____DaxBase_Disuses_zeroDs(b *testing.B) {
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

func BenchmarkDaxPrev_DaxBase_Disuses_zeroDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.AddLocalDaxSrcForTest(base, "cliargs", prev_supp.FooDaxSrc{})
		prev.AddLocalDaxSrcForTest(base, "database", prev_supp.FooDaxSrc{})
		prev.AddLocalDaxSrcForTest(base, "pubsub", prev_supp.FooDaxSrc{})
		prev.AddLocalDaxSrcForTest(base, "json", prev_supp.FooDaxSrc{})
		prev.AddLocalDaxSrcForTest(base, "env", prev_supp.FooDaxSrc{})
	}
	b.StopTimer()
}

func BenchmarkDax_____DaxBase_Disuses_oneDs(b *testing.B) {
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

		base.Disuses("cliargs")
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_DaxBase_Disuses_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.AddLocalDaxSrcForTest(base, "cliargs", prev_supp.FooDaxSrc{})
		prev.AddLocalDaxSrcForTest(base, "database", prev_supp.FooDaxSrc{})
		prev.AddLocalDaxSrcForTest(base, "pubsub", prev_supp.FooDaxSrc{})
		prev.AddLocalDaxSrcForTest(base, "json", prev_supp.FooDaxSrc{})
		prev.AddLocalDaxSrcForTest(base, "env", prev_supp.FooDaxSrc{})

		base.FreeLocalDaxSrc("cliargs")
	}
	b.StopTimer()
}

func BenchmarkDax_____DaxBase_Disuses_fiveDs(b *testing.B) {
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

		base.Disuses("cliargs")
		base.Disuses("database")
		base.Disuses("pubsub")
		base.Disuses("json")
		base.Disuses("env")
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_DaxBase_Disuses_fiveDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.AddLocalDaxSrcForTest(base, "cliargs", prev_supp.FooDaxSrc{})
		prev.AddLocalDaxSrcForTest(base, "database", prev_supp.FooDaxSrc{})
		prev.AddLocalDaxSrcForTest(base, "pubsub", prev_supp.FooDaxSrc{})
		prev.AddLocalDaxSrcForTest(base, "json", prev_supp.FooDaxSrc{})
		prev.AddLocalDaxSrcForTest(base, "env", prev_supp.FooDaxSrc{})

		base.FreeLocalDaxSrc("cliargs")
		base.FreeLocalDaxSrc("database")
		base.FreeLocalDaxSrc("pubsub")
		base.FreeLocalDaxSrc("json")
		base.FreeLocalDaxSrc("env")
	}
	b.StopTimer()
}
