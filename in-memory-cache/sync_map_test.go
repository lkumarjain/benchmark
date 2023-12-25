package inmemorycache

import (
	"sync"
	"testing"
)

func BenchmarkSyncMap(b *testing.B) {
	for _, tt := range tests {
		benchmarkSyncMap(b, tt.name, tt.valueGenerator)
	}
}

func benchmarkSyncMap(b *testing.B, prefix string, valueGenerator func(int) string) {
	var m sync.Map

	b.Run(testName(prefix, "Set"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Store(generateKey(prefix, i), valueGenerator(i))
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Get"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Load(generateKey(prefix, i))
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Remove"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Delete(generateKey(prefix, i))
		}
		b.StopTimer()
	})
}
