package expressionevaluation

import (
	"context"
	"testing"

	expr "github.com/PaesslerAG/gval"
)

func BenchmarkGVAL(b *testing.B) {
	for _, tt := range tests {
		benchmarkGVAL(b, tt.name, tt.input, tt.expression)
	}
}

func benchmarkGVAL(b *testing.B, prefix string, input Input, expression string) {

	b.Run(testName(prefix, "Compile"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			programGVAL(b, expression)
		}
		b.StopTimer()
	})

	eval := programGVAL(b, expression)
	b.Run(testName(prefix, "Evaluate"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			out, err := eval(context.Background(), input)
			if err != nil && !out.(bool) {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}

func programGVAL(b *testing.B, expression string) expr.Evaluable {
	program, err := expr.Full().NewEvaluable(expression)

	if err != nil {
		b.Fatal(err)
	}

	return program
}
