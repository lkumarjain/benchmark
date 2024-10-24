package workerpool

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"

	"github.com/devchat-ai/gopool"
)

func BenchmarkGoPool(b *testing.B) {
	for _, test := range tests {

		runtime.GC()

		for _, concur := range concurrency {
			b.Run(fmt.Sprintf("%s_Concurrency=%d", test.name, concur), func(b *testing.B) {

				mu := &sync.Mutex{}
				response := make([]interface{}, 0)

				wp := gopool.NewGoPool(maxPoolSize, gopool.WithResultCallback(func(result interface{}) {
					mu.Lock()
					response = append(response, result)
					mu.Unlock()
				}))

				defer wp.Release()

				b.ReportAllocs()
				b.SetParallelism(concur)
				b.ResetTimer()

				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						wg.Add(1)
						wp.AddTask(func() (interface{}, error) {
							defer wg.Done()
							return test.executor(context.Background(), "1")
						})
					}
				})

				wp.Wait()
				wg.Wait()

				b.StopTimer()
			})
		}
	}
}
