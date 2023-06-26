package orderedmap_isasync_sync_test

import (
	"sync"
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_isasync_sync"
)

func Benchmark_SetUpLocalDaxSrc_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	var wg sync.WaitGroup
	base := sabi.NewDaxBase(&wg)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.FreeAllLocalDaxSrcsForTest(base)
	}
	b.StopTimer()
}

func Benchmark_SetUpLocalDaxSrc_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	var wg sync.WaitGroup
	base := sabi.NewDaxBase(&wg)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})
		sabi.FreeAllLocalDaxSrcsForTest(base)
	}
	b.StopTimer()
}

func Benchmark_SetUpLocalDaxSrc_fiveDs(b *testing.B) {
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

		sabi.FreeAllLocalDaxSrcsForTest(base)
	}
	b.StopTimer()
}
