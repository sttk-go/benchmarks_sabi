# Benchmark of Dax in sabi v0.4.0

## orderedmap_isasync_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_isasync_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	35033134	        34.22 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 8240978	       142.3 ns/op	     320 B/op	       3 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 3072844	       386.7 ns/op	     576 B/op	       7 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2437 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	39872698	        28.56 ns/op	      16 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	22166484	        46.98 ns/op	      16 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2436 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	452908680	         2.611 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	140400019	         8.466 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	140771281	         9.048 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	34521769	        33.39 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 7956842	       158.6 ns/op	     320 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 2808211	       462.3 ns/op	     576 B/op	       7 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	34246527	        35.46 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	18957375	        59.77 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 8179860	       146.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	32823846	        34.67 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7606563	       158.8 ns/op	     320 B/op	       3 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 2706357	       437.7 ns/op	     576 B/op	       7 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	92615088	        12.84 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	17244031	        69.20 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	94832295	        12.57 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	16986939	        69.18 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 4441520	       267.6 ns/op	     384 B/op	       5 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1729261	       698.0 ns/op	     640 B/op	       9 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 5245804	       233.4 ns/op	     336 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1874678	       633.6 ns/op	     592 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_isasync_sync	37.799s
```

## orderedmap_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	35563122	        33.02 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 8105902	       143.2 ns/op	     320 B/op	       3 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 3095998	       376.6 ns/op	     576 B/op	       7 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2445 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	46583828	        25.67 ns/op	      16 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	30350209	        39.03 ns/op	      16 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2679 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	379551974	         3.097 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	135285717	         8.816 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	143757280	         8.298 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	36173857	        32.85 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 7903516	       147.2 ns/op	     320 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 2820202	       409.0 ns/op	     576 B/op	       7 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	35147030	        33.04 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	22729096	        52.68 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 8232686	       140.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	33605806	        34.89 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7754737	       155.2 ns/op	     320 B/op	       3 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 2809405	       426.3 ns/op	     576 B/op	       7 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	97839108	        12.22 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	17749440	        66.43 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	98939149	        12.14 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	17039484	        67.88 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 4465789	       268.7 ns/op	     384 B/op	       5 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1738520	       678.8 ns/op	     640 B/op	       9 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 5116464	       231.4 ns/op	     336 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1861424	       624.2 ns/op	     592 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_sync	37.569s
```
