# Benchmark of Err in sabi v0.4.0

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_sabi/err/v0_4_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_ok-12                         	1000000000	         0.2782 ns/op	       0 B/op	       0 allocs/op
Benchmark_ok_isOk-12                    	1000000000	         0.2463 ns/op	       0 B/op	       0 allocs/op
Benchmark_ok_typeSwitch-12              	1000000000	         0.2468 ns/op	       0 B/op	       0 allocs/op
Benchmark_ok_ErrorString-12             	576219542	         1.787 ns/op	       0 B/op	       0 allocs/op
Benchmark_empty-12                      	302526063	         3.962 ns/op	       0 B/op	       0 allocs/op
Benchmark_empty_isNotOk-12              	1000000000	         0.2467 ns/op	       0 B/op	       0 allocs/op
Benchmark_empty_typeSwitch-12           	1000000000	         0.5537 ns/op	       0 B/op	       0 allocs/op
Benchmark_empty_ErrorString-12          	12593649	        93.73 ns/op	      48 B/op	       2 allocs/op
Benchmark_oneField-12                   	36816217	        31.81 ns/op	      16 B/op	       1 allocs/op
Benchmark_oneFieldPtr-12                	37750860	        30.15 ns/op	      16 B/op	       1 allocs/op
Benchmark_oneField_isNotOk-12           	1000000000	         0.2492 ns/op	       0 B/op	       0 allocs/op
Benchmark_oneField_typeSwitch-12        	1000000000	         0.5643 ns/op	       0 B/op	       0 allocs/op
Benchmark_oneField_ErrorString-12       	 4253035	       279.1 ns/op	     136 B/op	       5 allocs/op
Benchmark_fiveField-12                  	27297681	        42.48 ns/op	      64 B/op	       1 allocs/op
Benchmark_fiveField_isNotOk-12          	1000000000	         0.2513 ns/op	       0 B/op	       0 allocs/op
Benchmark_fiveField_typeSwitch-12       	1000000000	         0.5589 ns/op	       0 B/op	       0 allocs/op
Benchmark_fiveField_ErrorString-12      	 1000000	      1038 ns/op	     536 B/op	      17 allocs/op
Benchmark_havingCause-12                	246848727	         4.635 ns/op	       0 B/op	       0 allocs/op
Benchmark_havingCause_ErrorString-12    	 7963022	       145.9 ns/op	     128 B/op	       3 allocs/op
PASS
ok  	github.com/sttk/benchmarks_sabi/err/v0_4_0	17.036s
```
