# In-Memory Cache Benchmarks

Benchmarks of in-memory cache libraries for Golang.

## Execute Benchmark

```bash
 go test -bench=. -benchmem -count 5 -benchtime=100000x > results/results.out
```

## Results

All the [benchmarks](/results.out) are performed in the `Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz` machine with `100K` samples and `5` iterations.

![Average](/in-memory-cache/results/Average_Cache.png)

### Average ns / operation

#### Set Function

![Average_ns_per_operation_set.png](/in-memory-cache/results/Average_ns_per_operation_set.png)

#### Get & Remove Function

![Average_ns_per_operation_get_remove](/in-memory-cache/results/Average_ns_per_operation_get_remove.png)

## Libraries

:warning: Please note that these libraries are benchmarked against storage of sample payloads (i.e. 1, 5, and 10 KB). You are encouraged to benchmark with your custom payloads.

- [akyoto/cache](https://github.com/akyoto/cache) - Cache arbitrary data with an expiration time.
- [allegro/bigcache](https://github.com/allegro/bigcache) - Efficient cache for gigabytes of data written in Go.
- [bluele/gcache](https://github.com/bluele/gcache) - An in-memory cache library for golang. It supports multiple eviction policies: LRU, LFU, ARC.
- [coocood/freecache](https://github.com/coocood/freecache) - A cache library for Go with zero GC overhead.
- [dgraph-io/ristretto](https://github.com/dgraph-io/ristretto) - A high performance memory-bound Go cache.
- [floatdrop/2q](https://github.com/floatdrop/2q) - Thread safe GoLang [2Q](https://www.vldb.org/conf/1994/P439.PDF) cache.
- [floatdrop/lru](https://github.com/floatdrop/lru) - Thread safe GoLang LRU cache
- [floatdrop/slru](https://github.com/floatdrop/slru) - Thread safe GoLang S(2)LRU cache.
- [hashicorp/golang-lru](https://github.com/hashicorp/golang-lru) - Golang LRU cache
- [jellydator/ttlcache](https://github.com/jellydator/ttlcache) - An in-memory cache with item expiration and generics
- [koding/cache](https://github.com/koding/cache) - Caching package for Go
- [irr123/wtfcache](https://github.com/irr123/wtfcache) - Threadsafe cache with generic interface
- [muesli/cache2go](https://github.com/muesli/cache2go) - Concurrency-safe Go caching library with expiration capabilities and access counters.
- [patrickmn/go-cache](https://github.com/patrickmn/go-cache) - An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.
- [sync#Map](https://pkg.go.dev/sync#Map) - Safe for concurrent use by multiple goroutine without additional locking or coordination. Loads, stores, and deletes run in amortized constant time.
- [VictoriaMetrics/fastcache](https://github.com/VictoriaMetrics/fastcache) - Fast thread-safe in memory cache for big number of entries in Go. Minimizes GC overhead.
  
## Credits

- Test data is generated using [mockaroo](https://www.mockaroo.com/)
