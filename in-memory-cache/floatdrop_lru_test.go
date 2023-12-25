package inmemorycache

import (
	"testing"

	cache "github.com/floatdrop/lru"
)

func BenchmarkFloatdropLRU(b *testing.B) {
	for _, tt := range tests {
		benchmarkFloatdropLRU(b, tt.name, tt.valueGenerator)
	}
}

func benchmarkFloatdropLRU(b *testing.B, prefix string, valueGenerator func(int) string) {
	c := cache.New[string, string](size * 1000)

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
