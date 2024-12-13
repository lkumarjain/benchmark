package synchronizationtechniques

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
)

func BenchmarkAtomic(b *testing.B) {
	var value int64
	for _, concur := range concurrency {
		runtime.GC()

		b.Run(fmt.Sprintf("%s/Parallelism=%d", "Atomic", concur), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			b.SetParallelism(concur)

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					atomic.AddInt64(&value, 1)
				}
			})

			b.StopTimer()
		})
	}
}
