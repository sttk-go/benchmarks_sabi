# Benchmark of Dax in sabi v0.4.0

## adds Ready method to DaxSrc and adds Committed method to DaxConn

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/ready_async
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12        	298844008	         3.857 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12       	42650397	        28.20 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12     	 2885840	       409.1 ns/op	      96 B/op	       4 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12    	  771073	      1563 ns/op	     224 B/op	      12 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12      	27707588	        41.37 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12     	 9514581	       123.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        10.24 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12       	216004623	         5.462 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12      	33953202	        34.70 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12        	37423992	        31.76 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12       	 6409270	       183.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12      	89400211	        12.96 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12     	17127253	        70.18 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12       	90436812	        13.43 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12      	16685865	        69.22 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12          	 1781894	       667.5 ns/op	     432 B/op	       6 allocs/op
Benchmark_RunTxn_commit_fiveDs-12         	  586051	      2031 ns/op	     560 B/op	      14 allocs/op
Benchmark_RunTxn_rollback_oneDs-12        	 2011890	       597.8 ns/op	     384 B/op	       5 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12       	  600097	      1932 ns/op	     512 B/op	      13 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/ready_async	27.006s
```

## unlock

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/unlock
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12        	297490818	         3.868 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12       	42367635	        28.44 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12     	 1827039	       651.6 ns/op	     208 B/op	       3 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12    	  384792	      3085 ns/op	     464 B/op	       7 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12      	 1750386	       687.3 ns/op	      64 B/op	       3 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12     	  520966	      2305 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        10.34 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12       	224806663	         5.312 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12      	32021836	        37.10 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12        	37795502	        32.36 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12       	 6500716	       181.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12      	100000000	        11.17 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12     	18155869	        64.54 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12       	100000000	        11.16 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12      	18127329	        64.45 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12          	  738268	      1594 ns/op	     608 B/op	       8 allocs/op
Benchmark_RunTxn_commit_fiveDs-12         	  191364	      5997 ns/op	    1056 B/op	      20 allocs/op
Benchmark_RunTxn_rollback_oneDs-12        	  732496	      1594 ns/op	     464 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12       	  234860	      4922 ns/op	     848 B/op	      24 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/unlock	25.409s
```

## original

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12        	68298651	        17.34 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12       	13002093	        91.55 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12     	 1835984	       657.4 ns/op	     208 B/op	       3 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12    	  364510	      3089 ns/op	     464 B/op	       7 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12      	 1684418	       698.7 ns/op	      64 B/op	       3 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12     	  516620	      2328 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        10.36 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12       	59602483	        19.64 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12      	11637241	       102.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12        	29315618	        40.42 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12       	 5477397	       216.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12      	98950856	        12.35 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12     	17967630	        65.25 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12       	99228013	        12.05 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12      	18360937	        65.17 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12          	  694435	      1599 ns/op	     608 B/op	       8 allocs/op
Benchmark_RunTxn_commit_fiveDs-12         	  192574	      6071 ns/op	    1056 B/op	      20 allocs/op
Benchmark_RunTxn_rollback_oneDs-12        	  730015	      1607 ns/op	     464 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12       	  233893	      4991 ns/op	     848 B/op	      24 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0	25.667s
```
