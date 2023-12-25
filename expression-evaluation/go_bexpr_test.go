package expressionevaluation

import (
	"strings"
	"testing"

	expr "github.com/hashicorp/go-bexpr"
)

func BenchmarkBExpr(b *testing.B) {
	for _, tt := range tests {
		benchmarkBExpr(b, tt.name, tt.input, tt.expression)
	}
}

func benchmarkBExpr(b *testing.B, prefix string, input Input, expression string) {
	expression = strings.ReplaceAll(expression, "(", "")
	expression = strings.ReplaceAll(expression, ")", "")
	expression = strings.ReplaceAll(expression, "||", "and")
	expression = strings.ReplaceAll(expression, "&&", "and")

	b.Run(testName(prefix, "Compile"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := expr.CreateEvaluator(expression)
			if err != nil {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})

	eval, _ := expr.CreateEvaluator(expression)
	b.Run(testName(prefix, "Evaluate"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			out, err := eval.Evaluate(input)
			if err != nil && !out {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}
