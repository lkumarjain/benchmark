# Worker Pool Benchmarks

Benchmarks of Synchronization Techniques for Golang.

## Execute Benchmark

```bash
go test -timeout=5h -bench=. -benchmem -count 5 -benchtime=1000000x > results/results.out
```

## Results

All the [benchmarks](/synchronization-techniques/results/) are performed in the `Intel(R) Core(TM) i7-1165G7 CPU @ 2.80GHz` machine with `100K` samples and `5` iterations.

#### Time / Operation
![concurrency_time_bar.png](/synchronization-techniques/results/concurrency_time_bar.png)
![concurrency_time_table.png](/synchronization-techniques/results/concurrency_time_table.png)

#### Allocations / Operation
![concurrency_allocations_bar.png](/synchronization-techniques/results/concurrency_allocations_bar.png)
![concurrency_allocations_table.png](/synchronization-techniques/results/concurrency_allocations_table.png)

#### Bytes / Operation
![concurrency_memory_bar.png](/synchronization-techniques/results/concurrency_memory_bar.png)
![concurrency_memory_table.png](/synchronization-techniques/results/concurrency_memory_table.png)

## Libraries

:warning: Please note that these techniques are benchmarked against a number increment. You are encouraged to benchmark with your custom payloads.

- [Atomic](https://pkg.go.dev/sync/atomic) - Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.
- [Channel](https://go.dev/tour/concurrency/2) - Channels are a typed conduit through which you can send and receive values.
- [Mutex](https://pkg.go.dev/sync#Mutex) - A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex..
