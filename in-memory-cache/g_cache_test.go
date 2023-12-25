package inmemorycache

import (
	"testing"

	cache "github.com/bluele/gcache"
)

func BenchmarkGCache(b *testing.B) {
	for _, tt := range tests {
		benchmarkGCache(b, tt.name, tt.valueGenerator)
	}
}

func benchmarkGCache(b *testing.B, prefix string, valueGenerator func(int) string) {
	c := cache.New(1024 * 1024 * size).LRU().Build()

	b.Run(testName(prefix, "Set"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			c.Set(generateKey(prefix, i), valueGenerator(i))
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
			c.Remove(generateKey(prefix, i))
		}
		b.StopTimer()
	})
}
