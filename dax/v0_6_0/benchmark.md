# Benchmark of Dax in sabi v0.6.0

> v0.6.0 is slightly slower than v0.5.0, but this is due to the use of `orderedmap` in v0.6.0.
> The use of `orderedmap` is necessary for controlling the order in `DaxSrc#Setup` and `DaxConn#Commit`
> (executing asynchronous process first).

> - `BenchmarkDax_____*` is v0.6.0.
> - `BenchmarkDaxPrev_*` is v0.5.0.

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_6_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkDax_____Sabi_Uses_zeroDs-12                 	1000000000	         0.7970 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_Sabi_Uses_zeroDs-12                 	1000000000	         0.7520 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____Sabi_Uses_oneDs-12                  	33206764	        35.11 ns/op	      64 B/op	       1 allocs/op
BenchmarkDaxPrev_Sabi_Uses_oneDs-12                  	33387213	        35.81 ns/op	      64 B/op	       1 allocs/op
BenchmarkDax_____Sabi_Uses_fiveDs-12                 	 6906590	       176.0 ns/op	     320 B/op	       5 allocs/op
BenchmarkDaxPrev_Sabi_Uses_fiveDs-12                 	 6991933	       173.2 ns/op	     320 B/op	       5 allocs/op
BenchmarkDax_____Sabi_Setup_zeroDs-12                	1000000000	         0.2496 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_Sabi_Setup_zeroDs-12                	1000000000	         0.2533 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____Sabi_Setup_oneDs-12                 	27534898	        40.13 ns/op	      64 B/op	       1 allocs/op
BenchmarkDaxPrev_Sabi_Setup_oneDs-12                 	28899537	        40.73 ns/op	      64 B/op	       1 allocs/op
BenchmarkDax_____Sabi_Setup_fiveDs-12                	23774223	        50.07 ns/op	      64 B/op	       1 allocs/op
BenchmarkDaxPrev_Sabi_Setup_fiveDs-12                	23918625	        50.25 ns/op	      64 B/op	       1 allocs/op
BenchmarkDax_____Sabi_Close_zeroDs-12                	1000000000	         0.4988 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_Sabi_Close_zeroDs-12                	1000000000	         0.2504 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____Sabi_Close_oneDs-12                 	498210192	         2.390 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_Sabi_Close_oneDs-12                 	463511821	         2.212 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____Sabi_Close_freeDs-12                	142120209	         9.157 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_Sabi_Close_freeDs-12                	100000000	        11.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____Sabi_NewDaxBase-12                  	75476475	        15.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_Sabi_NewDaxBase-12                  	83701246	        14.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____DaxBase_Uses_zeroDs-12              	1000000000	         0.5547 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_DaxBase_Uses_zeroDs-12              	1000000000	         0.6313 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____DaxBase_Uses_oneDs-12               	17254132	        64.79 ns/op	      64 B/op	       1 allocs/op
BenchmarkDaxPrev_DaxBase_Uses_oneDs-12               	23676795	        53.16 ns/op	      64 B/op	       1 allocs/op
BenchmarkDax_____DaxBase_Uses_fiveDs-12              	 4508553	       263.8 ns/op	     320 B/op	       5 allocs/op
BenchmarkDaxPrev_DaxBase_Uses_fiveDs-12              	 4590952	       264.9 ns/op	     320 B/op	       5 allocs/op
BenchmarkDax_____DaxBase_Disuses_zeroDs-12           	 3572648	       302.2 ns/op	     320 B/op	       5 allocs/op
BenchmarkDaxPrev_DaxBase_Disuses_zeroDs-12           	 6052312	       251.0 ns/op	     320 B/op	       5 allocs/op
BenchmarkDax_____DaxBase_Disuses_oneDs-12            	 5931915	       213.2 ns/op	     320 B/op	       5 allocs/op
BenchmarkDaxPrev_DaxBase_Disuses_oneDs-12            	 5824082	       268.3 ns/op	     320 B/op	       5 allocs/op
BenchmarkDax_____DaxBase_Disuses_fiveDs-12           	 4846239	       260.0 ns/op	     320 B/op	       5 allocs/op
BenchmarkDaxPrev_DaxBase_Disuses_fiveDs-12           	 4780958	       314.2 ns/op	     320 B/op	       5 allocs/op
BenchmarkDax_____DaxBase_Close_zeroDs-12             	355681702	         2.996 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_DaxBase_Close_zeroDs-12             	593412916	         1.758 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____DaxBase_Close_oneDs-12              	21266913	        56.49 ns/op	      64 B/op	       1 allocs/op
BenchmarkDaxPrev_DaxBase_Close_oneDs-12              	14286130	        86.57 ns/op	      64 B/op	       1 allocs/op
BenchmarkDax_____DaxBase_Close_fiveDs-12             	 3859932	       302.4 ns/op	     320 B/op	       5 allocs/op
BenchmarkDaxPrev_DaxBase_Close_fiveDs-12             	 2863428	       490.4 ns/op	     320 B/op	       5 allocs/op
BenchmarkDax_____Sabi_GetDaxConn_global_oneDs-12     	63669894	        15.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_Sabi_GetDaxConn_global_oneDs-12     	89711731	        14.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____Sabi_GetDaxConn_global_fiveDs-12    	14732191	        85.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_Sabi_GetDaxConn_global_fiveDs-12    	15846766	        73.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____Sabi_GetDaxConn_local_oneDs-12      	89885505	        15.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_Sabi_GetDaxConn_local_oneDs-12      	76390438	        13.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____Sabi_GetDaxConn_local_fiveDs-12     	17645712	        74.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkDaxPrev_Sabi_GetDaxConn_local_fiveDs-12     	16679683	        72.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkDax_____Sabi_Txn_commit_oneDs-12            	 6084776	       165.4 ns/op	      80 B/op	       2 allocs/op
BenchmarkDaxPrev_Sabi_Txn_commit_oneDs-12            	 8548458	       124.8 ns/op	      64 B/op	       1 allocs/op
BenchmarkDax_____Sabi_Txn_commit_fiveDs-12           	 3052701	       471.7 ns/op	      80 B/op	       2 allocs/op
BenchmarkDaxPrev_Sabi_Txn_commit_fiveDs-12           	 3094785	       385.3 ns/op	      64 B/op	       1 allocs/op
BenchmarkDax_____Sabi_Txn_rollback_oneDs-12          	 6831175	       169.3 ns/op	      80 B/op	       2 allocs/op
BenchmarkDaxPrev_Sabi_Txn_rollback_oneDs-12          	10189089	       125.7 ns/op	      64 B/op	       1 allocs/op
BenchmarkDax_____Sabi_Txn_rollback_fiveDs-12         	 2910828	       400.6 ns/op	      80 B/op	       2 allocs/op
BenchmarkDaxPrev_Sabi_Txn_rollback_fiveDs-12         	 3459291	       348.0 ns/op	      64 B/op	       1 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_6_0	77.829s
```
