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
Benchmark_MerchantV1_Proto_Marshal-8     |    3833323 |    366 ns/op |     2 |   3 |   1.40 |    1149 |  122.17
Benchmark_MerchantV1_Proto_Unmarshal-8   |    5772439 |    304 ns/op |     3 |  48 |   1.76 |    1731 |    6.34
Benchmark_MerchantV38Fields_Proto_Marshal-8 |    1794448 |    667 ns/op |     3 |   3 |   1.20 |     538 |  222.40
Benchmark_MerchantV38Fields_Proto_Unmarshal-8 |    5314509 |    304 ns/op |     3 |  48 |   1.62 |    1594 |    6.35


Totals:


benchmark                                | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_MerchantV1_Proto_-8            |    9605762 |    670 ns/op |     5 |  51 |   6.44 |    5762 |   13.15
Benchmark_MerchantV38Fields_Proto_-8     |    7108957 |    971 ns/op |     6 |  51 |   6.91 |    4265 |   19.06

## Contributing

- add .proto file with your message. Eg: `vi merchant_v1.proto`
- `make generate`
- add your test functions in `bencher_test.go`
- `make test`
- `make stats`