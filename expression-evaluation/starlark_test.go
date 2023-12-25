package expressionevaluation

import (
	"strings"
	"testing"

	expr "go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

func BenchmarkStarlark(b *testing.B) {
	for _, tt := range tests {
		benchmarkStarlark(b, tt.name, tt.input, tt.expression)
	}
}

func benchmarkStarlark(b *testing.B, prefix string, input Input, expression string) {
	expression = strings.ReplaceAll(expression, "||", "and")
	expression = strings.ReplaceAll(expression, "&&", "and")

	b.Run(testName(prefix, "Compile"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			programStarlark(b, expression)
		}
		b.StopTimer()
	})

	thread := &expr.Thread{Name: "example"}
	predeclared := expr.StringDict{
		"ID":       expr.MakeInt(input.ID),
		"Name":     expr.String(input.Name),
		"City":     expr.String(input.City),
		"Country":  expr.String(input.Country),
		"Currency": expr.String(input.Currency),
	}

	eval := programStarlark(b, expression)
	b.Run(testName(prefix, "Evaluate"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			out, err := expr.EvalExpr(thread, eval, predeclared)
			if err != nil && !out.(expr.Bool).Truth() {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}

func programStarlark(b *testing.B, expression string) syntax.Expr {
	program, err := syntax.ParseExpr("example.star", expression, syntax.RetainComments)
	if err != nil {
		b.Fatal(err)
	}

	if err != nil {
		b.Fatal(err)
	}

	return program
}
