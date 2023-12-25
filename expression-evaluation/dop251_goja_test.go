package expressionevaluation

import (
	"testing"

	expr "github.com/dop251/goja"
)

func BenchmarkGoja(b *testing.B) {
	for _, tt := range tests {
		benchmarkGoja(b, tt.name, tt.input, tt.expression)
	}
}

func benchmarkGoja(b *testing.B, prefix string, input Input, expression string) {

	b.Run(testName(prefix, "Compile"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			programGoja(b, expression)
		}
		b.StopTimer()
	})

	vm := expr.New()
	vm.Set("ID", input.ID)
	vm.Set("Name", input.Name)
	vm.Set("City", input.City)
	vm.Set("Country", input.Country)
	vm.Set("Currency", input.Currency)

	eval := programGoja(b, expression)
	b.Run(testName(prefix, "Evaluate"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			out, err := vm.RunProgram(eval)
			if err != nil && !out.ToBoolean() {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}

func programGoja(b *testing.B, expression string) *expr.Program {
	program, err := expr.Compile("", expression, false)
	if err != nil {
		b.Fatal(err)
	}
	return program
}
