package orderedmap_isasync_async_test

import (
	"sync"
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_isasync_async"
)

func Benchmark_FreeAllLocalDaxSrcs_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	var wg sync.WaitGroup
	base := sabi.NewDaxBase(&wg)

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

	var wg sync.WaitGroup
	base := sabi.NewDaxBase(&wg)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})
		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}

func Benchmark_FreeAllLocalDaxSrcs_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	var wg sync.WaitGroup
	base := sabi.NewDaxBase(&wg)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})
		base.SetUpLocalDaxSrc("database", FooDaxSrc{})
		base.SetUpLocalDaxSrc("pubsub", FooDaxSrc{})
		base.SetUpLocalDaxSrc("json", FooDaxSrc{})
		base.SetUpLocalDaxSrc("env", FooDaxSrc{})

		base.FreeAllLocalDaxSrcs()
	}
	b.StopTimer()
}
