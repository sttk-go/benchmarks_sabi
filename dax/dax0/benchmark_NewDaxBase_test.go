package dax0_test

import (
  sabi "github.com/sttk-go/benchmarks_sabi/dax/dax0"
  "testing"
)

func Benchmark_NewDaxBase(b *testing.B) {
  unused := func(a any) {}

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    base := sabi.NewDaxBase()
    unused(base)
  }
  b.StopTimer()
}
