package v0_6_0_test

import (
	"testing"

	prev "github.com/sttk/benchmarks_sabi/dax/v0_5_0"
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_6_0"
	errs "github.com/sttk/benchmarks_sabi/dax/v0_6_0/errs"

	prev_supp "github.com/sttk/benchmarks_sabi/dax/v0_5_0/supp"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_6_0/supp"
)

func BenchmarkDax_____Sabi_Txn_commit_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.Uses("cliargs", supp.FooDaxSrc{})

	logic := func(dax sabi.Dax) errs.Err {
		conn, err := sabi.GetDaxConn[supp.FooDaxConn](dax, "cliargs")
		_ = conn
		_ = err
		return errs.Ok()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.Txn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_Txn_commit_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()
	base.AddLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})

	logic := func(dax prev.Dax) prev.Err {
		conn, err := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "cliargs")
		_ = conn
		_ = err
		return prev.Ok()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.RunTxn[prev.Dax](base, logic)
	}
	b.StopTimer()
}

func BenchmarkDax_____Sabi_Txn_commit_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.Uses("cliargs", supp.FooDaxSrc{})
	base.Uses("database", supp.FooDaxSrc{})
	base.Uses("pubsub", supp.FooDaxSrc{})
	base.Uses("json", supp.FooDaxSrc{})
	base.Uses("env", supp.FooDaxSrc{})

	logic := func(dax sabi.Dax) errs.Err {
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
		return errs.Ok()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.Txn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_Txn_commit_fiveDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()
	base.AddLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	base.AddLocalDaxSrc("database", prev_supp.FooDaxSrc{})
	base.AddLocalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
	base.AddLocalDaxSrc("json", prev_supp.FooDaxSrc{})
	base.AddLocalDaxSrc("env", prev_supp.FooDaxSrc{})

	logic := func(dax prev.Dax) prev.Err {
		conn0, err0 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "cliargs")
		conn1, err1 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "database")
		conn2, err2 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "pubsub")
		conn3, err3 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "json")
		conn4, err4 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "env")
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
		return prev.Ok()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.RunTxn[prev.Dax](base, logic)
	}
	b.StopTimer()
}

func BenchmarkDax_____Sabi_Txn_rollback_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.Uses("cliargs", supp.FooDaxSrc{})

	type Fail struct{}

	logic := func(dax sabi.Dax) errs.Err {
		conn, err := sabi.GetDaxConn[supp.FooDaxConn](dax, "cliargs")
		_ = conn
		_ = err
		return errs.New(Fail{})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.Txn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_Txn_rollback_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()
	base.AddLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})

	type Fail struct{}

	logic := func(dax prev.Dax) prev.Err {
		conn, err := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "cliargs")
		_ = conn
		_ = err
		return prev.NewErr(Fail{})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.RunTxn[prev.Dax](base, logic)
	}
	b.StopTimer()
}

func BenchmarkDax_____Sabi_Txn_rollback_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.Uses("cliargs", supp.FooDaxSrc{})
	base.Uses("database", supp.FooDaxSrc{})
	base.Uses("pubsub", supp.FooDaxSrc{})
	base.Uses("json", supp.FooDaxSrc{})
	base.Uses("env", supp.FooDaxSrc{})

	type Fail struct{}

	logic := func(dax sabi.Dax) errs.Err {
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
		return errs.New(Fail{})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.Txn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_Txn_rollback_fiveDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()
	base.AddLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	base.AddLocalDaxSrc("database", prev_supp.FooDaxSrc{})
	base.AddLocalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
	base.AddLocalDaxSrc("json", prev_supp.FooDaxSrc{})
	base.AddLocalDaxSrc("env", prev_supp.FooDaxSrc{})

	type Fail struct{}

	logic := func(dax prev.Dax) prev.Err {
		conn0, err0 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "cliargs")
		conn1, err1 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "database")
		conn2, err2 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "pubsub")
		conn3, err3 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "json")
		conn4, err4 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "env")
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
		return prev.NewErr(Fail{})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.RunTxn[prev.Dax](base, logic)
	}
	b.StopTimer()
}
