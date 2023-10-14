package v0_4_0_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_4_0/supp"
)

func Benchmark_GetDaxConn_global_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", supp.FooDaxSrc{})

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := sabi.GetDaxConn[supp.FooDaxConn](base, "cliargs")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func Benchmark_GetDaxConn_global_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("database", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("pubsub", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("json", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("env", supp.FooDaxSrc{})

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn0, err0 := sabi.GetDaxConn[supp.FooDaxConn](base, "cliargs")
		conn1, err1 := sabi.GetDaxConn[supp.FooDaxConn](base, "database")
		conn2, err2 := sabi.GetDaxConn[supp.FooDaxConn](base, "pubsub")
		conn3, err3 := sabi.GetDaxConn[supp.FooDaxConn](base, "json")
		conn4, err4 := sabi.GetDaxConn[supp.FooDaxConn](base, "env")
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

	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", supp.FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := sabi.GetDaxConn[supp.FooDaxConn](base, "cliargs")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func Benchmark_GetDaxConn_local_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("database", supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("pubsub", supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("json", supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("env", supp.FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn0, err0 := sabi.GetDaxConn[supp.FooDaxConn](base, "cliargs")
		conn1, err1 := sabi.GetDaxConn[supp.FooDaxConn](base, "database")
		conn2, err2 := sabi.GetDaxConn[supp.FooDaxConn](base, "pubsub")
		conn3, err3 := sabi.GetDaxConn[supp.FooDaxConn](base, "json")
		conn4, err4 := sabi.GetDaxConn[supp.FooDaxConn](base, "env")
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
