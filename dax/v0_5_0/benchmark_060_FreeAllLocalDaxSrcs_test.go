package v0_5_0_test

import (
	"testing"

	prev "github.com/sttk/benchmarks_sabi/dax/v0_4_0"
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_5_0"

	prev_supp "github.com/sttk/benchmarks_sabi/dax/v0_4_0/supp"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_5_0/supp"
)

func BenchmarkDax_____FreeAllLocalDaxSrcs_zeroDs(b *testing.B) {
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

func BenchmarkDaxPrev_FreeAllLocalDaxSrcs_zeroDs(b *testing.B) {
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

func BenchmarkDax_____FreeAllLocalDaxSrcs_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.AddLocalDaxSrc("cliargs", supp.FooDaxSrc{})
		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_FreeAllLocalDaxSrcs_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.SetUpLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}

func BenchmarkDax_____FreeAllLocalDaxSrcs_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.AddLocalDaxSrc("cliargs", supp.FooDaxSrc{})
		base.AddLocalDaxSrc("database", supp.FooDaxSrc{})
		base.AddLocalDaxSrc("pubsub", supp.FooDaxSrc{})
		base.AddLocalDaxSrc("json", supp.FooDaxSrc{})
		base.AddLocalDaxSrc("env", supp.FooDaxSrc{})

		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_FreeAllLocalDaxSrcs_fiveDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.SetUpLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
		base.SetUpLocalDaxSrc("database", prev_supp.FooDaxSrc{})
		base.SetUpLocalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
		base.SetUpLocalDaxSrc("json", prev_supp.FooDaxSrc{})
		base.SetUpLocalDaxSrc("env", prev_supp.FooDaxSrc{})

		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}
