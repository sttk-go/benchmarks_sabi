package ready_async_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/ready_async"
	"testing"
)

func Benchmark_GetDaxConn_global_oneDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := sabi.GetDaxConn[FooDaxConn](base, "cliargs")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func Benchmark_GetDaxConn_global_fiveDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("database", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("pubsub", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("json", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("env", FooDaxSrc{})

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn0, err := sabi.GetDaxConn[FooDaxConn](base, "cliargs")
		conn1, err := sabi.GetDaxConn[FooDaxConn](base, "database")
		conn2, err := sabi.GetDaxConn[FooDaxConn](base, "pubsub")
		conn3, err := sabi.GetDaxConn[FooDaxConn](base, "json")
		conn4, err := sabi.GetDaxConn[FooDaxConn](base, "env")
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
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
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
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})
	base.SetUpLocalDaxSrc("database", FooDaxSrc{})
	base.SetUpLocalDaxSrc("pubsub", FooDaxSrc{})
	base.SetUpLocalDaxSrc("json", FooDaxSrc{})
	base.SetUpLocalDaxSrc("env", FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn0, err := sabi.GetDaxConn[FooDaxConn](base, "cliargs")
		conn1, err := sabi.GetDaxConn[FooDaxConn](base, "database")
		conn2, err := sabi.GetDaxConn[FooDaxConn](base, "pubsub")
		conn3, err := sabi.GetDaxConn[FooDaxConn](base, "json")
		conn4, err := sabi.GetDaxConn[FooDaxConn](base, "env")
		_ = conn0
		_ = conn1
		_ = conn2
		_ = conn3
		_ = conn4
		_ = err
	}
	b.StopTimer()
}
