# Benchmark of Dax in sabi v0.4.0

## unlock

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/sabi_0_4_0/unlock
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12        	304369935	         3.871 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12       	42540140	        28.55 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12     	 1836556	       652.7 ns/op	     208 B/op	       3 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12    	  388362	      3079 ns/op	     464 B/op	       7 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12      	 1746771	       685.6 ns/op	      64 B/op	       3 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12     	  520314	      2314 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        10.34 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12       	225095977	         5.311 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12      	32087640	        37.10 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12        	36763765	        32.30 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12       	 6514738	       181.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12      	100000000	        11.27 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12     	17902971	        65.36 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12       	100000000	        11.29 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12      	18303699	        66.29 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12          	  721860	      1579 ns/op	     608 B/op	       8 allocs/op
Benchmark_RunTxn_commit_fiveDs-12         	  191293	      6009 ns/op	    1056 B/op	      20 allocs/op
Benchmark_RunTxn_rollback_oneDs-12        	  723151	      1586 ns/op	     464 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12       	  238455	      4964 ns/op	     848 B/op	      24 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/sabi_0_4_0/unlock	25.375s
```

## original

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/sabi_0_4_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_oneDs-12        	67004229	        17.56 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12       	13215918	        91.25 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupGlobalDaxSrcs_oneDs-12     	 1843701	       647.8 ns/op	     208 B/op	       3 allocs/op
Benchmark_SetupGlobalDaxSrcs_fiveDs-12    	  386847	      3092 ns/op	     464 B/op	       7 allocs/op
Benchmark_FreeGlobalDaxSrcs_oneDs-12      	 1695716	       681.6 ns/op	      64 B/op	       3 allocs/op
Benchmark_FreeGlobalDaxSrcs_fiveDs-12     	  523662	      2319 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                   	100000000	        10.25 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_oneDs-12       	59208042	        19.60 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetupLocalDaxSrc_fiveDs-12      	11425267	       103.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12        	29259128	        40.43 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12       	 5478446	       215.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_oneDs-12      	100000000	        12.02 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12     	18191206	        65.33 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12       	98955450	        12.04 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12      	18064498	        64.94 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12          	  679263	      1587 ns/op	     608 B/op	       8 allocs/op
Benchmark_RunTxn_commit_fiveDs-12         	  187744	      6082 ns/op	    1056 B/op	      20 allocs/op
Benchmark_RunTxn_rollback_oneDs-12        	  726408	      1602 ns/op	     464 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12       	  234010	      4990 ns/op	     848 B/op	      24 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/sabi_0_4_0	24.576s
```
