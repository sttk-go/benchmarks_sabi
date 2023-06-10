# Benchmark of Dax in sabi v0.4.0

## unlock

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/unlock
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12*        	299033228	         3.950 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12*       	40878951	        30.01 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12     	 1785477	       668.7 ns/op	     208 B/op	       3 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12    	  380688	      3129 ns/op	     464 B/op	       7 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12*      	 1704133	       701.0 ns/op	      64 B/op	       3 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12*     	  505640	      2412 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        10.57 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12*       	204593811	         5.489 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12*      	31246194	        39.59 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12*        	36237418	        33.56 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12*       	 6293496	       194.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12*      	100000000	        11.64 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12*     	18048814	        66.10 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12*       	100000000	        11.73 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12*      	18269988	        65.99 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12          	 7249626	       162.9 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12         	 7479223	       162.1 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12        	10126107	       115.9 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12       	10289503	       116.3 ns/op	      80 B/op	       3 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/unlock	26.370s
```

## original

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12        	68295711	        17.32 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12       	13099276	        90.72 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12     	 1827309	       649.0 ns/op	     208 B/op	       3 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12    	  388081	      3103 ns/op	     464 B/op	       7 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12      	 1722231	       686.8 ns/op	      64 B/op	       3 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12     	  522172	      2314 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        10.26 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12       	56734749	        19.56 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12      	11484742	       102.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12        	29622751	        40.40 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12       	 5454741	       217.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12      	99111391	        12.15 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12     	18087664	        64.89 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12       	97366408	        12.10 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12      	17838622	        65.34 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12          	 7214811	       161.0 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12         	 6213116	       163.0 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12        	10421632	       114.6 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12       	10495599	       114.1 ns/op	      80 B/op	       3 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0	26.039s
```
