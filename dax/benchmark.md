# Dax

## dax0

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/dax0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12           	64613872	        18.66 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	10845070	       107.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 6524212	       184.1 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 3179898	       372.3 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	19838622	        61.34 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 3436761	       354.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	82861723	        14.03 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	 1000000	      1009 ns/op	     704 B/op	      12 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	83642322	        14.26 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	 1000000	      1003 ns/op	     704 B/op	      12 allocs/op
Benchmark_NewDaxBase-12                      	95604007	        12.19 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 7617122	       160.5 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 7590507	       158.7 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	10296888	       116.8 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	10205098	       117.8 ns/op	      80 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	58125292	        20.11 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 9841222	       121.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	 1786282	       663.5 ns/op	      64 B/op	       3 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_fiveDs-12    	  537992	      2287 ns/op	     256 B/op	      11 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 1789880	       660.3 ns/op	     208 B/op	       3 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  376281	      3144 ns/op	     464 B/op	       7 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/dax0	29.171s
```

## dax1_noLock

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/dax1_noLock
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12           	289389169	         4.075 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	24760960	        48.04 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7012440	       168.5 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 3564448	       334.8 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	29733372	        39.65 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 4580565	       260.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	100000000	        11.93 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	 1000000	      1032 ns/op	     832 B/op	      12 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	100000000	        11.83 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	 1174717	      1000 ns/op	     832 B/op	      12 allocs/op
Benchmark_NewDaxBase-12                      	93036042	        12.56 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 7329841	       161.1 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 7588900	       168.6 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 9881638	       113.4 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	10479270	       115.1 ns/op	      80 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	216420318	         5.513 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	21408070	        57.50 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	 1745301	       680.4 ns/op	      64 B/op	       3 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_fiveDs-12    	  517816	      2348 ns/op	     256 B/op	      11 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 1769203	       675.4 ns/op	     208 B/op	       3 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  377104	      3135 ns/op	     464 B/op	       7 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/dax1_noLock	29.992s
```
