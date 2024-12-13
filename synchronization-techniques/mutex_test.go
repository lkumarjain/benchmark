package synchronizationtechniques

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	var value int64
	for _, concur := range concurrency {
		runtime.GC()
		var mutex sync.Mutex

		b.Run(fmt.Sprintf("%s/Parallelism=%d", "Mutex", concur), func(b *testing.B) {
			b.ResetTimer()
			b.SetParallelism(concur)

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					mutex.Lock()
					value++
					mutex.Unlock()
				}
			})

			b.StopTimer()
		})
	}
}
