package expressionevaluation

import (
	"testing"

	expr "github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
)

func BenchmarkExpr(b *testing.B) {
	for _, tt := range tests {
		benchmarkExpr(b, tt.name, tt.input, tt.expression)
	}
}

func benchmarkExpr(b *testing.B, prefix string, input Input, expression string) {

	b.Run(testName(prefix, "Compile"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			programExpr(b, input, expression)
		}
		b.StopTimer()
	})

	eval := programExpr(b, input, expression)
	b.Run(testName(prefix, "Evaluate"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			out, err := expr.Run(eval, input)
			if err != nil && !out.(bool) {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}

func programExpr(b *testing.B, input Input, expression string) *vm.Program {
	program, err := expr.Compile(expression, expr.Env(input))
	if err != nil {
		b.Fatal(err)
	}

	return program
}
