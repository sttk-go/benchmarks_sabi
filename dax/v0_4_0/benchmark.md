# Benchmark of Dax in sabi v0.4.0

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	701858133	         1.648 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	31979253	        37.35 ns/op	       0 B/op	       0 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 5505736	       212.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2375 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	 1844935	       650.3 ns/op	     208 B/op	       3 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	  365540	      2999 ns/op	     464 B/op	       7 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2399 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	 1841827	       649.7 ns/op	      64 B/op	       3 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	  552769	      2212 ns/op	     256 B/op	      11 allocs/op
Benchmark_NewDaxBase-12                      	100000000	        10.13 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	685279536	         1.718 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	28733091	        41.47 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 5088334	       231.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	21347149	        55.04 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	12866899	        90.27 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 5487015	       217.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	24816471	        48.02 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 6384585	       188.8 ns/op	     336 B/op	       2 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 3228111	       356.6 ns/op	     336 B/op	       2 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	99709280	        11.88 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	16617764	        69.92 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	95633074	        12.39 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	16957789	        68.36 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	  679772	      1582 ns/op	     608 B/op	       8 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	  186225	      5981 ns/op	    1056 B/op	      20 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	  724579	      1594 ns/op	     464 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	  232851	      4892 ns/op	     848 B/op	      24 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0	33.326s
```
