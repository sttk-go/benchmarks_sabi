package v0_6_0_test

import (
	"testing"

	prev "github.com/sttk/benchmarks_sabi/dax/v0_5_0"
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_6_0"

	prev_supp "github.com/sttk/benchmarks_sabi/dax/v0_5_0/supp"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_6_0/supp"
)

func BenchmarkDax_____Sabi_GetDaxConn_global_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.Uses("cliargs", supp.FooDaxSrc{})

	base := sabi.NewDaxBase()

	sabi.Begin(base)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := sabi.GetDaxConn[supp.FooDaxConn](base, "cliargs")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_GetDaxConn_global_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})

	base := prev.NewDaxBase()

	prev.Begin(base)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := prev.GetDaxConn[prev_supp.FooDaxConn](base, "cliargs")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func BenchmarkDax_____Sabi_GetDaxConn_global_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.Uses("cliargs", supp.FooDaxSrc{})
	sabi.Uses("database", supp.FooDaxSrc{})
	sabi.Uses("pubsub", supp.FooDaxSrc{})
	sabi.Uses("json", supp.FooDaxSrc{})
	sabi.Uses("env", supp.FooDaxSrc{})

	base := sabi.NewDaxBase()

	sabi.Begin(base)

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

func BenchmarkDaxPrev_Sabi_GetDaxConn_global_fiveDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("database", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("json", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("env", prev_supp.FooDaxSrc{})

	base := prev.NewDaxBase()

	prev.Begin(base)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn0, err0 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "cliargs")
		conn1, err1 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "database")
		conn2, err2 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "pubsub")
		conn3, err3 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "json")
		conn4, err4 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "env")
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

func BenchmarkDax_____Sabi_GetDaxConn_local_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.Uses("cliargs", supp.FooDaxSrc{})

	sabi.Begin(base)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := sabi.GetDaxConn[supp.FooDaxConn](base, "cliargs")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_Sabi_GetDaxConn_local_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()
	base.AddLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})

	prev.Begin(base)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := prev.GetDaxConn[prev_supp.FooDaxConn](base, "cliargs")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func BenchmarkDax_____Sabi_GetDaxConn_local_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	base.Uses("cliargs", supp.FooDaxSrc{})
	base.Uses("database", supp.FooDaxSrc{})
	base.Uses("pubsub", supp.FooDaxSrc{})
	base.Uses("json", supp.FooDaxSrc{})
	base.Uses("env", supp.FooDaxSrc{})

	sabi.Begin(base)

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

func BenchmarkDaxPrev_Sabi_GetDaxConn_local_fiveDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()

	base.AddLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	base.AddLocalDaxSrc("database", prev_supp.FooDaxSrc{})
	base.AddLocalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
	base.AddLocalDaxSrc("json", prev_supp.FooDaxSrc{})
	base.AddLocalDaxSrc("env", prev_supp.FooDaxSrc{})

	prev.Begin(base)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn0, err0 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "cliargs")
		conn1, err1 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "database")
		conn2, err2 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "pubsub")
		conn3, err3 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "json")
		conn4, err4 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "env")
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
