# bencher
benchmarks proto and json, store data in Redis

## Usage
### Run benchmarks
```
make test
```
### Print the table to update the results in the README.
```
make stats
```

## Benchmark Results
Test result if we have 3 field struct and 38 field struct

benchmark                                | iter      | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-----------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_MerchantV1_Proto_Marshal-8     |    3737200 |    374 ns/op |     2 |   3 |   1.40 |    1120 |  124.67
Benchmark_MerchantV1_Proto_Unmarshal-8   |    4949146 |    301 ns/op |     3 |  48 |   1.49 |    1485 |    6.27
Benchmark_MerchantV38Fields_Proto_Marshal-8 |    1000000 |   1039 ns/op |    49 |  56 |   1.04 |    4950 |   18.55
Benchmark_MerchantV38Fields_Proto_Unmarshal-8 |    2580718 |    522 ns/op |    49 |  80 |   1.35 |   12774 |    6.53


Totals:


benchmark                                | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_MerchantV1_Proto_-8            |    8686346 |    675 ns/op |     5 |  51 |   5.87 |    5210 |   13.24
Benchmark_MerchantV38Fields_Proto_-8     |    3580718 |   1561 ns/op |    99 | 136 |   5.59 |   35449 |   11.48

## Contributing

- add .proto file with your message. Eg: `vi merchant_v1.proto`
- `make generate`
- add your test functions in `bencher_test.go`
- `make test`
- `make stats`