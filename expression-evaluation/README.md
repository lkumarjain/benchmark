# Expression Evaluation Benchmarks

Benchmarks of expression evaluation libraries for Golang.

## Execute Benchmark

```bash
 go test -bench=. -benchmem -count 5 -benchtime=100000x > results/results.out
```

## Results

All the [benchmarks](/results.out) are performed in the `Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz` machine with `10K` samples and `5` iterations.

### Compile

Average Compile time for each expression library

![Average Compile](/expression-evaluation/results/Average_Compile.png)
![Average Compile Expand](/expression-evaluation/results/Average_Compile_Expand.png)

### Evaluation

Average Evaluation time for each expression library

![Average Evaluate](/expression-evaluation/results/Average_Evaluate.png)
![Average Evaluate Expand](/expression-evaluation/results/Average_Evaluate_Expand.png)

## Libraries

:warning: Please note that these libraries are benchmarked against sample expressions. You are encouraged to benchmark with your custom expressions.

- [antonmedv/expr](https://github.com/antonmedv/expr) - Expression language and expression evaluation for Go
- [d5/tengo](https://github.com/d5/tengo) - A fast script language for Go.
- [dop251/goja](https://github.com/dop251/goja) - ECMAScript/JavaScript engine in pure Go.
- [google/cel-go](https://github.com/google/cel-go) - Fast, portable, non-Turing complete expression evaluation with gradual typing (Go)
- [hashicorp/go-bexpr](https://github.com/hashicorp/go-bexpr) - Generic boolean expression evaluation in Go.
- [Knetic/govaluate](https://github.com/Knetic/govaluate) - Arbitrary expression evaluation for golang.
- [PaesslerAG/gval](https://github.com/PaesslerAG/gval) - Expression evaluation in golang.
- [robertkrimen/otto](https://github.com/robertkrimen/otto) - A JavaScript interpreter in Go (golang).
- [skx/evalfilter](https://github.com/skx/evalfilter) - A byte-code based virtual machine to implement scripting/filtering support in your golang project.
- [google/starlark-go](https://github.com/google/starlark-go) - Starlark in Go: The Starlark configuration language, implemented in Go.
  
## Credits

- Test data is generated using [mockaroo](https://www.mockaroo.com/)
