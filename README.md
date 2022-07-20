# bencher
benchmarks proto and json, store data in Redis

## Usage
```
make test
```

## Benchmark Results
```
Benchmark_MerchantV1_Proto_Marshal-8     	 3695186	       365.9 ns/op	         3.002 B/serial	       3 B/op	       0 allocs/op
Benchmark_MerchantV1_Proto_Unmarshal-8   	 4822508	       299.6 ns/op	         3.000 B/serial	      48 B/op	       1 allocs/op
```