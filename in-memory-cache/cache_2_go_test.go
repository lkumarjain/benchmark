package inmemorycache

import (
	"testing"
	"time"

	"github.com/muesli/cache2go"
)

func BenchmarkCache2Go(b *testing.B) {
	for _, tt := range tests {
		benchmarkCache2Go(b, tt.name, tt.valueGenerator)
	}
}

func benchmarkCache2Go(b *testing.B, prefix string, valueGenerator func(int) string) {
	c := cache2go.Cache("test")

	b.Run(testName(prefix, "Set"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			c.Add(generateKey(prefix, i), time.Minute, valueGenerator(i))
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Get"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			c.Value(generateKey(prefix, i))
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
