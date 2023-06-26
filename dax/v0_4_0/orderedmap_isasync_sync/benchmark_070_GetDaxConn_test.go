package orderedmap_isasync_sync_test

import (
	"sync"
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_isasync_sync"
)

func Benchmark_GetDaxConn_global_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})

	var wg sync.WaitGroup
	base := sabi.NewDaxBase(&wg)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := sabi.GetDaxConn[FooDaxConn](base, "cliargs")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func Benchmark_GetDaxConn_global_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("database", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("pubsub", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("json", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("env", FooDaxSrc{})

	var wg sync.WaitGroup
	base := sabi.NewDaxBase(&wg)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn0, err0 := sabi.GetDaxConn[FooDaxConn](base, "cliargs")
		conn1, err1 := sabi.GetDaxConn[FooDaxConn](base, "database")
		conn2, err2 := sabi.GetDaxConn[FooDaxConn](base, "pubsub")
		conn3, err3 := sabi.GetDaxConn[FooDaxConn](base, "json")
		conn4, err4 := sabi.GetDaxConn[FooDaxConn](base, "env")
		_ = conn0
		_ = conn1
		_ = conn2
		_ = conn3
		_ = conn4
		_ = err0
		_ = err1
		_ = err2
		_ = err3
		_ = err4
	}
	b.StopTimer()
}

func Benchmark_GetDaxConn_local_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	var wg sync.WaitGroup
	base := sabi.NewDaxBase(&wg)
	base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := sabi.GetDaxConn[FooDaxConn](base, "cliargs")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func Benchmark_GetDaxConn_local_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	var wg sync.WaitGroup
	base := sabi.NewDaxBase(&wg)
	base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})
	base.SetUpLocalDaxSrc("database", FooDaxSrc{})
	base.SetUpLocalDaxSrc("pubsub", FooDaxSrc{})
	base.SetUpLocalDaxSrc("json", FooDaxSrc{})
	base.SetUpLocalDaxSrc("env", FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn0, err0 := sabi.GetDaxConn[FooDaxConn](base, "cliargs")
		conn1, err1 := sabi.GetDaxConn[FooDaxConn](base, "database")
		conn2, err2 := sabi.GetDaxConn[FooDaxConn](base, "pubsub")
		conn3, err3 := sabi.GetDaxConn[FooDaxConn](base, "json")
		conn4, err4 := sabi.GetDaxConn[FooDaxConn](base, "env")
		_ = conn0
		_ = conn1
		_ = conn2
		_ = conn3
		_ = conn4
		_ = err0
		_ = err1
		_ = err2
		_ = err3
		_ = err4
	}
	b.StopTimer()
}
