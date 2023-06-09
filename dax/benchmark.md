# Dax

## dax0

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/dax0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12           	68839777	        18.20 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	10994059	       109.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 6413575	       183.6 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 3221197	       383.2 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	19191723	        61.56 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 3348549	       349.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	84325184	        13.94 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	12988977	        94.18 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	88968690	        14.02 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	11985762	        97.24 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	83735706	        12.50 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 7670098	       158.4 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 7374903	       157.9 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	10028804	       116.4 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	10217492	       116.3 ns/op	      80 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	61125988	        20.05 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 9814810	       120.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	 1801696	       663.2 ns/op	      64 B/op	       3 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_fiveDs-12    	  509618	      2272 ns/op	     256 B/op	      11 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 1815336	       656.9 ns/op	     208 B/op	       3 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  386509	      3106 ns/op	     464 B/op	       7 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/dax0	30.170s
```

## dax1_noLock

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/dax1_noLock
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12           	270324130	         4.040 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	24704067	        47.77 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 6940434	       167.1 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 3511857	       337.4 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	29876404	        39.61 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 4578314	       259.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	100000000	        11.12 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	13947904	        84.07 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	100000000	        11.30 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	14033695	        83.87 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	97259928	        12.23 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 7435768	       159.5 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 7276369	       158.2 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	10149468	       112.9 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	10070470	       111.9 ns/op	      80 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	216730340	         5.490 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	21043538	        56.83 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	 1806780	       657.6 ns/op	      64 B/op	       3 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_fiveDs-12    	  544866	      2247 ns/op	     256 B/op	      11 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 1798779	       659.9 ns/op	     208 B/op	       3 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  380410	      3105 ns/op	     464 B/op	       7 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/dax1_noLock	28.963s
```

## dax2_ready_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/dax2_ready_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12           	301572358	         3.836 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	24525600	        48.05 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 5142834	       220.6 ns/op	     368 B/op	       4 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 2281500	       522.2 ns/op	     432 B/op	       8 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	12941109	        90.78 ns/op	      32 B/op	       2 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 2137970	       566.1 ns/op	     160 B/op	      10 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	100000000	        11.48 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	12883806	        90.20 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	100000000	        11.16 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	13507099	        87.72 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	99191546	        12.16 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 7422187	       155.5 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 7706353	       155.1 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	10367853	       112.4 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	10409209	       113.2 ns/op	      80 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	181616311	         6.571 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	20399294	        57.99 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	20955897	        55.48 ns/op	      16 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_fiveDs-12    	13546245	        87.33 ns/op	      16 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 8708875	       134.7 ns/op	      64 B/op	       2 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	 5494310	       214.1 ns/op	      64 B/op	       2 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/dax2_ready_sync	28.884s
```

## dax3_ready_async

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/dax3_ready_async
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12           	302904734	         3.869 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	24443527	        48.38 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 1578192	       757.5 ns/op	     432 B/op	       8 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	  366979	      3204 ns/op	     752 B/op	      28 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	 1912308	       631.6 ns/op	      96 B/op	       6 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	  365749	      3292 ns/op	     480 B/op	      30 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	100000000	        11.88 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	12935763	        90.46 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	100000000	        11.46 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	13298268	        89.80 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	93170037	        12.70 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 7480966	       160.2 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 7416972	       159.8 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	10458277	       115.7 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	10479346	       114.8 ns/op	      80 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	177390304	         7.026 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	19898253	        60.96 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	 3586550	       328.7 ns/op	      48 B/op	       3 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_fiveDs-12    	  862464	      1411 ns/op	     176 B/op	      11 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 2924139	       408.8 ns/op	      96 B/op	       4 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  764928	      1566 ns/op	     224 B/op	      12 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/dax3_ready_async	29.633s
```
