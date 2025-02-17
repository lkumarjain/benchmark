# Kafka Client Benchmarks

Benchmarks of sql-builder libraries for Golang.

## Execute Benchmark

```bash
go test -timeout=5h -bench=, -benchmem -count 5 -benchtime=100000x > results/results.out
```

## Results

All the [benchmarks](/results/results.out) are performed in the `Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz` machine with `1M` samples and `5` iterations where lower values are good.

### Time / Operation

|Function| Chart View                                                      | Table View                                                        |
|-----| ------------------------------------------------------------------- | --------------------------------------------------------------------- |
|Insert| ![Insert_time_bar.png](/sql-builder/results/Insert_time_bar.png) | ![Insert_time_table.png](/sql-builder/results/Insert_time_table.png) |
|Select| ![Select_time_bar.png](/sql-builder/results/Select_time_bar.png) | ![Select_time_table.png](/sql-builder/results/Select_time_table.png) |
|Update| ![Update_time_bar.png](/sql-builder/results/Update_time_bar.png) | ![Update_time_table.png](/sql-builder/results/Update_time_table.png) |
|Delete| ![Delete_time_bar.png](/sql-builder/results/Delete_time_bar.png) | ![Delete_time_table.png](/sql-builder/results/Delete_time_table.png) |

#### Allocations / Operation

|Function| Chart View                                                      | Table View                                                        |
|-----| ------------------------------------------------------------------- | --------------------------------------------------------------------- |
|Insert| ![Insert_memory_bar.png](/sql-builder/results/Insert_memory_bar.png) | ![Insert_memory_table.png](/sql-builder/results/Insert_memory_table.png) |
|Select| ![Select_memory_bar.png](/sql-builder/results/Select_memory_bar.png) | ![Select_memory_table.png](/sql-builder/results/Select_memory_table.png) |
|Update| ![Update_memory_bar.png](/sql-builder/results/Update_memory_bar.png) | ![Update_memory_table.png](/sql-builder/results/Update_memory_table.png) |
|Delete| ![Delete_memory_bar.png](/sql-builder/results/Delete_memory_bar.png) | ![Delete_memory_table.png](/sql-builder/results/Delete_memory_table.png) |

#### Bytes / Operation

|Function| Chart View                                                      | Table View                                                        |
|-----| ------------------------------------------------------------------- | --------------------------------------------------------------------- |
|Insert| ![Insert_allocations_bar.png](/sql-builder/results/Insert_allocations_bar.png) | ![Insert_allocations_table.png](/sql-builder/results/Insert_allocations_table.png) |
|Select| ![Select_allocations_bar.png](/sql-builder/results/Select_allocations_bar.png) | ![Select_allocations_table.png](/sql-builder/results/Select_allocations_table.png) |
|Update| ![Update_allocations_bar.png](/sql-builder/results/Update_allocations_bar.png) | ![Update_allocations_table.png](/sql-builder/results/Update_allocations_table.png) |
|Delete| ![Delete_allocations_bar.png](/sql-builder/results/Delete_allocations_bar.png) | ![Delete_allocations_table.png](/sql-builder/results/Delete_allocations_table.png) |

## Libraries

:warning: Please note that these libraries are benchmarked against storage of sample payloads (i.e. 1, 5, and 10 KB). You are encouraged to benchmark with your custom payloads.

- [Doug-martin/Goqu] (https://github.com/doug-martin/goqu) - SQL builder and query library for golang
- [Gocraft/Dbr](https://github.com/gocraft/dbr) - Additions to Go's database/sql for super fast performance and convenience.
- [Masterminds/Squirrel](https://github.com/Masterminds/squirrel) - Fluent SQL generation for golang
- [Stephenafamo/Bob](https://github.com/stephenafamo/bob) - SQL query builder and ORM/Factory generator for Go with support for PostgreSQL, MySQL and SQLite.
