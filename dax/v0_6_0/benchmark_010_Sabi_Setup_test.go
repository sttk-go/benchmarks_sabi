package v0_6_0_test

import (
	"testing"

	prev "github.com/sttk/benchmarks_sabi/dax/v0_5_0"
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_6_0"

	prev_supp "github.com/sttk/benchmarks_sabi/dax/v0_5_0/supp"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_6_0/supp"
)

func BenchmarkDax_____Sabi_Setup_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.Uses("cliargs", supp.FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_Setup_zeroDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
	}
	b.StopTimer()
}

func BenchmarkDax_____Sabi_Setup_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.Uses("cliargs", supp.FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.Setup()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_Setup_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.SetupGlobalDaxSrcs()
	}
	b.StopTimer()
}

func BenchmarkDax_____Sabi_Setup_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.Uses("cliargs", supp.FooDaxSrc{})
	sabi.Uses("database", supp.FooDaxSrc{})
	sabi.Uses("pubsub", supp.FooDaxSrc{})
	sabi.Uses("json", supp.FooDaxSrc{})
	sabi.Uses("env", supp.FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.Setup()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_Setup_fiveDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("database", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("json", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("env", prev_supp.FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.SetupGlobalDaxSrcs()
	}
	b.StopTimer()
}
