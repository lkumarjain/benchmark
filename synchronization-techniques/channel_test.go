package synchronizationtechniques

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkChannel(b *testing.B) {
	var value int64

	for _, concur := range concurrency {
		runtime.GC()
		var ch = make(chan struct{}, 1)

		b.Run(fmt.Sprintf("%s/Parallelism=%d", "Channel", concur), func(b *testing.B) {
			b.ResetTimer()
			b.SetParallelism(100)

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					ch <- struct{}{}
					value++
					<-ch
				}
			})

			b.StopTimer()
		})
	}
}
