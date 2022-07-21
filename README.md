# bencher
benchmarks for proto with redis

## Usage
### Run benchmarks
```
make test
```

## Benchmark Results
```
Benchmark_MerchantV1_Proto_Marshal-8      	 3713834	       374.9 ns/op	         2.999 B/serial	       3 B/op	       0 allocs/op
Benchmark_MerchantV1_Proto_Unmarshal-8    	 4975596	       316.5 ns/op	         3.000 B/serial	      48 B/op	       1 allocs/op
Benchmark_MerchantV1_RedisSet-8           	   37658	     28998 ns/op
Benchmark_MerchantV38_Proto_Marshal-8     	 1285106	       979.7 ns/op	        49.49 B/serial	      57 B/op	       1 allocs/op
Benchmark_MerchantV38_Proto_Unmarshal-8   	 2517289	       499.3 ns/op	        49.50 B/serial	      80 B/op	       1 allocs/op
Benchmark_MerchantV38_RedisSet-8          	   37695	     28892 ns/op
PASS
ok  	github.com/alok87/bencher	17.859s
```

## Contributing

- add `.proto` file with your message. Eg: `vi merchant_v1.proto`
- `make generate` to generate the `.pg.go` Eg: `merchant_v1.pb.go`
- add your test functions in `bencher_v1_test.go`
- `make test` to run the benchmark tests