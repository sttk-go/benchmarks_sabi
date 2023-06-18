package unlock_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/sabi_0_4_0/unlock"
)

func Benchmark_SetupLocalDaxSrc_oneDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})
	}
	b.StopTimer()
}

func Benchmark_SetupLocalDaxSrc_fiveDs(b *testing.B) {
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base.SetUpLocalDaxSrc("cliargs", FooDaxSrc{})
		base.SetUpLocalDaxSrc("database", FooDaxSrc{})
		base.SetUpLocalDaxSrc("pubsub", FooDaxSrc{})
		base.SetUpLocalDaxSrc("json", FooDaxSrc{})
		base.SetUpLocalDaxSrc("env", FooDaxSrc{})
	}
	b.StopTimer()
}
