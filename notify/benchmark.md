# notify

## sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/notify/sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddSyncErrHandler-12    	84087196	        14.95 ns/op	       0 B/op	       0 allocs/op
Benchmark_ok-12                   	1000000000	         0.2486 ns/op	       0 B/op	       0 allocs/op
Benchmark_emptyReason-12          	 1780969	       690.2 ns/op	     232 B/op	       2 allocs/op
Benchmark_oneFieldReason-12       	 1663237	       742.4 ns/op	     248 B/op	       3 allocs/op
Benchmark_fiveFieldReason-12      	 1550002	       767.6 ns/op	     296 B/op	       3 allocs/op
Benchmark_havingCauseReason-12    	 1455140	       797.2 ns/op	     248 B/op	       3 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/notify/sync	11.716s
```

## async

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/notify/async
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddAsyncErrHandler-12    	87209588	        13.57 ns/op	       0 B/op	       0 allocs/op
Benchmark_ok-12                    	1000000000	         0.2411 ns/op	       0 B/op	       0 allocs/op
Benchmark_emptyReason-12           	 1292944	       921.1 ns/op	     328 B/op	       3 allocs/op
Benchmark_oneFieldReason-12        	 1243851	       962.8 ns/op	     344 B/op	       4 allocs/op
Benchmark_fiveFieldReason-12       	 1226570	       979.6 ns/op	     392 B/op	       4 allocs/op
Benchmark_havingCauseReason-12     	 1235872	       969.8 ns/op	     344 B/op	       4 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/notify/async	11.222s
```
