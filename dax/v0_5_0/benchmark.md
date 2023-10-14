# Benchmark of Dax in sabi v0.5.0

> - `BenchmarkDax_____*` is v0.5.0.
> - `BenchmarkDaxPrev_*` is v0.4.0.

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_5_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkDax_____AddGlobalDaxSrc_zeroDs-12        	1000000000	         0.7503 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_AddGlobalDaxSrc_zeroDs-12        	33376377	        34.42 ns/op	      48 B/op	       1 allocs/op
BenchmarkDax_____AddGlobalDaxSrc_oneDs-12         	31556976	        40.48 ns/op	      64 B/op	       1 allocs/op
BenchmarkDaxPrev_AddGlobalDaxSrc_oneDs-12         	 9437352	       126.8 ns/op	     336 B/op	       2 allocs/op
BenchmarkDax_____AddGlobalDaxSrc_fiveDs-12        	 6658290	       176.3 ns/op	     320 B/op	       5 allocs/op
BenchmarkDaxPrev_AddGlobalDaxSrc_fiveDs-12        	 4773223	       247.9 ns/op	     336 B/op	       2 allocs/op
BenchmarkDax_____SetupGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2488 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_SetupGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2488 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____SetupGlobalDaxSrcs_oneDs-12      	28557936	        41.87 ns/op	      64 B/op	       1 allocs/op
BenchmarkDaxPrev_SetupGlobalDaxSrcs_oneDs-12      	 1818165	       655.1 ns/op	     208 B/op	       3 allocs/op
BenchmarkDax_____SetupGlobalDaxSrcs_fiveDs-12     	22314506	        52.35 ns/op	      64 B/op	       1 allocs/op
BenchmarkDaxPrev_SetupGlobalDaxSrcs_fiveDs-12     	  369990	      3073 ns/op	     464 B/op	       7 allocs/op
BenchmarkDax_____CloseGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2495 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_CloseGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2489 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____CloseGlobalDaxSrcs_oneDs-12      	531284226	         2.248 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_CloseGlobalDaxSrcs_oneDs-12      	 1763954	       672.7 ns/op	      64 B/op	       3 allocs/op
BenchmarkDax_____CloseGlobalDaxSrcs_freeDs-12     	141112470	         8.420 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_CloseGlobalDaxSrcs_freeDs-12     	  503298	      2298 ns/op	     256 B/op	      11 allocs/op
BenchmarkDax_____NewDaxBase-12                    	77358645	        14.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_NewDaxBase-12                    	100000000	        10.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____AddLocalDaxSrc_zeroDs-12         	1000000000	         0.4986 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_AddLocalDaxSrc_zeroDs-12         	33787790	        34.61 ns/op	      48 B/op	       1 allocs/op
BenchmarkDax_____AddLocalDaxSrc_oneDs-12          	22835356	        52.77 ns/op	      64 B/op	       1 allocs/op
BenchmarkDaxPrev_AddLocalDaxSrc_oneDs-12          	 8896658	       131.2 ns/op	     336 B/op	       2 allocs/op
BenchmarkDax_____AddLocalDaxSrc_fiveDs-12         	 4419768	       274.2 ns/op	     320 B/op	       5 allocs/op
BenchmarkDaxPrev_AddLocalDaxSrc_fiveDs-12         	 4436720	       271.5 ns/op	     336 B/op	       2 allocs/op
BenchmarkDax_____FreeLocalDaxSrc_zeroDs-12        	 3580320	       293.5 ns/op	     320 B/op	       5 allocs/op
BenchmarkDaxPrev_FreeLocalDaxSrc_zeroDs-12        	19579417	        60.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____FreeLocalDaxSrc_oneDs-12         	 5488020	       264.2 ns/op	     320 B/op	       5 allocs/op
BenchmarkDaxPrev_FreeLocalDaxSrc_oneDs-12         	12052033	       102.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____FreeLocalDaxSrc_fiveDs-12        	 5085693	       394.3 ns/op	     320 B/op	       5 allocs/op
BenchmarkDaxPrev_FreeLocalDaxSrc_fiveDs-12        	 4706419	       237.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____FreeAllLocalDaxSrcs_zeroDs-12    	590861132	         2.309 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_FreeAllLocalDaxSrcs_zeroDs-12    	20013699	        61.14 ns/op	      48 B/op	       1 allocs/op
BenchmarkDax_____FreeAllLocalDaxSrcs_oneDs-12     	12038931	        86.26 ns/op	      64 B/op	       1 allocs/op
BenchmarkDaxPrev_FreeAllLocalDaxSrcs_oneDs-12     	 6137978	       192.1 ns/op	     336 B/op	       2 allocs/op
BenchmarkDax_____FreeAllLocalDaxSrcs_fiveDs-12    	 2997429	       398.9 ns/op	     320 B/op	       5 allocs/op
BenchmarkDaxPrev_FreeAllLocalDaxSrcs_fiveDs-12    	 3021103	       364.3 ns/op	     336 B/op	       2 allocs/op
BenchmarkDax_____GetDaxConn_global_oneDs-12       	92518327	        13.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_GetDaxConn_global_oneDs-12       	91300986	        13.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____GetDaxConn_global_fiveDs-12      	14840193	        80.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_GetDaxConn_global_fiveDs-12      	15539475	        76.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____GetDaxConn_local_oneDs-12        	83828145	        14.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_GetDaxConn_local_oneDs-12        	89615745	        12.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____GetDaxConn_local_fiveDs-12       	15002965	        84.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_GetDaxConn_local_fiveDs-12       	12021766	        85.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____RunTxn_commit_oneDs-12           	 8502117	       124.0 ns/op	      64 B/op	       1 allocs/op
BenchmarkDaxPrev_RunTxn_commit_oneDs-12           	  690764	      1662 ns/op	     608 B/op	       8 allocs/op
BenchmarkDax_____RunTxn_commit_fiveDs-12          	 3373076	       381.3 ns/op	      64 B/op	       1 allocs/op
BenchmarkDaxPrev_RunTxn_commit_fiveDs-12          	  184630	      6479 ns/op	    1056 B/op	      20 allocs/op
BenchmarkDax_____RunTxn_rollback_oneDs-12         	 9477416	       145.3 ns/op	      64 B/op	       1 allocs/op
BenchmarkDaxPrev_RunTxn_rollback_oneDs-12         	  674092	      1794 ns/op	     464 B/op	       8 allocs/op
BenchmarkDax_____RunTxn_rollback_fiveDs-12        	 3284464	       372.9 ns/op	      64 B/op	       1 allocs/op
BenchmarkDaxPrev_RunTxn_rollback_fiveDs-12        	  218294	      5875 ns/op	     848 B/op	      24 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_5_0	70.128s
```
