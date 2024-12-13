package expressionevaluation

import (
	"fmt"
	"testing"

	expr "github.com/skx/evalfilter/v2"
)

func BenchmarkEvalFilter(b *testing.B) {
	for _, tt := range tests {
		benchmarkEvalFilter(b, tt.name, tt.input, tt.expression)
	}
}

func benchmarkEvalFilter(b *testing.B, prefix string, input Input, expression string) {
	expression = fmt.Sprintf(`if ( %s ) { return true; } else { return false; }`, expression)

	b.Run(testName(prefix, "Compile"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			programEvalFilter(b, input, expression)
		}
		b.StopTimer()
	})

	eval := programEvalFilter(b, input, expression)
	b.Run(testName(prefix, "Evaluate"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			out, err := eval.Run(input)
			if err != nil && !out {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}

func programEvalFilter(b *testing.B, _ Input, expression string) *expr.Eval {
	eval := expr.New(expression)

	err := eval.Prepare()
	if err != nil {
		b.Fatal(err)
	}

	return eval
}
