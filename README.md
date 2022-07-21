# bencher
benchmarks for proto with redis

## Usage
### Run benchmarks
```
make test
```

## Benchmark Results
Takes 0.03ms for basic SET and GET in Redis without load in local 
### Local results when service and redis are in same host (local imac)
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

### EC2 (mumbai) -> Redis Labs (Mumbai)
Takes 0.3ms for basic SET and GET in Redis without load in Mumbai.

```
Benchmark_ProtoUnmarshal_MerchantV1-2    	 3226676	       407.9 ns/op	         3.000 B/serial	       3 B/op	       0 allocs/op
Benchmark_ProtoMarshal_MerchantV1-2      	 3087302	       431.0 ns/op	         2.998 B/serial	      48 B/op	       1 allocs/op

Benchmark_ProtoMarshal_MerchantV38-2     	  982110	      1223 ns/op	        49.51 B/serial	      57 B/op	       1 allocs/op
Benchmark_ProtoUnmarshal_MerchantV38-2   	 1792717	       719.2 ns/op	        49.51 B/serial	      80 B/op	       1 allocs/op

Benchmark_RedisSet_MerchantV1-2          	    3054	    388322 ns/op
Benchmark_RedisSet_MerchantV38-2         	    3162	    365419 ns/op

Benchmark_RedisGet_MerchantV1-2          	    2562	    397538 ns/op
Benchmark_RedisGet_MerchantV38-2         	    3160	    370145 ns/op
```
## Contributing

- add `.proto` file with your message. Eg: `vi merchant_v1.proto`
- `make generate` to generate the `.pg.go` Eg: `merchant_v1.pb.go`
- add your test functions in `bencher_v1_test.go`
- `make test` to run the benchmark tests (local)
- export REDIS_HOST=https://someurl.com make test