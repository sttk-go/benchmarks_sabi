package v0_5_0_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_5_0"
)

func Benchmark_RunTxn_commit_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.AddLocalDaxSrc("cliargs", FooDaxSrc{})

	logic := func(dax sabi.Dax) sabi.Err {
		conn, err := sabi.GetDaxConn[FooDaxConn](dax, "cliargs")
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
	base.AddLocalDaxSrc("cliargs", FooDaxSrc{})
	base.AddLocalDaxSrc("database", FooDaxSrc{})
	base.AddLocalDaxSrc("pubsub", FooDaxSrc{})
	base.AddLocalDaxSrc("json", FooDaxSrc{})
	base.AddLocalDaxSrc("env", FooDaxSrc{})

	logic := func(dax sabi.Dax) sabi.Err {
		conn0, err0 := sabi.GetDaxConn[FooDaxConn](dax, "cliargs")
		conn1, err1 := sabi.GetDaxConn[FooDaxConn](dax, "database")
		conn2, err2 := sabi.GetDaxConn[FooDaxConn](dax, "pubsub")
		conn3, err3 := sabi.GetDaxConn[FooDaxConn](dax, "json")
		conn4, err4 := sabi.GetDaxConn[FooDaxConn](dax, "env")
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
	base.AddLocalDaxSrc("cliargs", FooDaxSrc{})

	type Fail struct{}

	logic := func(dax sabi.Dax) sabi.Err {
		conn, err := sabi.GetDaxConn[FooDaxConn](dax, "cliargs")
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
	base.AddLocalDaxSrc("cliargs", FooDaxSrc{})
	base.AddLocalDaxSrc("database", FooDaxSrc{})
	base.AddLocalDaxSrc("pubsub", FooDaxSrc{})
	base.AddLocalDaxSrc("json", FooDaxSrc{})
	base.AddLocalDaxSrc("env", FooDaxSrc{})

	type Fail struct{}

	logic := func(dax sabi.Dax) sabi.Err {
		conn0, err0 := sabi.GetDaxConn[FooDaxConn](dax, "cliargs")
		conn1, err1 := sabi.GetDaxConn[FooDaxConn](dax, "database")
		conn2, err2 := sabi.GetDaxConn[FooDaxConn](dax, "pubsub")
		conn3, err3 := sabi.GetDaxConn[FooDaxConn](dax, "json")
		conn4, err4 := sabi.GetDaxConn[FooDaxConn](dax, "env")
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
