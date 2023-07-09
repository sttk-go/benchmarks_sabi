# Benchmark of Dax in sabi v0.4.0

## register_daxsrc_in_list_async

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/register_daxsrc_in_list_async
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	1000000000	         0.6040 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	33542286	        34.02 ns/op	      64 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 6980875	       172.1 ns/op	     320 B/op	       5 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2421 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 2112610	       564.6 ns/op	      32 B/op	       2 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  599376	      2004 ns/op	      96 B/op	       6 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2431 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	624127725	         1.912 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	140234799	         8.588 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	143336984	         8.309 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	1000000000	         0.4785 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	27805158	        41.48 ns/op	      64 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 5680309	       209.3 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	 4197866	       240.9 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	 4957056	       222.0 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 5232993	       226.0 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	600020328	         1.968 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	23741653	        45.09 ns/op	      64 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 5237934	       223.1 ns/op	     320 B/op	       5 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	88406509	        13.29 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	16640314	        71.49 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	88041660	        13.38 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	16453153	        72.44 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 1285921	       927.5 ns/op	     352 B/op	       5 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	  352708	      3333 ns/op	     672 B/op	      13 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 1294119	       923.6 ns/op	     352 B/op	       5 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	  349576	      3353 ns/op	     672 B/op	      13 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/register_daxsrc_in_list_async	38.542s
```

## orderedmap_async

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_async
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	35613368	        33.38 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 8305567	       141.6 ns/op	     320 B/op	       3 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 3147579	       378.3 ns/op	     576 B/op	       7 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2425 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 2140351	       558.7 ns/op	      32 B/op	       2 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  609810	      2016 ns/op	      96 B/op	       6 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2426 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	417414655	         2.852 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	135677966	         8.786 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	144195602	         8.290 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	34917196	        33.22 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 7841844	       153.4 ns/op	     320 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 2803820	       428.0 ns/op	     576 B/op	       7 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	35665041	        32.53 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	22333728	        54.50 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 8033294	       147.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	33183504	        35.13 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7655464	       156.7 ns/op	     320 B/op	       3 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 2746215	       436.8 ns/op	     576 B/op	       7 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	94771929	        12.14 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	17845200	        67.46 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	96595460	        12.16 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	17769390	        66.02 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 1350398	       887.4 ns/op	     352 B/op	       5 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	  366168	      3202 ns/op	     672 B/op	      13 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 1352938	       880.5 ns/op	     352 B/op	       5 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	  377204	      3160 ns/op	     672 B/op	      13 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_async	38.205s
```
