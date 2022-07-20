package bencher

import (
	"math/rand"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// func randString(l int) string {
// 	buf := make([]byte, l)
// 	for i := 0; i < (l+1)/2; i++ {
// 		buf[i] = byte(rand.Intn(256))
// 	}
// 	return fmt.Sprintf("%x", buf)[:l]
// }

func randomMerchantV1(n int) []*MerchantV1 {
	merchant := make([]*MerchantV1, 0, n)
	for i := 0; i < n; i++ {
		merchant = append(merchant, &MerchantV1{
			RefundDisabled:     rand.Intn(2) == 1,
			AutoRefundDisabled: rand.Intn(2) == 1,
			BlockSettlement:    rand.Intn(2) == 1,
		})
	}
	return merchant
}

func randomMerchantV38Fields(n int) []*MerchantV38Fields {
	merchant := make([]*MerchantV38Fields, 0, n)
	for i := 0; i < n; i++ {
		merchant = append(merchant, &MerchantV38Fields{
			RefundDisabled:     rand.Intn(2) == 1,
			AutoRefundDisabled: rand.Intn(2) == 1,
			BlockSettlement:    rand.Intn(2) == 1,
		})
	}
	return merchant
}

func Benchmark_MerchantV1_Proto_Marshal(b *testing.B) {
	merchant := randomMerchantV1(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := proto.Marshal(merchant[rand.Intn(len(merchant))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_MerchantV1_Proto_Unmarshal(b *testing.B) {
	b.StopTimer()
	merchant := randomMerchantV1(b.N)
	ser := make([][]byte, len(merchant))
	var serialSize int
	for i, d := range merchant {
		var err error
		ser[i], err = proto.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(merchant)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &MerchantV1{}
		err := proto.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}

func Benchmark_MerchantV38Fields_Proto_Marshal(b *testing.B) {
	merchant := randomMerchantV38Fields(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := proto.Marshal(merchant[rand.Intn(len(merchant))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_MerchantV38Fields_Proto_Unmarshal(b *testing.B) {
	b.StopTimer()
	merchant := randomMerchantV38Fields(b.N)
	ser := make([][]byte, len(merchant))
	var serialSize int
	for i, d := range merchant {
		var err error
		ser[i], err = proto.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(merchant)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &MerchantV1{}
		err := proto.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}
