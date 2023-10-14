package v0_6_0_test

import (
	"testing"

	prev "github.com/sttk/benchmarks_sabi/dax/v0_5_0"
	prev_supp "github.com/sttk/benchmarks_sabi/dax/v0_5_0/supp"
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_6_0"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_6_0/supp"
)

func BenchmarkDax_____Sabi_Uses_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.ResetGlobals()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_Uses_zeroDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.ResetGlobals()
	}
	b.StopTimer()
}

func BenchmarkDax_____Sabi_Uses_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.Uses("cliargs", supp.FooDaxSrc{})
		sabi.ResetGlobals()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_Uses_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
		prev.ResetGlobals()
	}
	b.StopTimer()
}

func BenchmarkDax_____Sabi_Uses_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.Uses("cliargs", supp.FooDaxSrc{})
		sabi.Uses("database", supp.FooDaxSrc{})
		sabi.Uses("pubsub", supp.FooDaxSrc{})
		sabi.Uses("json", supp.FooDaxSrc{})
		sabi.Uses("env", supp.FooDaxSrc{})

		sabi.ResetGlobals()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_Uses_fiveDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
		prev.AddGlobalDaxSrc("database", prev_supp.FooDaxSrc{})
		prev.AddGlobalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
		prev.AddGlobalDaxSrc("json", prev_supp.FooDaxSrc{})
		prev.AddGlobalDaxSrc("env", prev_supp.FooDaxSrc{})

		prev.ResetGlobals()
	}
	b.StopTimer()
}
