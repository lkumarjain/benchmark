package inmemorycache

import (
	"testing"

	cache "github.com/irr123/wtfcache"
)

func BenchmarkWTF(b *testing.B) {
	for _, tt := range tests {
		benchmarkWTF(b, tt.name, tt.valueGenerator)
	}
}

func benchmarkWTF(b *testing.B, prefix string, valueGenerator func(int) string) {
	c := cache.New[string, string]().MakeWithLock(b.N)

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
			c.Del(generateKey(prefix, i))
		}
		b.StopTimer()
	})
}
