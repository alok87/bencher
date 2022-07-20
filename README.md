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

benchmark                                | iter      | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-----------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_MerchantV1_Proto_Marshal-8     |    3695186 |    365 ns/op |     3 |   3 |   1.35 |    1109 |  121.97
Benchmark_MerchantV1_Proto_Unmarshal-8   |    4822508 |    299 ns/op |     3 |  48 |   1.44 |    1446 |    6.24


Totals:


benchmark                                | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_MerchantV1_Proto_-8            |    8517694 |    665 ns/op |     6 |  51 |   5.67 |    5112 |   13.05

## Contributing

- add .proto file with your message. Eg: `vi merchant_v1.proto`
- `make generate`
- add your test functions in `bencher_test.go`
- `make test`
- `make stats`