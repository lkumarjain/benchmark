package workerpool

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/maurice2k/ultrapool"
)

func BenchmarkUltraPool(b *testing.B) {
	for _, test := range tests {

		runtime.GC()

		for _, concur := range concurrency {
			b.Run(fmt.Sprintf("%s_Concurrency=%d", test.name, concur), func(b *testing.B) {
				job := func(_ ultrapool.Task) {
					defer wg.Done()

					test.executor(context.Background(), "1")
				}

				wp := ultrapool.NewWorkerPool(job)
				wp.SetIdleWorkerLifetime(time.Second * 5)
				wp.Start()

				b.ReportAllocs()
				b.SetParallelism(concur)
				b.ResetTimer()

				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						wg.Add(1)
						wp.AddTask("")
					}
				})

				wp.Stop()
				wg.Wait()

				b.StopTimer()
			})
		}
	}
}
