package register_daxsrc_in_list_sync_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0/register_daxsrc_in_list_sync"
)

func Benchmark_AddGlobalDaxSrc_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.ResetGlobals()
	}
	b.StopTimer()
}

func Benchmark_AddGlobalDaxSrc_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})
		sabi.ResetGlobals()
	}
	b.StopTimer()
}

func Benchmark_AddGlobalDaxSrc_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})
		sabi.AddGlobalDaxSrc("database", FooDaxSrc{})
		sabi.AddGlobalDaxSrc("pubsub", FooDaxSrc{})
		sabi.AddGlobalDaxSrc("json", FooDaxSrc{})
		sabi.AddGlobalDaxSrc("env", FooDaxSrc{})

		sabi.ResetGlobals()
	}
	b.StopTimer()
}
