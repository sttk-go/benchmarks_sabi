package dax2_ready_sync_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/dax2_ready_sync"
	"testing"
)

func Benchmark_GetDaxConn_global_oneDs(b *testing.B) {
	sabi.AddGlobalDaxSrc("foo", FooDaxSrc{})
	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := sabi.GetDaxConn[FooDaxConn](base, "foo")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func Benchmark_GetDaxConn_global_fiveDs(b *testing.B) {
	sabi.AddGlobalDaxSrc("foo", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("bar", BarDaxSrc{})
	sabi.AddGlobalDaxSrc("baz", BazDaxSrc{})
	sabi.AddGlobalDaxSrc("qux", QuxDaxSrc{})
	sabi.AddGlobalDaxSrc("corge", CorgeDaxSrc{})

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn0, err := sabi.GetDaxConn[FooDaxConn](base, "foo")
		conn1, err := sabi.GetDaxConn[BarDaxConn](base, "bar")
		conn2, err := sabi.GetDaxConn[BazDaxConn](base, "baz")
		conn3, err := sabi.GetDaxConn[QuxDaxConn](base, "qux")
		conn4, err := sabi.GetDaxConn[CorgeDaxConn](base, "corge")
		_ = conn0
		_ = conn1
		_ = conn2
		_ = conn3
		_ = conn4
		_ = err
	}
	b.StopTimer()
}

func Benchmark_GetDaxConn_local_oneDs(b *testing.B) {
	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("foo", FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := sabi.GetDaxConn[FooDaxConn](base, "foo")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func Benchmark_GetDaxConn_local_fiveDs(b *testing.B) {
	base := sabi.NewDaxBase()

	base.SetUpLocalDaxSrc("foo", FooDaxSrc{})
	base.SetUpLocalDaxSrc("bar", BarDaxSrc{})
	base.SetUpLocalDaxSrc("baz", BazDaxSrc{})
	base.SetUpLocalDaxSrc("qux", QuxDaxSrc{})
	base.SetUpLocalDaxSrc("corge", CorgeDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn0, err := sabi.GetDaxConn[FooDaxConn](base, "foo")
		conn1, err := sabi.GetDaxConn[BarDaxConn](base, "bar")
		conn2, err := sabi.GetDaxConn[BazDaxConn](base, "baz")
		conn3, err := sabi.GetDaxConn[QuxDaxConn](base, "qux")
		conn4, err := sabi.GetDaxConn[CorgeDaxConn](base, "corge")
		_ = conn0
		_ = conn1
		_ = conn2
		_ = conn3
		_ = conn4
		_ = err
	}
	b.StopTimer()
}
