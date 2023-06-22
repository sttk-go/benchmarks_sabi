# Benchmark of Dax in sabi v0.4.0

## ready_async

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/ready_async
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	693328550	         1.644 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	43126464	        27.85 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 7207584	       160.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2427 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 1646998	       722.5 ns/op	      80 B/op	       3 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  546970	      2194 ns/op	     144 B/op	       7 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2426 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	37866961	        31.55 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	20327539	        58.18 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	100000000	        10.18 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	675605257	         1.742 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 1950352	       610.5 ns/op	      32 B/op	       2 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	  379138	      3170 ns/op	     160 B/op	      10 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	20882013	        55.16 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	14206041	        80.64 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 6299506	       183.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	30939840	        37.21 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 1467603	       815.5 ns/op	     368 B/op	       4 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	  350440	      3318 ns/op	     496 B/op	      12 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	96622316	        12.36 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	16763215	        69.18 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	98396222	        12.23 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	16994620	        68.05 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	  997198	      1091 ns/op	     416 B/op	       5 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	  372792	      3114 ns/op	     480 B/op	       9 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 1243606	       962.9 ns/op	     368 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	  406794	      2811 ns/op	     432 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/ready_async	34.510s
```

## unlock

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/unlock
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	672017583	         1.642 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	43227232	        27.88 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 7239038	       164.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2383 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 1870320	       638.6 ns/op	     208 B/op	       3 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  397155	      2990 ns/op	     464 B/op	       7 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2398 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	 1815012	       654.2 ns/op	      64 B/op	       3 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	  551445	      2213 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                      	100000000	        10.23 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	714590703	         1.650 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	38499928	        31.28 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 6783117	       176.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	21275480	        55.71 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	14522481	        80.87 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 6100206	       180.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	30847309	        37.92 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 6794308	       173.3 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 3930520	       305.8 ns/op	     336 B/op	       2 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	95786227	        11.85 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	16361152	        69.76 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	95812850	        12.25 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	16776498	        68.98 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	  713598	      1568 ns/op	     608 B/op	       8 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	  194941	      5914 ns/op	    1056 B/op	      20 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	  720855	      1576 ns/op	     464 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	  237524	      4838 ns/op	     848 B/op	      24 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/unlock	33.201s
```
