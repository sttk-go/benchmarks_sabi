# err

## go_error

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/err/go_error
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_ok-12                         	1000000000	         0.2459 ns/op	       0 B/op	       0 allocs/op
Benchmark_ok_isOk-12                    	1000000000	         0.2426 ns/op	       0 B/op	       0 allocs/op
Benchmark_ok_typeSwitch-12              	1000000000	         0.2416 ns/op	       0 B/op	       0 allocs/op
Benchmark_ok_ErrorString-12             	1000000000	         0.2420 ns/op	       0 B/op	       0 allocs/op
Benchmark_empty-12                      	1000000000	         0.2407 ns/op	       0 B/op	       0 allocs/op
Benchmark_empty_isOk-12                 	1000000000	         0.2415 ns/op	       0 B/op	       0 allocs/op
Benchmark_empty_typeSwitch-12           	1000000000	         0.2402 ns/op	       0 B/op	       0 allocs/op
Benchmark_empty_ErrorString-12          	961870692	         1.207 ns/op	       0 B/op	       0 allocs/op
Benchmark_oneField-12                   	41542165	        27.96 ns/op	      16 B/op	       1 allocs/op
Benchmark_oneFieldPtr-12                	44212088	        27.00 ns/op	      16 B/op	       1 allocs/op
Benchmark_oneField_isOk-12              	1000000000	         0.2403 ns/op	       0 B/op	       0 allocs/op
Benchmark_oneField_typeSwitch-12        	1000000000	         0.2407 ns/op	       0 B/op	       0 allocs/op
Benchmark_oneField_ErrorString-12       	27827743	        39.15 ns/op	      32 B/op	       1 allocs/op
Benchmark_fiveField-12                  	31044304	        38.64 ns/op	      64 B/op	       1 allocs/op
Benchmark_fiveField_isOk-12             	1000000000	         0.4800 ns/op	       0 B/op	       0 allocs/op
Benchmark_fiveField_typeSwitch-12       	1000000000	         0.2395 ns/op	       0 B/op	       0 allocs/op
Benchmark_fiveField_ErrorString-12      	10611986	       113.8 ns/op	      83 B/op	       2 allocs/op
Benchmark_havingCause-12                	35803155	        32.50 ns/op	      16 B/op	       1 allocs/op
Benchmark_havingCause_ErrorString-12    	27869539	        45.32 ns/op	      48 B/op	       1 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/err/go_error	13.781s
```

## sabi_err

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/err/sabi_err
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_ok-12                         	1000000000	         0.2428 ns/op	       0 B/op	       0 allocs/op
Benchmark_ok_isOk-12                    	1000000000	         0.2461 ns/op	       0 B/op	       0 allocs/op
Benchmark_ok_typeSwitch-12              	1000000000	         0.2415 ns/op	       0 B/op	       0 allocs/op
Benchmark_ok_ErrorString-12             	701801460	         1.678 ns/op	       0 B/op	       0 allocs/op
Benchmark_empty-12                      	310454064	         3.863 ns/op	       0 B/op	       0 allocs/op
Benchmark_empty_isNotOk-12              	1000000000	         0.2477 ns/op	       0 B/op	       0 allocs/op
Benchmark_empty_typeSwitch-12           	1000000000	         0.7513 ns/op	       0 B/op	       0 allocs/op
Benchmark_empty_ErrorString-12          	12881445	        90.08 ns/op	      48 B/op	       2 allocs/op
Benchmark_oneField-12                   	38000936	        32.32 ns/op	      16 B/op	       1 allocs/op
Benchmark_oneFieldPtr-12                	38169219	        33.11 ns/op	      16 B/op	       1 allocs/op
Benchmark_oneField_isNotOk-12           	1000000000	         0.2526 ns/op	       0 B/op	       0 allocs/op
Benchmark_oneField_typeSwitch-12        	1000000000	         0.7359 ns/op	       0 B/op	       0 allocs/op
Benchmark_oneField_ErrorString-12       	 4197573	       300.1 ns/op	     136 B/op	       5 allocs/op
Benchmark_fiveField-12                  	26935065	        48.32 ns/op	      64 B/op	       1 allocs/op
Benchmark_fiveField_isNotOk-12          	1000000000	         0.2499 ns/op	       0 B/op	       0 allocs/op
Benchmark_fiveField_typeSwitch-12       	1000000000	         0.7842 ns/op	       0 B/op	       0 allocs/op
Benchmark_fiveField_ErrorString-12      	  922323	      1229 ns/op	     536 B/op	      17 allocs/op
Benchmark_havingCause-12                	251967697	         6.640 ns/op	       0 B/op	       0 allocs/op
Benchmark_havingCause_ErrorString-12    	 7394916	       163.4 ns/op	     128 B/op	       3 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/err/sabi_err	20.837s
```
