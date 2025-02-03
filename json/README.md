# JSON  Benchmarks

Benchmarks of json libraries for Golang.

## Execute Benchmark

```bash
 go test -bench=. -benchmem -timeout=5h -count 5 -benchtime=100000x > results/results.out
```

## Results

All the [benchmarks](/results.out) are performed in the `Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz` machine with `100K` samples and `5` iterations.

### Time / Operation

|Function| Chart View                                                      | Table View                                                        |
|-----| ------------------------------------------------------------------- | --------------------------------------------------------------------- |
|Decode Interface| ![Interface_time_bar](/json/results/Interface_time_bar.png) | ![Interface_time_table](/json/results/Interface_time_table.png) |
|Iterate Object| ![Iterate_time_bar](/json/results/Iterate_time_bar.png) | ![Interface_time_table](/json/results/Iterate_time_table.png) |
|Decode Struct| ![Struct_time_bar](/json/results/Struct_time_bar.png) | ![Interface_time_table](/json/results/Struct_time_table.png) |

#### Allocations / Operation

|Function| Chart View                                                      | Table View                                                        |
|-----| ------------------------------------------------------------------- | --------------------------------------------------------------------- |
|Decode Interface| ![Interface_allocations_bar](/json/results/Interface_allocations_bar.png) | ![Interface_time_table](/json/results/Interface_allocations_table.png) |
|Iterate Object| ![Iterate_allocations_bar](/json/results/Iterate_allocations_bar.png) | ![Interface_time_table](/json/results/Iterate_allocations_table.png) |
|Decode Struct| ![Struct_allocations_bar](/json/results/Struct_allocations_bar.png) | ![Interface_time_table](/json/results/Struct_allocations_table.png) |

#### Bytes / Operation

|Function| Chart View                                                      | Table View                                                        |
|-----| ------------------------------------------------------------------- | --------------------------------------------------------------------- |
|Decode Interface| ![Interface_memory_bar](/json/results/Interface_memory_bar.png) | ![Interface_time_table](/json/results/Interface_memory_table.png) |
|Iterate Object| ![Iterate_memory_bar](/json/results/Iterate_memory_bar.png) | ![Interface_time_table](/json/results/Iterate_memory_table.png) |
|Decode Struct| ![Struct_memory_bar](/json/results/Struct_memory_bar.png) | ![Interface_time_table](/json/results/Struct_memory_table.png) |

## Libraries

:warning: Please note that these libraries are benchmarked against storage of sample payloads (i.e. 1, 5, and 10 KB). You are encouraged to benchmark with your custom payloads.

- [buger/jsonparser](https://github.com/buger/jsonparser) - One of the fastest alternative JSON parser for Go that does not require schema
- [bytedance/sonic](https://github.com/bytedance/sonic) - A blazingly fast JSON serializing & deserializing library
- [encoding/json](https://pkg.go.dev/encoding/json) -  Implements encoding and decoding of JSON as defined in RFC 7159
- [goccy/go-json](https://github.com/goccy/go-json) - Fast JSON encoder/decoder compatible with encoding/json for Go
- [json-iterator](https://github.com/json-iterator/go) - A high-performance 100% compatible drop-in replacement of "encoding/json"
- [mailru/easyjson](https://github.com/mailru/easyjson) - Fast JSON serializer for golang.
- [segmentio/encoding](https://github.com/segmentio/encoding) - Go package containing implementations of efficient encoding, decoding, and validation APIs.
- [sugawarayuuta/sonnet](https://github.com/sugawarayuuta/sonnet) - High performance JSON decoder in Go
- [tidwall/gjson](https://github.com/tidwall/gjson) - Get JSON values quickly - JSON parser for Go
- [ugorji/go](https://github.com/ugorji/go) - idiomatic codec and rpc lib for msgpack, cbor, json, etc. msgpack.org
- [zerosnake0/jzon](https://github.com/zerosnake0/jzon) - A golang json library inspired by json-iterator

## Credits

- Test data is generated using [mockaroo](https://www.mockaroo.com/)
