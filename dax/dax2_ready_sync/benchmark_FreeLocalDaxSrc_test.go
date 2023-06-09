package dax2_ready_sync_test

import (
	sabi "github.com/sttk-go/benchmarks_sabi/dax/dax2_ready_sync"
	"testing"
)

func Benchmark_FreeLocalDaxSrc_oneDs(b *testing.B) {
  base := sabi.NewDaxBase()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    base.SetUpLocalDaxSrc("foo", FooDaxSrc{})
    base.FreeLocalDaxSrc("foo")
  }
  b.StopTimer()
}

func Benchmark_FreeLocalDaxSrc_fiveDs(b *testing.B) {
  base := sabi.NewDaxBase()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    base.SetUpLocalDaxSrc("foo", FooDaxSrc{})
    base.SetUpLocalDaxSrc("bar", BarDaxSrc{})
    base.SetUpLocalDaxSrc("baz", BazDaxSrc{})
    base.SetUpLocalDaxSrc("qux", QuxDaxSrc{})
    base.SetUpLocalDaxSrc("corge", CorgeDaxSrc{})

    base.FreeLocalDaxSrc("foo")
    base.FreeLocalDaxSrc("bar")
    base.FreeLocalDaxSrc("baz")
    base.FreeLocalDaxSrc("qux")
    base.FreeLocalDaxSrc("corge")
  }
  b.StopTimer()
}
