package inmemorycache

import (
	"testing"

	cache "github.com/coocood/freecache"
)

func BenchmarkFreecache(b *testing.B) {
	for _, tt := range tests {
		benchmarkFreecache(b, tt.name, tt.valueGenerator)
	}
}

func benchmarkFreecache(b *testing.B, prefix string, valueGenerator func(int) string) {
	c := cache.NewCache(1024 * 1024 * size)

	b.Run(testName(prefix, "Set"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			c.Set([]byte(generateKey(prefix, i)), []byte(valueGenerator(i)), 60)
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Get"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			c.Get([]byte(generateKey(prefix, i)))
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Remove"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			c.Del([]byte(generateKey(prefix, i)))
		}
		b.StopTimer()
	})
}
