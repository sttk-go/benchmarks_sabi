# Benchmark of Err in sabi v0.4.0

> - `BenchmarkGoError_*` is go standard error.
> - `BenchmarkErr_____*` is v0.4.0.

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/err/v0_4_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkGoError_ok-12                         	1000000000	         0.2570 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____ok-12                         	1000000000	         0.2504 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_ok_isOk-12                    	1000000000	         0.2509 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____ok_isOk-12                    	1000000000	         0.2506 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_ok_typeSwitch-12              	1000000000	         0.2501 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____ok_typeSwitch-12              	1000000000	         0.2514 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_ok_ErrorString-12             	1000000000	         0.2495 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____ok_ErrorString-12             	640279188	         1.928 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_empty-12                      	1000000000	         0.2515 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____empty-12                      	299973758	         3.996 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_empty_isNotOk-12              	1000000000	         0.2509 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____empty_isNotOk-12              	1000000000	         0.2502 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_empty_typeSwitch-12           	1000000000	         0.2506 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____empty_typeSwitch-12           	1000000000	         0.7739 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_empty_ErrorString-12          	1000000000	         1.006 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____empty_ErrorString-12          	12266066	        96.05 ns/op	      48 B/op	       2 allocs/op
BenchmarkGoError_oneField-12                   	41210586	        30.12 ns/op	      16 B/op	       1 allocs/op
BenchmarkErr_____oneField-12                   	36676057	        33.46 ns/op	      16 B/op	       1 allocs/op
BenchmarkGoError_oneFieldPtr-12                	41965396	        28.55 ns/op	      16 B/op	       1 allocs/op
BenchmarkErr_____oneFieldPtr-12                	37552299	        31.56 ns/op	      16 B/op	       1 allocs/op
BenchmarkGoError_oneField_isNotOk-12           	1000000000	         0.2622 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____oneField_isNotOk-12           	1000000000	         0.2928 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_oneField_typeSwitch-12        	1000000000	         0.2593 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____oneField_typeSwitch-12        	1000000000	         0.7746 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_oneField_ErrorString-12       	28284916	        41.90 ns/op	      32 B/op	       1 allocs/op
BenchmarkErr_____oneField_ErrorString-12       	 4231530	       287.0 ns/op	     136 B/op	       5 allocs/op
BenchmarkGoError_fiveField-12                  	30384162	        41.74 ns/op	      64 B/op	       1 allocs/op
BenchmarkErr_____fiveField-12                  	25718036	        46.01 ns/op	      64 B/op	       1 allocs/op
BenchmarkGoError_fiveField_isNotOk-12          	1000000000	         0.2684 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____fiveField_isNotOk-12          	1000000000	         0.2503 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_fiveField_typeSwitch-12       	1000000000	         0.2539 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____fiveField_typeSwitch-12       	1000000000	         0.7615 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_fiveField_ErrorString-12      	10846029	       109.3 ns/op	      83 B/op	       2 allocs/op
BenchmarkErr_____fiveField_ErrorString-12      	 1000000	      1047 ns/op	     536 B/op	      17 allocs/op
BenchmarkGoError_havingCause-12                	37899298	        32.77 ns/op	      16 B/op	       1 allocs/op
BenchmarkErr_____havingCause-12                	251143336	         4.746 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_havingCause_ErrorString-12    	25850918	        44.83 ns/op	      48 B/op	       1 allocs/op
BenchmarkErr_____havingCause_ErrorString-12    	 7745919	       154.1 ns/op	     128 B/op	       3 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/err/v0_4_0	31.208s
```
