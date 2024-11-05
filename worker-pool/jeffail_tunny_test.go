package workerpool

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"

	"github.com/Jeffail/tunny"
)

func BenchmarkTunny(b *testing.B) {
	for _, test := range tests {

		runtime.GC()

		for _, concur := range concurrency {
			b.Run(fmt.Sprintf("%s/Concurrency=%d", test.name, concur), func(b *testing.B) {
				mu := &sync.Mutex{}
				response := make([]interface{}, 0)

				job := func(_ interface{}) interface{} {
					defer wg.Done()
					s, _ := test.executor(context.Background(), "1")

					return s
				}

				wp := tunny.NewFunc(runtime.GOMAXPROCS(0), job)
				wp.SetSize(50)

				b.ReportAllocs()
				b.SetParallelism(concur)
				b.ResetTimer()

				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						wg.Add(1)
						s := wp.Process("")

						mu.Lock()
						response = append(response, s)
						mu.Unlock()
					}
				})

				wp.Close()
				wg.Wait()

				b.StopTimer()
			})
		}
	}
}
