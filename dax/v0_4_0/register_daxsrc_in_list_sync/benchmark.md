# Benchmark of Dax in sabi v0.4.0

## register_daxsrc_in_list_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/register_daxsrc_in_list_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	1000000000	         0.5871 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	33565353	        34.04 ns/op	      64 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 7058150	       169.6 ns/op	     320 B/op	       5 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2391 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	45852760	        24.68 ns/op	      16 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	31477484	        36.98 ns/op	      16 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2379 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	636490131	         1.874 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	145778431	         8.201 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	99466717	        11.98 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	1000000000	         0.4158 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	23826708	        50.49 ns/op	      64 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 4495912	       267.9 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	 5318566	       220.2 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	 5626162	       221.4 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 5414763	       223.8 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	602804790	         1.966 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	19397810	        56.03 ns/op	      64 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 4164927	       282.7 ns/op	     320 B/op	       5 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	96867279	        11.99 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	17577278	        67.22 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	100000000	        12.02 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	17820243	        66.13 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 5320132	       218.0 ns/op	     320 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 2110636	       560.6 ns/op	     512 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 5353162	       219.8 ns/op	     320 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 2198718	       545.9 ns/op	     512 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/register_daxsrc_in_list_sync	37.991s
```

## orderedmap_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	33498493	        35.35 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 8367165	       142.9 ns/op	     320 B/op	       3 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 3077402	       384.6 ns/op	     576 B/op	       7 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2495 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	45482864	        25.70 ns/op	      16 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	30532304	        38.87 ns/op	      16 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2483 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	438045627	         2.710 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	138165579	         8.676 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	137596050	         8.733 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	35559105	        33.56 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 8112730	       146.5 ns/op	     320 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 2945577	       407.7 ns/op	     576 B/op	       7 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	34949241	        34.55 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	22308955	        53.89 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 8091321	       145.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	33568606	        35.68 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7783567	       152.0 ns/op	     320 B/op	       3 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 2826013	       421.3 ns/op	     576 B/op	       7 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	95199501	        12.75 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	16600123	        71.19 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	91494823	        12.82 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	16461879	        71.38 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 5193310	       227.7 ns/op	     336 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1845044	       641.1 ns/op	     592 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 5232573	       229.4 ns/op	     336 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1878931	       637.2 ns/op	     592 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_sync	37.733s
```
