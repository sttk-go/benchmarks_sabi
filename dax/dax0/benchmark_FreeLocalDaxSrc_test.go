package dax0_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/dax0"
	"testing"
)

func Benchmark_FreeLocalDaxSrc_oneDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddLocalDaxSrcForTest(base, "foo", FooDaxSrc{})
		base.FreeLocalDaxSrc("foo")
	}
	b.StopTimer()
}

func Benchmark_FreeLocalDaxSrc_fiveDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddLocalDaxSrcForTest(base, "foo", FooDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "bar", BarDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "baz", BazDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "qux", QuxDaxSrc{})
		sabi.AddLocalDaxSrcForTest(base, "corge", CorgeDaxSrc{})

		base.FreeLocalDaxSrc("foo")
		base.FreeLocalDaxSrc("bar")
		base.FreeLocalDaxSrc("baz")
		base.FreeLocalDaxSrc("qux")
		base.FreeLocalDaxSrc("corge")
	}
	b.StopTimer()
}
