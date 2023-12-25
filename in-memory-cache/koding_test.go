package inmemorycache

import (
	"testing"
	"time"

	"github.com/koding/cache"
)

func BenchmarkKoding(b *testing.B) {
	for _, tt := range tests {
		benchmarkKoding(b, tt.name, tt.valueGenerator)
	}
}

func benchmarkKoding(b *testing.B, prefix string, valueGenerator func(int) string) {
	c := cache.NewMemoryWithTTL(time.Duration(60) * time.Second)

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
			c.Delete(generateKey(prefix, i))
		}
		b.StopTimer()
	})
}
