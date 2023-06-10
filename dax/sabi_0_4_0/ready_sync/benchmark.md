# Benchmark of Dax in sabi v0.4.0

## adds Ready method to DaxSrc and adds Committed method to DaxConn

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/ready_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12*       	304575050	         3.958 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12*      	42806258	        28.68 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12*    	 8926680	       134.0 ns/op	      64 B/op	       2 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12*   	 5703813	       209.8 ns/op	      64 B/op	       2 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12*     	27070966	        43.79 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12*    	 9673358	       122.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        10.38 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12*      	208599316	         5.703 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12*     	35392882	        34.95 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12*       	36392841	        33.91 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12*      	 6136545	       196.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12      	97540380	        12.32 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12     	16788832	        71.30 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12       	91462452	        12.45 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12      	17930511	        65.62 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12*         	 9316815	       120.1 ns/op	     112 B/op	       3 allocs/op
Benchmark_RunTxn_commit_fiveDs-12*        	 9962514	       120.6 ns/op	     112 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_oneDs-12*       	14008962	        86.04 ns/op	      64 B/op	       2 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12*      	14468781	        82.94 ns/op	      64 B/op	       2 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/ready_sync	25.077s
```

## unlock

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/unlock
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12*       	298354960	         3.889 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12*      	41970232	        28.64 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12     	 1834860	       651.0 ns/op	     208 B/op	       3 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12    	  369907	      3085 ns/op	     464 B/op	       7 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12*     	 1746042	       688.5 ns/op	      64 B/op	       3 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12*    	  518077	      2309 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        10.30 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12*      	223531206	         5.326 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12*     	32002276	        38.11 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12*       	36232384	        33.07 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12*      	 6566602	       182.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12*     	100000000	        11.21 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12*    	19572531	        61.68 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12*      	100000000	        11.26 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12*     	19567251	        61.53 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12          	 7450989	       160.2 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12         	 7299812	       160.0 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12        	 9877399	       114.3 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12       	10342538	       113.9 ns/op	      80 B/op	       3 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/unlock	26.006s
```

## original

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12        	68305056	        17.30 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12       	13019295	        90.82 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12     	 1839874	       651.4 ns/op	     208 B/op	       3 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12    	  381088	      3078 ns/op	     464 B/op	       7 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12      	 1749529	       684.7 ns/op	      64 B/op	       3 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12     	  523287	      2313 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        10.32 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12       	59306278	        20.57 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12      	11462022	       104.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12        	28161942	        40.59 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12       	 5328586	       216.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12      	97057099	        12.08 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12     	18036222	        64.89 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12       	99439418	        12.05 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12      	18075672	        65.24 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12          	 7343053	       161.2 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12         	 7476615	       160.4 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12        	 9988776	       113.4 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12       	10375681	       113.2 ns/op	      80 B/op	       3 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0	26.578s
```
