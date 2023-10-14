# Banchmark of Err in sabi v0.6.0

> - `BenchmarkGoError_*` is go standard error.
> - `BenchmarkErrPrev_*` is v0.4.0.
> - `BenchmarkErr_____*` is v0.6.0.

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/err/v0_6_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkGoError_ok-12                         	1000000000	         0.2626 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrPrev_ok-12                         	1000000000	         0.2552 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____ok-12                         	1000000000	         0.2594 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_ok_isOk-12                    	1000000000	         0.2575 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrPrev_ok_isOk-12                    	1000000000	         0.2555 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____ok_isOk-12                    	1000000000	         0.2545 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_ok_typeSwitch-12              	1000000000	         0.2548 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrPrev_ok_typeSwitch-12              	1000000000	         0.2578 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____ok_typeSwitch-12              	1000000000	         0.5148 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_ok_ErrorString-12             	1000000000	         0.2545 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrPrev_ok_ErrorString-12             	620144182	         1.938 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____ok_ErrorString-12             	615272054	         1.952 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_empty-12                      	1000000000	         0.2634 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrPrev_empty-12                      	288217788	         4.139 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____empty-12                      	291390579	         4.220 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_empty_isNotOk-12              	1000000000	         0.2590 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____empty_isNotOk-12              	1000000000	         0.2599 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrPrev_empty_isNotOk-12              	1000000000	         0.2581 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_empty_typeSwitch-12           	1000000000	         0.2555 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrPrev_empty_typeSwitch-12           	1000000000	         0.7734 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____empty_typeSwitch-12           	1000000000	         0.7696 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_empty_ErrorString-12          	1000000000	         1.028 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrPrev_empty_ErrorString-12          	12182444	        95.37 ns/op	      48 B/op	       2 allocs/op
BenchmarkErr_____empty_ErrorString-12          	12205104	        98.32 ns/op	      48 B/op	       2 allocs/op
BenchmarkGoError_oneField-12                   	41419143	        28.68 ns/op	      16 B/op	       1 allocs/op
BenchmarkErrPrev_oneField-12                   	35353527	        33.18 ns/op	      16 B/op	       1 allocs/op
BenchmarkErr_____oneField-12                   	35641498	        33.76 ns/op	      16 B/op	       1 allocs/op
BenchmarkGoError_oneFieldPtr-12                	42807829	        27.68 ns/op	      16 B/op	       1 allocs/op
BenchmarkErrPrev_oneFieldPtr-12                	37972203	        31.47 ns/op	      16 B/op	       1 allocs/op
BenchmarkErr_____oneFieldPtr-12                	37996641	        31.24 ns/op	      16 B/op	       1 allocs/op
BenchmarkGoError_oneField_isNotOk-12           	1000000000	         0.2550 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrPrev_oneField_isNotOk-12           	1000000000	         0.2534 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____oneField_isNotOk-12           	1000000000	         0.2570 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_oneField_typeSwitch-12        	1000000000	         0.2547 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrPrev_oneField_typeSwitch-12        	1000000000	         0.7659 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____oneField_typeSwitch-12        	1000000000	         0.7999 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_oneField_ErrorString-12       	28217631	        42.22 ns/op	      32 B/op	       1 allocs/op
BenchmarkErrPrev_oneField_ErrorString-12       	 4042230	       287.4 ns/op	     136 B/op	       5 allocs/op
BenchmarkErr_____oneField_ErrorString-12       	 4166324	       286.2 ns/op	     136 B/op	       5 allocs/op
BenchmarkGoError_fiveField-12                  	29228023	        40.84 ns/op	      64 B/op	       1 allocs/op
BenchmarkErrPrev_fiveField-12                  	27602919	        43.38 ns/op	      64 B/op	       1 allocs/op
BenchmarkErr_____fiveField-12                  	26800788	        43.31 ns/op	      64 B/op	       1 allocs/op
BenchmarkGoError_fiveField_isNotOk-12          	1000000000	         0.2501 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrPrev_fiveField_isNotOk-12          	1000000000	         0.2505 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____fiveField_isNotOk-12          	1000000000	         0.2504 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_fiveField_typeSwitch-12       	1000000000	         0.2509 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrPrev_fiveField_typeSwitch-12       	1000000000	         0.7632 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____fiveField_typeSwitch-12       	1000000000	         0.7520 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_fiveField_ErrorString-12      	10681282	       114.3 ns/op	      83 B/op	       2 allocs/op
BenchmarkErrPrev_fiveField_ErrorString-12      	 1000000	      1074 ns/op	     536 B/op	      17 allocs/op
BenchmarkErr_____fiveField_ErrorString-12      	 1000000	      1057 ns/op	     536 B/op	      17 allocs/op
BenchmarkGoError_havingCause-12                	36899782	        32.15 ns/op	      16 B/op	       1 allocs/op
BenchmarkErrPrev_havingCause-12                	251106838	         4.756 ns/op	       0 B/op	       0 allocs/op
BenchmarkErr_____havingCause-12                	251690734	         4.762 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoError_havingCause_ErrorString-12    	27842971	        43.25 ns/op	      48 B/op	       1 allocs/op
BenchmarkErrPrev_havingCause_ErrorString-12    	 7508644	       155.1 ns/op	     128 B/op	       3 allocs/op
BenchmarkErr_____havingCause_ErrorString-12    	 7389758	       155.5 ns/op	     128 B/op	       3 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/err/v0_6_0	50.422s
```
