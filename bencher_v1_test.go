package bencher

import (
	"math/rand"
	reflect "reflect"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"
)

func newRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:36379",
		Password: "",
		DB:       0,
	})

	return client
}

func UpdateField(s interface{}, field string, value interface{}) {
	v := reflect.ValueOf(s).Elem().FieldByName(field)
	if v.IsValid() {
		v.Set(reflect.ValueOf(value))
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

func Benchmark_MerchantV1_RedisSet(b *testing.B) {
	b.StopTimer()
	client := newRedisClient()
	key := "refund/merchant/mid_v1"
	merchant := randomMerchantV1(b.N)[0]
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

// func Benchmark_MerchantV1_RedisSet(b *testing.B) {
// 	client := newRedisClient()
// 	ctx := context.Background()
// 	key := "refund/merchant/mid_657_swiggy"

// 	// optimistic tx
// 	// https://pkg.go.dev/github.com/go-redis/redis/v8#Client.Watch
// 	txFunc := func(tx *redis.Tx) error {
// 		// Get current value or zero.
// 		b, err := tx.Get(ctx, key).Bytes()
// 		if err != nil && err != redis.Nil {
// 			return err
// 		}

// 		// Operation is committed only if the watched keys remain unchanged.
// 		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
// 			pipe.Set(ctx, key, n, 0)
// 			return nil
// 		})
// 		return err
// 	}

// }
