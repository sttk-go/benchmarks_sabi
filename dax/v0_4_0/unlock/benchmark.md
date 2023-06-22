# Benchmark of Dax in sabi v0.4.0

## unlock

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/unlock
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	672033044	         1.668 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	43176595	        28.36 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 7032249	       167.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2459 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 1825756	       656.2 ns/op	     208 B/op	       3 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  385932	      3054 ns/op	     464 B/op	       7 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2432 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	 1777165	       668.3 ns/op	      64 B/op	       3 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	  544554	      2249 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                      	100000000	        10.27 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	692771485	         1.687 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	37710522	        31.72 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 6442788	       178.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	20985402	        55.79 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	14420409	        82.18 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 6417178	       183.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	29625877	        38.07 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 6830809	       173.0 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 3891985	       310.5 ns/op	     336 B/op	       2 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	96295395	        11.99 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	16072542	        70.87 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	95701563	        12.39 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	16599384	        70.33 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	  732890	      1594 ns/op	     608 B/op	       8 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	  187446	      6086 ns/op	    1056 B/op	      20 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	  711324	      1610 ns/op	     464 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	  225082	      4939 ns/op	     848 B/op	      24 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/unlock	33.336s
```

## original

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	666196428	         1.662 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	27203445	        41.35 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 5018512	       235.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2409 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 1795323	       663.7 ns/op	     208 B/op	       3 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  370544	      3083 ns/op	     464 B/op	       7 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2442 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	 1802262	       665.2 ns/op	      64 B/op	       3 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	  541734	      2240 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                      	100000000	        10.30 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	665944494	         1.740 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	27511218	        41.85 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 4955572	       240.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	20918829	        55.80 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	12605403	        91.65 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 5319708	       220.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	24499591	        48.24 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 6249046	       189.6 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 3310654	       359.8 ns/op	     336 B/op	       2 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	96208657	        12.21 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	16658388	        69.87 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	95678970	        12.61 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	16406976	        69.73 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	  700404	      1603 ns/op	     608 B/op	       8 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	  184063	      6119 ns/op	    1056 B/op	      20 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	  723436	      1626 ns/op	     464 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	  226898	      4986 ns/op	     848 B/op	      24 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0	33.555s
```
