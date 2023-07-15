# Benchmark of Dax in sabi v0.4.0

## register_daxsrc_in_list_async

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/register_daxsrc_in_list_async
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	1000000000	         0.6262 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	33841669	        38.09 ns/op	      64 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 6867734	       173.5 ns/op	     320 B/op	       5 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2413 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 2154201	       545.6 ns/op	      32 B/op	       2 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  599134	      2004 ns/op	      96 B/op	       6 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2399 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	496443106	         2.383 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	140175253	         8.343 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	100000000	        11.96 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	1000000000	         0.4932 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	23201197	        51.01 ns/op	      64 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 4443690	       268.7 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	 5412044	       221.0 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	 5139154	       211.7 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 5304961	       222.9 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	548023576	         2.165 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	17775276	        58.80 ns/op	      64 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 4130972	       286.0 ns/op	     320 B/op	       5 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	88731859	        12.89 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	17123137	        69.58 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	91963281	        12.93 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	16933298	        69.72 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 1350170	       875.1 ns/op	     336 B/op	       5 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	  394482	      3009 ns/op	     592 B/op	      13 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 1374561	       869.4 ns/op	     336 B/op	       5 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	  399840	      2921 ns/op	     592 B/op	      13 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/register_daxsrc_in_list_async	39.262s
```

## orderedmap_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_async
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	35704796	        33.73 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 8147667	       143.0 ns/op	     320 B/op	       3 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 3140148	       383.8 ns/op	     576 B/op	       7 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2499 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 2041052	       583.7 ns/op	      32 B/op	       2 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  574429	      2095 ns/op	      96 B/op	       6 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2508 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	404499268	         2.970 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	131079636	         9.056 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	139176177	         8.626 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	33647347	        33.89 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 7896608	       153.1 ns/op	     320 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 2776054	       430.1 ns/op	     576 B/op	       7 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	35732656	        33.35 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	21440049	        55.48 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 7732389	       153.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	33168475	        35.30 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7646227	       155.3 ns/op	     320 B/op	       3 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 2722699	       438.3 ns/op	     576 B/op	       7 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	93562012	        12.70 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	16817739	        70.06 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	96053512	        12.73 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	17014132	        68.63 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 1316002	       912.9 ns/op	     352 B/op	       5 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	  361906	      3286 ns/op	     672 B/op	      13 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 1299494	       914.4 ns/op	     352 B/op	       5 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	  363655	      3252 ns/op	     672 B/op	      13 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_async	38.514s
```
