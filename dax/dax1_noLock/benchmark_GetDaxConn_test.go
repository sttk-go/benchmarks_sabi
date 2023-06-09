package dax1_noLock_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/dax1_noLock"
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
		conn, err := sabi.GetDaxConn[FooDaxConn](base, "foo")
		conn, err = sabi.GetDaxConn[FooDaxConn](base, "bar")
		conn, err = sabi.GetDaxConn[FooDaxConn](base, "baz")
		conn, err = sabi.GetDaxConn[FooDaxConn](base, "qux")
		conn, err = sabi.GetDaxConn[FooDaxConn](base, "corge")
		_ = conn
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
		conn, err := sabi.GetDaxConn[FooDaxConn](base, "foo")
		conn, err = sabi.GetDaxConn[FooDaxConn](base, "bar")
		conn, err = sabi.GetDaxConn[FooDaxConn](base, "baz")
		conn, err = sabi.GetDaxConn[FooDaxConn](base, "qux")
		conn, err = sabi.GetDaxConn[FooDaxConn](base, "corge")
		_ = conn
		_ = err
	}
	b.StopTimer()
}
