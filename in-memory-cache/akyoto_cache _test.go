package inmemorycache

import (
	"testing"
	"time"

	"github.com/akyoto/cache"
)

func BenchmarkAkyotoCache(b *testing.B) {
	for _, tt := range tests {
		benchmarkAkyotoCache(b, tt.name, tt.valueGenerator)
	}
}

func benchmarkAkyotoCache(b *testing.B, prefix string, valueGenerator func(int) string) {
	c := cache.New(time.Minute)

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
