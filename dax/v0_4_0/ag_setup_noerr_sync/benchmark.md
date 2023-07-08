# Benchmark of Dax in sabi v0.4.0

## ag_setup_noerr_sync (`FooDaxSrc{err Err)`, `FooDaxConn{err Err}`)

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/ag_setup_noerr_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	36129513	        33.58 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 6867609	       173.8 ns/op	     352 B/op	       4 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 2226220	       555.6 ns/op	     736 B/op	      12 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.4818 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	43478040	        26.81 ns/op	      24 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	30180351	        40.30 ns/op	      24 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2400 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	411578144	         2.859 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	150771272	         7.905 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	142235084	         8.407 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	33904303	        34.22 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 5848060	       201.7 ns/op	     376 B/op	       5 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 1768234	       672.5 ns/op	     856 B/op	      17 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	 6546687	       180.7 ns/op	     160 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	 5585734	       209.9 ns/op	     160 B/op	       5 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 3815524	       309.0 ns/op	     160 B/op	       5 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	34926063	        35.54 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 5746252	       204.4 ns/op	     376 B/op	       5 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 1741267	       687.2 ns/op	     856 B/op	      17 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	94434903	        12.40 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	17867030	        67.27 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	95126042	        12.32 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	17204355	        68.32 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 4619359	       256.1 ns/op	     376 B/op	       5 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1410369	       868.0 ns/op	     760 B/op	      13 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 4566908	       259.8 ns/op	     376 B/op	       5 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1451059	       838.3 ns/op	     760 B/op	      13 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/ag_setup_noerr_sync	39.603s
```

## ag_setup_noerr_sync (`FooDaxSrc{)`, `FooDaxConn{}`)

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/ag_setup_noerr_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	35877436	        33.76 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 8369018	       144.0 ns/op	     320 B/op	       3 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 3091483	       388.4 ns/op	     576 B/op	       7 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2442 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	44096616	        27.02 ns/op	      24 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	29922019	        39.96 ns/op	      24 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2418 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	456007190	         2.617 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	151151263	         7.909 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	143739090	         8.405 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	31222717	        33.25 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 7065795	       170.2 ns/op	     344 B/op	       4 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 2289168	       514.0 ns/op	     696 B/op	      12 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	34870059	        33.36 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	22820354	        51.82 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 8169276	       143.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	33857328	        34.82 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 6875002	       176.0 ns/op	     344 B/op	       4 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 2167164	       530.7 ns/op	     696 B/op	      12 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	94659253	        12.42 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	17500635	        68.66 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	95711868	        12.61 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	16745740	        70.53 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 5133286	       229.7 ns/op	     344 B/op	       4 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1889331	       643.5 ns/op	     600 B/op	       8 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 5184817	       233.0 ns/op	     344 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1848537	       627.4 ns/op	     600 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/ag_setup_noerr_sync	37.688s
```

## orderedmap_sync

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_sync
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_AddGlobalDaxSrc_zeroDs-12          	35065758	        34.50 ns/op	      48 B/op	       1 allocs/op
Benchmark_AddGlobalDaxSrc_oneDs-12           	 8048066	       146.1 ns/op	     320 B/op	       3 allocs/op
Benchmark_AddGlobalDaxSrc_fiveDs-12          	 3097280	       387.0 ns/op	     576 B/op	       7 allocs/op
Benchmark_StartUpGlobalDaxSrcs_zeroDs-12     	1000000000	         0.2506 ns/op	       0 B/op	       0 allocs/op
Benchmark_StartUpGlobalDaxSrcs_oneDs-12      	44978307	        25.84 ns/op	      16 B/op	       1 allocs/op
Benchmark_StartUpGlobalDaxSrcs_fiveDs-12     	29820260	        40.24 ns/op	      16 B/op	       1 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_zeroDs-12    	1000000000	         0.2412 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_oneDs-12     	385634155	         3.087 ns/op	       0 B/op	       0 allocs/op
Benchmark_ShutdownGlobalDaxSrcs_freeDs-12    	135947154	         8.819 ns/op	       0 B/op	       0 allocs/op
Benchmark_NewDaxBase-12                      	144207121	         8.513 ns/op	       0 B/op	       0 allocs/op
Benchmark_SetUpLocalDaxSrc_zeroDs-12         	36495001	        33.37 ns/op	      48 B/op	       1 allocs/op
Benchmark_SetUpLocalDaxSrc_oneDs-12          	 7814625	       150.9 ns/op	     320 B/op	       3 allocs/op
Benchmark_SetUpLocalDaxSrc_fiveDs-12         	 2854320	       415.3 ns/op	     576 B/op	       7 allocs/op
Benchmark_FreeLocalDaxSrc_zeroDs-12          	34728174	        33.09 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_oneDs-12           	22413698	        52.44 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeLocalDaxSrc_fiveDs-12          	 8386680	       141.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_FreeAllLocalDaxSrcs_zeroDs-12      	32952954	        35.44 ns/op	      48 B/op	       1 allocs/op
Benchmark_FreeAllLocalDaxSrcs_oneDs-12       	 7713186	       156.7 ns/op	     320 B/op	       3 allocs/op
Benchmark_FreeAllLocalDaxSrcs_fiveDs-12      	 2751710	       431.8 ns/op	     576 B/op	       7 allocs/op
Benchmark_GetDaxConn_global_oneDs-12         	95433308	        12.22 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_global_fiveDs-12        	17879406	        66.34 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_oneDs-12          	98523808	        12.18 ns/op	       0 B/op	       0 allocs/op
Benchmark_GetDaxConn_local_fiveDs-12         	17321272	        68.06 ns/op	       0 B/op	       0 allocs/op
Benchmark_RunTxn_commit_oneDs-12             	 4544973	       266.5 ns/op	     384 B/op	       5 allocs/op
Benchmark_RunTxn_commit_fiveDs-12            	 1770678	       684.8 ns/op	     640 B/op	       9 allocs/op
Benchmark_RunTxn_rollback_oneDs-12           	 5143808	       230.9 ns/op	     336 B/op	       4 allocs/op
Benchmark_RunTxn_rollback_fiveDs-12          	 1883058	       620.5 ns/op	     592 B/op	       8 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/dax/v0_4_0/orderedmap_sync	37.757s
```
