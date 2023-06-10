# Benchmark of Dax in sabi v0.4.0

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12        	66776688	        17.43 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12       	13165652	        90.79 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12     	 1833920	       647.0 ns/op	     208 B/op	       3 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12    	  386816	      3052 ns/op	     464 B/op	       7 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12      	 1748050	       680.1 ns/op	      64 B/op	       3 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12     	  521553	      2296 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        10.26 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12       	53398462	        19.54 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12      	11542224	       102.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12        	29178195	        40.41 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12       	 5469159	       216.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12      	99699745	        12.05 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12     	18443382	        64.73 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12       	99429457	        12.04 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12      	17864029	        65.48 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12          	 7345782	       157.7 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12         	 7400504	       158.7 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12        	10483096	       113.9 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12       	10199563	       112.4 ns/op	      80 B/op	       3 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0	24.968s

```
