package expressionevaluation

import (
	"testing"

	expr "github.com/Knetic/govaluate"
)

func BenchmarkGoValuate(b *testing.B) {
	for _, tt := range tests {
		benchmarkGoValuate(b, tt.name, tt.input, tt.expression)
	}
}

func benchmarkGoValuate(b *testing.B, prefix string, input Input, expression string) {

	b.Run(testName(prefix, "Compile"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			programGoValuate(b, expression)
		}
		b.StopTimer()
	})

	eval := programGoValuate(b, expression)
	b.Run(testName(prefix, "Evaluate"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			out, err := eval.Eval(input)
			if err != nil && !out.(bool) {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}

func programGoValuate(b *testing.B, expression string) *expr.EvaluableExpression {
	program, err := expr.NewEvaluableExpression(expression)

	if err != nil {
		b.Fatal(err)
	}

	return program
}
