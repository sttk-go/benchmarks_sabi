# Benchmark of Dax in sabi v0.4.0

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12        	68064441	        17.31 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12       	13056586	        90.89 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12     	 1837280	       650.8 ns/op	     208 B/op	       3 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12    	  383868	      3136 ns/op	     464 B/op	       7 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12      	 1696023	       708.7 ns/op	      64 B/op	       3 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12     	  500344	      2587 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        11.30 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12       	58057123	        20.77 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12      	 9681686	       108.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12        	28310112	        43.29 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12       	 5167663	       228.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12      	98590233	        12.28 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12     	18525591	        65.28 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12       	98289774	        12.01 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12      	18265119	        64.96 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12          	  722646	      1590 ns/op	     608 B/op	       8 allocs/op
Benchmark_RunTxn_commit_fiveDs-12         	  191404	      6089 ns/op	    1056 B/op	      20 allocs/op
Benchmark_RunTxn_rollback_oneDs-12        	  720508	      1609 ns/op	     464 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12       	  224719	      5027 ns/op	     848 B/op	      24 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0	25.865s
```
