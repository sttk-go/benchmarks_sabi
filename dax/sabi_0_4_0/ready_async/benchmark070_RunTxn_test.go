package ready_async_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/ready_async"
	"testing"
)

func Benchmark_RunTxn_commit_oneDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})

	logic := func(dax sabi.Dax) sabi.Err { return sabi.Ok() }

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

	logic := func(dax sabi.Dax) sabi.Err { return sabi.Ok() }

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
	logic := func(dax sabi.Dax) sabi.Err { return sabi.NewErr(Fail{}) }

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
	logic := func(dax sabi.Dax) sabi.Err { return sabi.NewErr(Fail{}) }

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}
