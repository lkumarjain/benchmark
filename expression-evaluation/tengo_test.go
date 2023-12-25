package expressionevaluation

import (
	"context"
	"fmt"
	"testing"

	expr "github.com/d5/tengo/v2"
)

func BenchmarkTengo(b *testing.B) {
	for _, tt := range tests {
		benchmarkTengo(b, tt.name, tt.input, tt.expression)
	}
}

func benchmarkTengo(b *testing.B, prefix string, input Input, expression string) {
	script := expr.NewScript([]byte(fmt.Sprintf("result := (%s) ? true : false", expression)))

	script.Add("ID", input.ID)
	script.Add("Name", input.Name)
	script.Add("City", input.City)
	script.Add("Country", input.Country)
	script.Add("Currency", input.Currency)

	b.Run(testName(prefix, "Compile"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := script.RunContext(context.Background())
			if err != nil {
				panic(err)
			}
		}
		b.StopTimer()
	})

	eval, _ := script.RunContext(context.Background())

	b.Run(testName(prefix, "Evaluate"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			value := eval.Get("result")
			if !value.Bool() {
				b.Fatal("Error")
			}
		}
		b.StopTimer()
	})
}
