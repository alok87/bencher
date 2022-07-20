test:
	rm results.txt || true
	go test -count=1 -bench=. | grep Benchmark_Merchant > results.txt
# cd old && go test -count=1 -bench=. | grep Benchmark_Merchant >> ../results.txt
# sort -o results.txt results.txt
	cat results.txt

stats:
	make stats

all: old/structdef-go-v1.pb.go structdef-gogo-v1.pb.go structdef-go-v1.pb.go structdef-go-v2.pb.go

generate:
	@buf generate

