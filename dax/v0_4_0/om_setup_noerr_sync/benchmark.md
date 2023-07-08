# Benchmark of Dax in sabi v0.4.0

## om_setup_noerr_sync (`FooDaxSrc{err Err}`, `FooDaxConn{err Err}`)

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/om_setup_noerr_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	34246216	        33.85 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 6866948	       175.6 ns/op	     352 B/op	       4 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 2222516	       540.0 ns/op	     736 B/op	      12 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.4932 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	46357569	        25.34 ns/op	      16 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	30181417	        39.08 ns/op	      16 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2454 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	406273615	         2.883 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	146591377	         8.859 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	143332993	         8.373 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	35358080	        34.09 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 6678679	       179.4 ns/op	     352 B/op	       4 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 2144640	       562.7 ns/op	     736 B/op	      12 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	 6488299	       187.5 ns/op	     160 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	 5445489	       210.1 ns/op	     160 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 3807505	       309.2 ns/op	     160 B/op	       5 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	32335369	        35.63 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 6271308	       190.8 ns/op	     352 B/op	       4 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 1882323	       639.4 ns/op	     736 B/op	      12 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	74058086	        16.09 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	16373770	        71.36 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	91880912	        12.74 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	17550134	        68.76 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 4409733	       265.1 ns/op	     368 B/op	       5 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1417244	       838.4 ns/op	     752 B/op	      13 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 4654227	       259.0 ns/op	     368 B/op	       5 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1435492	       839.0 ns/op	     752 B/op	      13 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/om_setup_noerr_sync	38.375s
```

## om_setup_noerr_sync (`FooDaxSrc{}`, `FooDaxConn{}`)

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/om_setup_noerr_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	32270760	        34.17 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 8346688	       144.5 ns/op	     320 B/op	       3 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 3104644	       386.6 ns/op	     576 B/op	       7 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2633 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	47843713	        25.13 ns/op	      16 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	30807631	        38.20 ns/op	      16 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2422 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	452282794	         2.624 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	151776925	         7.944 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	142660273	         8.414 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	34698390	        33.20 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 8180868	       149.0 ns/op	     320 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 2962028	       407.7 ns/op	     576 B/op	       7 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	34191594	        33.23 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	22688180	        52.03 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 8206129	       144.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	34269640	        34.92 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7873468	       152.0 ns/op	     320 B/op	       3 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 2813714	       425.8 ns/op	     576 B/op	       7 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	94566276	        12.28 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	16975410	        67.46 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	96306214	        12.34 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	17363384	        68.28 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 5254120	       228.5 ns/op	     336 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1921756	       617.9 ns/op	     592 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 5111401	       230.5 ns/op	     336 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1957034	       616.5 ns/op	     592 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/om_setup_noerr_sync	36.249s
```

## orderedmap_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	35472178	        33.27 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 7995664	       143.2 ns/op	     320 B/op	       3 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 3133509	       383.2 ns/op	     576 B/op	       7 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2476 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	45606620	        26.33 ns/op	      16 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	30630718	        39.52 ns/op	      16 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2405 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	386292301	         3.103 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	135335476	         8.929 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	143341401	         8.329 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	34846280	        33.01 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 7920618	       150.7 ns/op	     320 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 2876320	       409.9 ns/op	     576 B/op	       7 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	35172376	        33.30 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	22180814	        52.51 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 8316104	       143.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	31367923	        35.24 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7706422	       160.0 ns/op	     320 B/op	       3 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 2791158	       435.5 ns/op	     576 B/op	       7 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	95233446	        12.40 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	17759610	        66.30 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	99537085	        12.23 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	15853816	        69.34 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 4472571	       270.3 ns/op	     384 B/op	       5 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1754631	       679.5 ns/op	     640 B/op	       9 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 5240346	       229.3 ns/op	     336 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1918126	       641.5 ns/op	     592 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_sync	37.709s
```
