package expressionevaluation

import (
	"testing"

	expr "github.com/google/cel-go/cel"
)

func BenchmarkCelGO(b *testing.B) {
	for _, tt := range tests {
		benchmarkCelGO(b, tt.name, tt.input, tt.expression)
	}
}

func benchmarkCelGO(b *testing.B, prefix string, input Input, expression string) {

	b.Run(testName(prefix, "Compile"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			programCelGO(b, input, expression)
		}
		b.StopTimer()
	})

	eval := programCelGO(b, input, expression)
	b.Run(testName(prefix, "Evaluate"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			out, _, err := eval.Eval(input)
			if err != nil && !out.Value().(bool) {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}

func programCelGO(b *testing.B, input Input, expression string) expr.Program {
	env, err := expr.NewEnv(
		expr.Variable("ID", expr.IntType),
		expr.Variable("Name", expr.StringType),
		expr.Variable("City", expr.StringType),
		expr.Variable("Country", expr.StringType),
		expr.Variable("Currency", expr.StringType),
	)

	if err != nil {
		b.Fatal(err)
	}

	parsed, issues := env.Parse(expression)
	if issues != nil && issues.Err() != nil {
		b.Fatalf("parse error: %s", issues.Err())
	}

	checked, issues := env.Check(parsed)
	if issues != nil && issues.Err() != nil {
		b.Fatalf("type-check error: %s", issues.Err())
	}

	prg, err := env.Program(checked)
	if err != nil {
		b.Fatalf("program construction error: %s", err)
	}

	return prg
}
