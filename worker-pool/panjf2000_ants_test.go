package workerpool

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/panjf2000/ants/v2"
)

func BenchmarkAnts(b *testing.B) {
	for _, test := range tests {

		runtime.GC()

		for _, concur := range concurrency {
			b.Run(fmt.Sprintf("%s/Concurrency=%d", test.name, concur), func(b *testing.B) {
				job := func(_ interface{}) {
					defer wg.Done()

					test.executor(context.Background(), "1")
				}

				wp, _ := ants.NewPoolWithFunc(50, job, ants.WithPreAlloc(false), ants.WithExpiryDuration(time.Second*5))

				defer wp.Release()

				b.ReportAllocs()
				b.SetParallelism(concur)
				b.ResetTimer()

				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						wg.Add(1)
						wp.Invoke("")
					}
				})

				wg.Wait()

				b.StopTimer()
			})
		}
	}
}
