package v0_6_0_test

import (
	"testing"

	prev "github.com/sttk/benchmarks_sabi/dax/v0_5_0"
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_6_0"

	prev_supp "github.com/sttk/benchmarks_sabi/dax/v0_5_0/supp"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_6_0/supp"
)

func BenchmarkDax_____Sabi_Close_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.Uses("cliargs", supp.FooDaxSrc{})
	sabi.Setup()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_Close_zeroDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	prev.SetupGlobalDaxSrcs()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
	}
	b.StopTimer()
}

func BenchmarkDax_____Sabi_Close_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.Uses("cliargs", supp.FooDaxSrc{})
	sabi.Setup()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.Close()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_Close_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	prev.SetupGlobalDaxSrcs()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.CloseGlobalDaxSrcs()
	}
	b.StopTimer()
}

func BenchmarkDax_____Sabi_Close_freeDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.Uses("cliargs", supp.FooDaxSrc{})
	sabi.Uses("database", supp.FooDaxSrc{})
	sabi.Uses("pubsub", supp.FooDaxSrc{})
	sabi.Uses("json", supp.FooDaxSrc{})
	sabi.Uses("env", supp.FooDaxSrc{})
	sabi.Setup()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.Close()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_Close_freeDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("database", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("json", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("env", prev_supp.FooDaxSrc{})
	prev.SetupGlobalDaxSrcs()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.CloseGlobalDaxSrcs()
	}
	b.StopTimer()
}
