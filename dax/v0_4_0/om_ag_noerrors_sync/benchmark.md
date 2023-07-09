# Benchmark of Dax in sabi v0.4.0

## om_ag_noerrors_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/om_ag_noerrors_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	34289904	        35.13 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 8102473	       143.1 ns/op	     320 B/op	       3 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 3067119	       390.1 ns/op	     576 B/op	       7 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2463 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	29770047	        38.08 ns/op	      48 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	24999752	        47.35 ns/op	      48 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2464 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	443134194	         2.687 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	137304121	         8.703 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	139895991	         8.525 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	35492070	        33.01 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 8075727	       148.3 ns/op	     320 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 2926333	       407.7 ns/op	     576 B/op	       7 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	35052055	        33.62 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	21728498	        53.74 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 8228388	       143.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	32980230	        36.14 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7618346	       157.9 ns/op	     320 B/op	       3 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 2716894	       461.8 ns/op	     576 B/op	       7 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	88558029	        13.58 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	13644326	        76.37 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	91279596	        13.22 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	13471392	        77.29 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 4936533	       241.3 ns/op	     368 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1831201	       653.6 ns/op	     624 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 4859222	       246.4 ns/op	     368 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1858632	       651.7 ns/op	     624 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/om_ag_noerrors_sync	38.454s
```

## orderedmap_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	30806689	        35.03 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 7066903	       145.9 ns/op	     320 B/op	       3 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 3141939	       385.3 ns/op	     576 B/op	       7 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2489 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	45185708	        26.18 ns/op	      16 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	27613196	        40.18 ns/op	      16 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2464 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	376559170	         3.198 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	130108695	         9.248 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	137716935	         8.710 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	36307406	        34.38 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 7828033	       152.3 ns/op	     320 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 2858307	       420.6 ns/op	     576 B/op	       7 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	34313535	        34.65 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	22170286	        54.27 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 8013774	       148.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	34339953	        35.66 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7690497	       157.5 ns/op	     320 B/op	       3 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 2773912	       444.4 ns/op	     576 B/op	       7 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	92467653	        13.33 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	15603603	        71.85 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	93096693	        13.31 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	16138250	        71.98 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 4416508	       316.8 ns/op	     384 B/op	       5 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1663382	       806.9 ns/op	     640 B/op	       9 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 5269154	       235.1 ns/op	     336 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1919974	       682.2 ns/op	     592 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_sync	37.422s
```
