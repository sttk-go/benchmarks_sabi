# Benchmark of Dax in sabi v0.4.0

## asyncgroup_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/asyncgroup_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	34090812	        34.90 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 8246877	       143.4 ns/op	     320 B/op	       3 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 3065690	       388.5 ns/op	     576 B/op	       7 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2449 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	16431454	        70.59 ns/op	      96 B/op	       2 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	10637293	       111.8 ns/op	     224 B/op	       2 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2439 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	357702726	         3.340 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	100000000	        10.01 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	142862359	         8.300 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	35067418	        33.40 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 5443998	       216.9 ns/op	     416 B/op	       5 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 1577400	       742.6 ns/op	    1056 B/op	      17 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	35224302	        33.31 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	22785189	        52.56 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 8239255	       144.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	32798256	        35.34 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 5280441	       225.0 ns/op	     416 B/op	       5 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 1517216	       777.5 ns/op	    1056 B/op	      17 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	96825682	        12.24 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	17641405	        67.51 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	98580651	        12.25 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	17305252	        68.80 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 4392880	       274.5 ns/op	     416 B/op	       5 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1678720	       715.4 ns/op	     800 B/op	       9 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 4942623	       242.0 ns/op	     384 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1888006	       630.9 ns/op	     640 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/asyncgroup_sync	37.621s
```

## orderedmap_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	35065758	        34.50 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 8048066	       146.1 ns/op	     320 B/op	       3 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 3097280	       387.0 ns/op	     576 B/op	       7 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2506 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	44978307	        25.84 ns/op	      16 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	29820260	        40.24 ns/op	      16 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2412 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	385634155	         3.087 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	135947154	         8.819 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	144207121	         8.513 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	36495001	        33.37 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 7814625	       150.9 ns/op	     320 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 2854320	       415.3 ns/op	     576 B/op	       7 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	34728174	        33.09 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	22413698	        52.44 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 8386680	       141.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	32952954	        35.44 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7713186	       156.7 ns/op	     320 B/op	       3 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 2751710	       431.8 ns/op	     576 B/op	       7 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	95433308	        12.22 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	17879406	        66.34 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	98523808	        12.18 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	17321272	        68.06 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 4544973	       266.5 ns/op	     384 B/op	       5 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1770678	       684.8 ns/op	     640 B/op	       9 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 5143808	       230.9 ns/op	     336 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1883058	       620.5 ns/op	     592 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_sync	37.757s
```
