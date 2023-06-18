package ready_async_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/sabi_0_4_0/ready_async"
)

func Benchmark_RunTxn_commit_oneDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})

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
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})
	base.SetUpLocalDaxSrc("database", FooDaxSrc{})
	base.SetUpLocalDaxSrc("pubsub", FooDaxSrc{})
	base.SetUpLocalDaxSrc("json", FooDaxSrc{})
	base.SetUpLocalDaxSrc("env", FooDaxSrc{})

	logic := func(dax sabi.Dax) sabi.Err {
		conn0, err := sabi.GetDaxConn[FooDaxConn](dax, "cliargs")
		conn1, err := sabi.GetDaxConn[FooDaxConn](dax, "database")
		conn2, err := sabi.GetDaxConn[FooDaxConn](dax, "pubsub")
		conn3, err := sabi.GetDaxConn[FooDaxConn](dax, "json")
		conn4, err := sabi.GetDaxConn[FooDaxConn](dax, "env")
		_ = conn0
		_ = conn1
		_ = conn2
		_ = conn3
		_ = conn4
		_ = err
		return sabi.Ok()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func Benchmark_RunTxn_rollback_oneDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})

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
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})
	base.SetUpLocalDaxSrc("database", FooDaxSrc{})
	base.SetUpLocalDaxSrc("pubsub", FooDaxSrc{})
	base.SetUpLocalDaxSrc("json", FooDaxSrc{})
	base.SetUpLocalDaxSrc("env", FooDaxSrc{})

	type Fail struct{}
	logic := func(dax sabi.Dax) sabi.Err {
		conn0, err := sabi.GetDaxConn[FooDaxConn](dax, "cliargs")
		conn1, err := sabi.GetDaxConn[FooDaxConn](dax, "database")
		conn2, err := sabi.GetDaxConn[FooDaxConn](dax, "pubsub")
		conn3, err := sabi.GetDaxConn[FooDaxConn](dax, "json")
		conn4, err := sabi.GetDaxConn[FooDaxConn](dax, "env")
		_ = conn0
		_ = conn1
		_ = conn2
		_ = conn3
		_ = conn4
		_ = err
		return sabi.NewErr(Fail{})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}
