# notify

## sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/notify/v0_4_0/sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddSyncErrHandler-12    	16173799	        62.14 ns/op	      16 B/op	       1 allocs/op
Benchmark_ok-12                   	1000000000	         0.2375 ns/op	       0 B/op	       0 allocs/op
Benchmark_emptyReason-12          	311186954	         3.786 ns/op	       0 B/op	       0 allocs/op
Benchmark_oneFieldReason-12       	34670158	        40.06 ns/op	      16 B/op	       1 allocs/op
Benchmark_fiveFieldReason-12      	28580348	        71.29 ns/op	      64 B/op	       1 allocs/op
Benchmark_havingCauseReason-12    	36939536	        41.96 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/notify/v0_4_0/sync	23.338s
```

## async

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/notify/v0_4_0/async
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddAsyncErrHandler-12    	15707230	        63.76 ns/op	      16 B/op	       1 allocs/op
Benchmark_ok-12                    	1000000000	         0.2487 ns/op	       0 B/op	       0 allocs/op
Benchmark_emptyReason-12           	307617741	         3.900 ns/op	       0 B/op	       0 allocs/op
Benchmark_oneFieldReason-12        	33746924	        40.14 ns/op	      16 B/op	       1 allocs/op
Benchmark_fiveFieldReason-12       	26183215	        72.16 ns/op	      64 B/op	       1 allocs/op
Benchmark_havingCauseReason-12     	35190273	        44.38 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/notify/v0_4_0/async	23.077s
```
