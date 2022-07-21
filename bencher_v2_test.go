package bencher

import (
	"math/rand"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomMerchantV38(n int) []*MerchantV38Fields {
	merchant := make([]*MerchantV38Fields, 0, n)
	for i := 0; i < n; i++ {
		merchant = append(merchant, &MerchantV38Fields{
			RefundDisabled:     rand.Intn(2) == 1,
			AutoRefundDisabled: rand.Intn(2) == 1,
			BlockSettlement:    rand.Intn(2) == 1,
			BlockSettlement4:   rand.Intn(2) == 1,
			BlockSettlement5:   rand.Intn(2) == 1,
			BlockSettlement6:   rand.Intn(2) == 1,
			BlockSettlement7:   rand.Intn(2) == 1,
			BlockSettlement8:   rand.Intn(2) == 1,
			BlockSettlement9:   rand.Intn(2) == 1,
			BlockSettlement10:  rand.Intn(2) == 1,
			BlockSettlement11:  rand.Intn(2) == 1,
			BlockSettlement12:  rand.Intn(2) == 1,
			BlockSettlement13:  rand.Intn(2) == 1,
			BlockSettlement14:  rand.Intn(2) == 1,
			BlockSettlement15:  rand.Intn(2) == 1,
			BlockSettlement16:  rand.Intn(2) == 1,
			BlockSettlement17:  rand.Intn(2) == 1,
			BlockSettlement18:  rand.Intn(2) == 1,
			BlockSettlement19:  rand.Intn(2) == 1,
			BlockSettlement20:  rand.Intn(2) == 1,
			BlockSettlement21:  rand.Intn(2) == 1,
			BlockSettlement22:  rand.Intn(2) == 1,
			BlockSettlement23:  rand.Intn(2) == 1,
			BlockSettlement24:  rand.Intn(2) == 1,
			BlockSettlement25:  rand.Intn(2) == 1,
			BlockSettlement26:  rand.Intn(2) == 1,
			BlockSettlement27:  rand.Intn(2) == 1,
			BlockSettlement28:  rand.Intn(2) == 1,
			BlockSettlement29:  rand.Intn(2) == 1,
			BlockSettlement30:  rand.Intn(2) == 1,
			BlockSettlement31:  rand.Intn(2) == 1,
			BlockSettlement32:  rand.Intn(2) == 1,
			BlockSettlement33:  rand.Intn(2) == 1,
			BlockSettlement34:  rand.Intn(2) == 1,
			BlockSettlement35:  rand.Intn(2) == 1,
			BlockSettlement36:  rand.Intn(2) == 1,
			BlockSettlement37:  rand.Intn(2) == 1,
			BlockSettlement38:  rand.Intn(2) == 1,
		})
	}
	return merchant
}

func Benchmark_ProtoMarshal_MerchantV38(b *testing.B) {
	merchant := randomMerchantV38(b.N)
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

func Benchmark_ProtoUnmarshal_MerchantV38(b *testing.B) {
	b.StopTimer()
	merchant := randomMerchantV38(b.N)
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
		o := &MerchantV38Fields{}
		err := proto.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}

func Benchmark_RedisSet_MerchantV38(b *testing.B) {
	b.StopTimer()
	client := newRedisClient()
	key := "refund/merchant/mid_v38"
	merchant := randomMerchantV38(b.N)[0]
	merchantBytes, err := proto.Marshal(merchant)
	if err != nil {
		b.Error(err)
		return
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		err := client.Set(key, merchantBytes, 0).Err()
		if err != nil {
			klog.Error(err)
			b.Error(err)
		}
		b.StartTimer()
	}
}

func Benchmark_RedisGet_MerchantV38(b *testing.B) {
	b.StopTimer()
	client := newRedisClient()
	key := "refund/merchant/mid_v38"

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		err := client.Get(key).Err()
		if err != nil {
			klog.Error(err)
			b.Error(err)
		}
		b.StartTimer()
	}
}
