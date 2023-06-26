package orderedmap_isasync_sync_test

import (
	"sync"
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_isasync_sync"
)

func Benchmark_FreeLocalDaxSrc_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	var wg sync.WaitGroup
	base := sabi.NewDaxBase(&wg)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddLocalDaxSrcForTest(base, "cliargs", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "database", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "pubsub", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "json", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "env", FooDaxSrc{})
	}
	b.StopTimer()
}

func Benchmark_FreeLocalDaxSrc_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	var wg sync.WaitGroup
	base := sabi.NewDaxBase(&wg)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddLocalDaxSrcForTest(base, "cliargs", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "database", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "pubsub", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "json", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "env", FooDaxSrc{})

		base.FreeLocalDaxSrc("cliargs")
	}
	b.StopTimer()
}

func Benchmark_FreeLocalDaxSrc_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	var wg sync.WaitGroup
	base := sabi.NewDaxBase(&wg)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddLocalDaxSrcForTest(base, "cliargs", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "database", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "pubsub", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "json", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "env", FooDaxSrc{})

		base.FreeLocalDaxSrc("cliargs")
		base.FreeLocalDaxSrc("database")
		base.FreeLocalDaxSrc("pubsub")
		base.FreeLocalDaxSrc("json")
		base.FreeLocalDaxSrc("env")
	}
	b.StopTimer()
}
