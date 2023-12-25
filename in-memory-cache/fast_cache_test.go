package inmemorycache

import (
	"testing"

	cache "github.com/VictoriaMetrics/fastcache"
)

func BenchmarkFastCache(b *testing.B) {
	for _, tt := range tests {
		benchmarkFastCache(b, tt.name, tt.valueGenerator)
	}
}

func benchmarkFastCache(b *testing.B, prefix string, valueGenerator func(int) string) {
	c := cache.New(size * 1000)

	b.Run(testName(prefix, "Set"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			c.Set([]byte(generateKey(prefix, i)), []byte(valueGenerator(i)))
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Get"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			c.Get(nil, []byte(generateKey(prefix, i)))
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
