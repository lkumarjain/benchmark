package expressionevaluation

import (
	"testing"

	expr "github.com/robertkrimen/otto"
)

func BenchmarkOtto(b *testing.B) {
	for _, tt := range tests {
		benchmarkOtto(b, tt.name, tt.input, tt.expression)
	}
}

func benchmarkOtto(b *testing.B, prefix string, input Input, expression string) {

	vm := expr.New()

	b.Run(testName(prefix, "Compile"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			programOtto(b, vm, expression)
		}
		b.StopTimer()
	})

	vm.Set("ID", input.ID)
	vm.Set("Name", input.Name)
	vm.Set("City", input.City)
	vm.Set("Country", input.Country)
	vm.Set("Currency", input.Currency)

	eval := programOtto(b, vm, expression)
	b.Run(testName(prefix, "Evaluate"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			out, err := vm.Run(eval)
			if err != nil {
				b.Fatal(err)
			}

			ok, err := out.ToBoolean()
			if err != nil && !ok {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}

func programOtto(b *testing.B, vm *expr.Otto, expression string) *expr.Script {
	program, err := vm.Compile("", expression)

	if err != nil {
		b.Fatal(err)
	}

	return program
}
