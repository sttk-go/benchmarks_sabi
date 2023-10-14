package v0_5_0_test

import (
	"testing"

	prev "github.com/sttk/benchmarks_sabi/dax/v0_4_0"
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_5_0"

	prev_supp "github.com/sttk/benchmarks_sabi/dax/v0_4_0/supp"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_5_0/supp"
)

func BenchmarkDax_____CloseGlobalDaxSrcs_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", supp.FooDaxSrc{})
	sabi.SetupGlobalDaxSrcs()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_CloseGlobalDaxSrcs_zeroDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	prev.StartUpGlobalDaxSrcs()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
	}
	b.StopTimer()
}

func BenchmarkDax_____CloseGlobalDaxSrcs_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", supp.FooDaxSrc{})
	sabi.SetupGlobalDaxSrcs()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.CloseGlobalDaxSrcs()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_CloseGlobalDaxSrcs_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	prev.StartUpGlobalDaxSrcs()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.ShutdownGlobalDaxSrcs()
	}
	b.StopTimer()
}

func BenchmarkDax_____CloseGlobalDaxSrcs_freeDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("database", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("pubsub", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("json", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("env", supp.FooDaxSrc{})
	sabi.SetupGlobalDaxSrcs()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.CloseGlobalDaxSrcs()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_CloseGlobalDaxSrcs_freeDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("database", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("json", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("env", prev_supp.FooDaxSrc{})
	prev.StartUpGlobalDaxSrcs()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.ShutdownGlobalDaxSrcs()
	}
	b.StopTimer()
}
