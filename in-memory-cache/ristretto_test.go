package inmemorycache

import (
	"testing"

	cache "github.com/dgraph-io/ristretto"
)

func BenchmarkRistretto(b *testing.B) {
	for _, tt := range tests {
		benchmarkRistretto(b, tt.name, tt.valueGenerator)
	}
}

func benchmarkRistretto(b *testing.B, prefix string, valueGenerator func(int) string) {
	c, _ := cache.NewCache(&cache.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})

	b.Run(testName(prefix, "Set"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			c.Set(generateKey(prefix, i), valueGenerator(i), 1)
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
			c.Del(generateKey(prefix, i))
		}
		b.StopTimer()
	})
}
