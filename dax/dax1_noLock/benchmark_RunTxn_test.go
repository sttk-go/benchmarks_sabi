package dax1_noLock_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/dax1_noLock"
	"testing"
)

func Benchmark_RunTxn_commit_oneDs(b *testing.B) {
	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("foo", FooDaxSrc{})

	logic := func(dax sabi.Dax) sabi.Err { return sabi.Ok() }

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func Benchmark_RunTxn_commit_fiveDs(b *testing.B) {
	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("foo", FooDaxSrc{})
	base.SetUpLocalDaxSrc("foo", FooDaxSrc{})
	base.SetUpLocalDaxSrc("foo", FooDaxSrc{})
	base.SetUpLocalDaxSrc("foo", FooDaxSrc{})
	base.SetUpLocalDaxSrc("foo", FooDaxSrc{})

	logic := func(dax sabi.Dax) sabi.Err { return sabi.Ok() }

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func Benchmark_RunTxn_rollback_oneDs(b *testing.B) {
	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("foo", FooDaxSrc{})

	type Fail struct{}
	logic := func(dax sabi.Dax) sabi.Err { return sabi.NewErr(Fail{}) }

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func Benchmark_RunTxn_rollback_fiveDs(b *testing.B) {
	base := sabi.NewDaxBase()
	base.SetUpLocalDaxSrc("foo", FooDaxSrc{})
	base.SetUpLocalDaxSrc("foo", FooDaxSrc{})
	base.SetUpLocalDaxSrc("foo", FooDaxSrc{})
	base.SetUpLocalDaxSrc("foo", FooDaxSrc{})
	base.SetUpLocalDaxSrc("foo", FooDaxSrc{})

	type Fail struct{}
	logic := func(dax sabi.Dax) sabi.Err { return sabi.NewErr(Fail{}) }

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}
