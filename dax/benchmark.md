# Dax

## dax0

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/dax0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12           	68327784	        17.35 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	11144739	       106.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7110122	       164.6 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 4518154	       258.3 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	28766413	        40.79 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 5246419	       222.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	88964138	        13.91 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	11940345	        98.01 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	89022250	        13.65 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	11623515	       103.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	100000000	        10.50 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 7763695	       153.8 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 7755943	       155.7 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	10007840	       115.0 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	10433086	       113.6 ns/op	      80 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	61266128	        19.63 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	10027503	       118.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	 1820577	       658.5 ns/op	      64 B/op	       3 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_fiveDs-12    	  542523	      2252 ns/op	     256 B/op	      11 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 1822387	       659.2 ns/op	     208 B/op	       3 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  359163	      3118 ns/op	     464 B/op	       7 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/dax0	29.185s
```

## dax1_noLock

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/dax1_noLock
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12           	282509034	         4.069 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	24826694	        47.98 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7410186	       160.7 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 4656500	       256.3 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	36023348	        32.61 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 6486324	       183.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	100000000	        11.38 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	13158900	        85.30 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	100000000	        11.17 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	14010847	        83.68 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	100000000	        10.23 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 7558574	       158.1 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 7637180	       157.4 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 9841532	       118.9 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	10454887	       114.7 ns/op	      80 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	218353562	         5.467 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	21206170	        56.91 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	 1801458	       665.5 ns/op	      64 B/op	       3 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_fiveDs-12    	  531367	      2247 ns/op	     256 B/op	      11 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 1790932	       667.8 ns/op	     208 B/op	       3 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  379215	      3107 ns/op	     464 B/op	       7 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/dax1_noLock	28.782s
```

## dax2_ready_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/dax2_ready_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12           	299321118	         3.858 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	24744036	        48.40 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7203262	       160.2 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 4669128	       258.2 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	36087945	        32.17 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 6402082	       184.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	100000000	        11.23 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	13939212	        86.74 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	100000000	        11.26 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	14192558	        83.85 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	100000000	        10.33 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 7379130	       158.5 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 7648110	       156.1 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	10483134	       112.4 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 9866866	       111.4 ns/op	      80 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	182831946	         6.494 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	20417871	        58.42 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	33851185	        32.96 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_fiveDs-12    	19858430	        59.41 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 8088582	       138.4 ns/op	      64 B/op	       2 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	 5517650	       218.9 ns/op	      64 B/op	       2 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/dax2_ready_sync	27.808s
```

## dax3_ready_async

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/dax3_ready_async
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12           	305125531	         3.846 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	24690583	        48.18 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7261178	       162.0 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 4740766	       257.5 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	34882328	        32.21 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 6381190	       187.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	100000000	        11.14 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	13906411	        85.28 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	100000000	        11.07 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	14153684	        83.46 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	100000000	        10.29 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 7333154	       157.9 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 7554034	       158.6 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	10346524	       114.4 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	10522440	       113.8 ns/op	      80 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	175636461	         6.556 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	20333757	        58.40 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	37251141	        32.17 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_fiveDs-12    	20153490	        60.16 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 2923024	       408.2 ns/op	      96 B/op	       4 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  679399	      1567 ns/op	     224 B/op	      12 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/dax3_ready_async	27.734s
```
