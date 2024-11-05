# Worker Pool Benchmarks

Benchmarks of worker pool libraries for Golang.

## Execute Benchmark

```bash
go test -timeout=5h -bench=. -benchmem -count 5 -benchtime=1000000x > results/results.out
```

## Results

All the [benchmarks](/worker-pool/results/) are performed in the `Intel(R) Core(TM) i7-1165G7 CPU @ 2.80GHz` machine with `100K` samples and `5` iterations.

#### Time / Operation
| Data                                                      | Concurrency                                                        |
| ------------------------------------------------------------------- | --------------------------------------------------------------------- |
| ![data_time_bar.png](/worker-pool/results/data_time_bar.png) | ![concurrency_time_bar.png](/worker-pool/results/concurrency_time_bar.png) |
| ![data_time_table.png](/worker-pool/results/data_time_table.png) | ![concurrency_time_table.png](/worker-pool/results/concurrency_time_table.png) |

#### Allocations / Operation
| Data                                                      | Concurrency                                                        |
| ------------------------------------------------------------------- | --------------------------------------------------------------------- |
| ![data_allocations_bar.png](/worker-pool/results/data_allocations_bar.png) | ![concurrency_allocations_bar.png](/worker-pool/results/concurrency_allocations_bar.png) |
| ![data_allocations_table.png](/worker-pool/results/data_allocations_table.png) | ![concurrency_allocations_table.png](/worker-pool/results/concurrency_allocations_table.png) |

#### Bytes / Operation
| Data                                                      | Concurrency                                                        |
| ------------------------------------------------------------------- | --------------------------------------------------------------------- |
| ![data_memory_bar.png](/worker-pool/results/data_memory_bar.png) | ![concurrency_memory_bar.png](/worker-pool/results/concurrency_memory_bar.png) |
| ![data_memory_table.png](/worker-pool/results/data_memory_table.png) | ![concurrency_memory_table.png](/worker-pool/results/concurrency_memory_table.png) |

## Libraries

:warning: Please note that these libraries are benchmarked against sample execution time of (i.e. 1ms, 10ms, and sha256 calculation of 1KB and 8 KB data). You are encouraged to benchmark with your custom payloads.

- [Alitto/pond](https://github.com/alitto/pond) - Minimalistic and High-performance goroutine worker pool written in Go.
- [Devchat-ai/gopool](https://github.com/devchat-ai/gopool) - GoPool is a high-performance, feature-rich, and easy-to-use worker pool library for Golang.
- [Jeffail/tunny](https://github.com/Jeffail/tunny) - A goroutine pool for Go.
- [Maurice2k/ultrapool](https://github.com/maurice2k/ultrapool) - Blazing fast worker pool for Golang.
- [Panjf2000/ants](https://github.com/panjf2000/ants) - Ants is the most powerful and reliable pooling solution for Go.

