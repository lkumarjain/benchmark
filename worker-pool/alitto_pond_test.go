package workerpool

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"

	"github.com/alitto/pond/v2"
)

func BenchmarkPond(b *testing.B) {
	for _, test := range tests {

		runtime.GC()

		for _, concur := range concurrency {
			b.Run(fmt.Sprintf("%s/Concurrency=%d", test.name, concur), func(b *testing.B) {
				mu := &sync.Mutex{}
				wp := pond.NewResultPool[string](100)
				results := make([]pond.Result[string], 0)

				b.ReportAllocs()
				b.SetParallelism(concur)
				b.ResetTimer()

				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						wg.Add(1)
						r := wp.Submit(func() string {
							defer wg.Done()

							s, _ := test.executor(context.Background(), "1")

							return s
						})

						mu.Lock()
						results = append(results, r)
						mu.Unlock()
					}
				})

				wp.StopAndWait()
				wg.Wait()

				response := make([]string, len(results))
				for i, v := range results {
					s, _ := v.Wait()
					response[i] = s
				}

				b.StopTimer()
			})
		}
	}
}
