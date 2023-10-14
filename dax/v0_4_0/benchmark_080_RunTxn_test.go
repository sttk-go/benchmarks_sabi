package v0_4_0_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_4_0"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_4_0/supp"
)

func Benchmark_RunTxn_commit_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", supp.FooDaxSrc{})

	logic := func(dax sabi.Dax) sabi.Err {
		conn, err := sabi.GetDaxConn[supp.FooDaxConn](dax, "cliargs")
		_ = conn
		_ = err
		return sabi.Ok()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func Benchmark_RunTxn_commit_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("database", supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("pubsub", supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("json", supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("env", supp.FooDaxSrc{})

	logic := func(dax sabi.Dax) sabi.Err {
		conn0, err0 := sabi.GetDaxConn[supp.FooDaxConn](dax, "cliargs")
		conn1, err1 := sabi.GetDaxConn[supp.FooDaxConn](dax, "database")
		conn2, err2 := sabi.GetDaxConn[supp.FooDaxConn](dax, "pubsub")
		conn3, err3 := sabi.GetDaxConn[supp.FooDaxConn](dax, "json")
		conn4, err4 := sabi.GetDaxConn[supp.FooDaxConn](dax, "env")
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
		return sabi.Ok()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func Benchmark_RunTxn_rollback_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", supp.FooDaxSrc{})

	type Fail struct{}

	logic := func(dax sabi.Dax) sabi.Err {
		conn, err := sabi.GetDaxConn[supp.FooDaxConn](dax, "cliargs")
		_ = conn
		_ = err
		return sabi.NewErr(Fail{})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func Benchmark_RunTxn_rollback_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("database", supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("pubsub", supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("json", supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("env", supp.FooDaxSrc{})

	type Fail struct{}

	logic := func(dax sabi.Dax) sabi.Err {
		conn0, err0 := sabi.GetDaxConn[supp.FooDaxConn](dax, "cliargs")
		conn1, err1 := sabi.GetDaxConn[supp.FooDaxConn](dax, "database")
		conn2, err2 := sabi.GetDaxConn[supp.FooDaxConn](dax, "pubsub")
		conn3, err3 := sabi.GetDaxConn[supp.FooDaxConn](dax, "json")
		conn4, err4 := sabi.GetDaxConn[supp.FooDaxConn](dax, "env")
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
		return sabi.NewErr(Fail{})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}
