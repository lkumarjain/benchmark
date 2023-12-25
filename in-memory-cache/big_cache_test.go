package inmemorycache

import (
	"testing"
	"time"

	cache "github.com/allegro/bigcache"
)

func BenchmarkBigCache(b *testing.B) {
	for _, tt := range tests {
		benchmarkBigCache(b, tt.name, tt.valueGenerator)
	}
}

func benchmarkBigCache(b *testing.B, prefix string, valueGenerator func(int) string) {
	c, _ := cache.NewBigCache(cache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,
		// time after which entry can be evicted
		LifeWindow: 10 * time.Minute,
		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,
		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,
		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 10,
	})

	b.Run(testName(prefix, "Set"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			c.Set(generateKey(prefix, i), []byte(valueGenerator(i)))
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
