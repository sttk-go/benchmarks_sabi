# Benchmark of Dax in sabi v0.4.0

## register_daxsrc_in_list_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/register_daxsrc_in_list_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	1000000000	         0.6196 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	31208605	        34.94 ns/op	      64 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 6774722	       172.6 ns/op	     320 B/op	       5 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2448 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	47246428	        25.21 ns/op	      16 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	31205007	        38.33 ns/op	      16 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2603 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	413212520	         2.823 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	125955319	         9.504 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	143901747	         8.311 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	1000000000	         0.4240 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	29990162	        39.92 ns/op	      64 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 5967441	       200.6 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	 3694419	       287.6 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	 4803978	       234.1 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 4986920	       227.7 ns/op	     320 B/op	       5 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	595094457	         1.963 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	23816497	        45.52 ns/op	      64 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 5453721	       220.1 ns/op	     320 B/op	       5 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	91552545	        12.64 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	16968477	        67.49 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	93865087	        12.59 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	17389602	        67.81 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 4805078	       243.2 ns/op	     336 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1700370	       694.9 ns/op	     592 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 4924094	       244.6 ns/op	     336 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1760666	       690.8 ns/op	     592 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/register_daxsrc_in_list_sync	35.315s
```

## orderedmap_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	34271924	        34.73 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 8441894	       143.5 ns/op	     320 B/op	       3 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 3085395	       383.9 ns/op	     576 B/op	       7 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2401 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	47393912	        25.10 ns/op	      16 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	30898572	        39.50 ns/op	      16 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2535 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	433777597	         2.758 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	134980707	         8.907 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	135506454	         8.806 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	33809064	        33.48 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 8151013	       145.1 ns/op	     320 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 2966322	       407.0 ns/op	     576 B/op	       7 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	35402526	        33.30 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	23104022	        52.09 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 8378400	       141.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	33451498	        35.45 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7715151	       153.6 ns/op	     320 B/op	       3 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 2809344	       421.7 ns/op	     576 B/op	       7 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	93539412	        12.74 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	17438576	        68.27 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	95645934	        12.53 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	17311896	        69.60 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 5140606	       229.4 ns/op	     336 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1867072	       643.2 ns/op	     592 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 5260694	       229.1 ns/op	     336 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1895728	       636.4 ns/op	     592 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_sync	37.808s
```
