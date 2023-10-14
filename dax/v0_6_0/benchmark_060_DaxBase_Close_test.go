package v0_6_0_test

import (
	"testing"

	prev "github.com/sttk/benchmarks_sabi/dax/v0_5_0"
	prev_supp "github.com/sttk/benchmarks_sabi/dax/v0_5_0/supp"
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_6_0"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_6_0/supp"
)

func BenchmarkDax_____DaxBase_Close_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.Close()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_DaxBase_Close_zeroDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}

func BenchmarkDax_____DaxBase_Close_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.Uses("cliargs", supp.FooDaxSrc{})
		base.Close()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_DaxBase_Close_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.AddLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}

func BenchmarkDax_____DaxBase_Close_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.Uses("cliargs", supp.FooDaxSrc{})
		base.Uses("database", supp.FooDaxSrc{})
		base.Uses("pubsub", supp.FooDaxSrc{})
		base.Uses("json", supp.FooDaxSrc{})
		base.Uses("env", supp.FooDaxSrc{})

		base.Close()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_DaxBase_Close_fiveDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.AddLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
		base.AddLocalDaxSrc("database", prev_supp.FooDaxSrc{})
		base.AddLocalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
		base.AddLocalDaxSrc("json", prev_supp.FooDaxSrc{})
		base.AddLocalDaxSrc("env", prev_supp.FooDaxSrc{})

		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}
