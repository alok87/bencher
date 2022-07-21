# bencher
benchmarks for proto with redis

## Usage
### Run benchmarks
```
make test
```

## Benchmark Results
### Local results when service and redis are in same host (local mac)
```
Benchmark_ProtoUnmarshal_MerchantV1-8    	 3567019	       371.9 ns/op	         3.001 B/serial	       3 B/op	       0 allocs/op
Benchmark_ProtoMarshal_MerchantV1-8      	 4795518	       329.9 ns/op	         3.002 B/serial	      48 B/op	       1 allocs/op
Benchmark_ProtoMarshal_MerchantV38-8     	 1221247	       981.7 ns/op	        49.52 B/serial	      57 B/op	       1 allocs/op
Benchmark_ProtoUnmarshal_MerchantV38-8   	 2493229	       512.5 ns/op	        49.50 B/serial	      80 B/op	       1 allocs/op

Benchmark_RedisSet_MerchantV1-8          	   37177	     31211 ns/op
Benchmark_RedisSet_MerchantV38-8         	   40179	     31315 ns/op

Benchmark_RedisGet_MerchantV1-8          	   39410	     30966 ns/op
Benchmark_RedisGet_MerchantV38-8         	   37562	     30866 ns/op
```

## Contributing

- add `.proto` file with your message. Eg: `vi merchant_v1.proto`
- `make generate` to generate the `.pg.go` Eg: `merchant_v1.pb.go`
- add your test functions in `bencher_v1_test.go`
- `make test` to run the benchmark tests (local)
- export REDIS_HOST=https://someurl.com make test