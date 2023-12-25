package inmemorycache

import (
	"testing"
	"time"

	cache "github.com/jellydator/ttlcache/v3"
)

func BenchmarkTTLCache(b *testing.B) {
	for _, tt := range tests {
		benchmarkTTLCache(b, tt.name, tt.valueGenerator)
	}
}

func benchmarkTTLCache(b *testing.B, prefix string, valueGenerator func(int) string) {
	c := cache.New[string, string]()

	b.Run(testName(prefix, "Set"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			c.Set(generateKey(prefix, i), valueGenerator(i), time.Minute)
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Get"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			c.Get(generateKey(prefix, i))
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Remove"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			c.Delete(generateKey(prefix, i))
		}
		b.StopTimer()
	})
}
