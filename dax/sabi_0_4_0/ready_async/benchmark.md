# Benchmark of Dax in sabi v0.4.0

## adds Ready method to DaxSrc and adds Committed method to DaxConn

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/ready_async
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12*       	305621150	         4.057 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12*      	41175874	        28.43 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12*    	 2943505	       405.6 ns/op	      96 B/op	       4 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12*   	  738932	      1552 ns/op	     224 B/op	      12 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12*     	27681424	        42.27 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12*    	10017271	       118.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        10.25 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12*      	216622744	         5.479 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12*     	34958104	        34.92 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12*       	35161506	        33.34 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12*      	 6383206	       181.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12      	90479572	        13.06 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12     	17214297	        69.28 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12       	90649598	        13.05 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12      	17228955	        69.35 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12*         	 9773616	       121.5 ns/op	     112 B/op	       3 allocs/op
Benchmark_RunTxn_commit_fiveDs-12*        	 9746608	       120.8 ns/op	     112 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_oneDs-12*       	12937051	        85.05 ns/op	      64 B/op	       2 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12*      	14112696	        83.65 ns/op	      64 B/op	       2 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/ready_async	24.870s
```

## unlock

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/unlock
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12*       	295736892	         3.898 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12*      	41263272	        28.68 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12     	 1826632	       652.8 ns/op	     208 B/op	       3 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12    	  389084	      3091 ns/op	     464 B/op	       7 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12*      	 1748542	       681.5 ns/op	      64 B/op	       3 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12*     	  526398	      2306 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        10.30 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12*       	224207263	         5.360 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12*      	31958334	        37.42 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12*        	37023637	        32.06 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12*       	 6507991	       181.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12*      	100000000	        11.10 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12*     	18984116	        61.04 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12*       	100000000	        11.07 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12*      	19559186	        61.76 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12          	 7481137	       159.1 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12         	 7469710	       158.9 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12        	10576123	       113.5 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12       	10375509	       115.2 ns/op	      80 B/op	       3 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0/unlock	26.964s
```

## original

```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12        	68692728	        17.32 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12       	12623794	        91.12 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12     	 1825360	       652.2 ns/op	     208 B/op	       3 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12    	  382323	      3078 ns/op	     464 B/op	       7 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12      	 1736794	       683.8 ns/op	      64 B/op	       3 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12     	  505197	      2363 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        10.25 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12       	59895855	        19.59 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12      	11659152	       102.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12        	29426410	        40.37 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12       	 5486424	       215.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12      	100000000	        12.09 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12     	18126808	        64.92 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12       	98655546	        12.04 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12      	18221239	        65.23 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12          	 7393803	       159.8 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12         	 7321879	       159.7 ns/op	     208 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_oneDs-12        	10374529	       115.5 ns/op	      80 B/op	       3 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12       	10530865	       114.0 ns/op	      80 B/op	       3 allocs/op
PASS
ok  	github.com/sttk-go/benchmarks_sabi/dax/sabi_0_4_0	26.154s
```
