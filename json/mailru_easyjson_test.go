package json

import (
	"testing"

	"github.com/mailru/easyjson/jlexer"
)

func BenchmarkEasyJson(b *testing.B) {
	for _, tt := range tests {
		benchmarkEasyJsonInterface(b, tt.name, []byte(tt.payload))
	}
}

func benchmarkEasyJsonInterface(b *testing.B, prefix string, payload []byte) {
	b.Run(testName(prefix, "Interface@Decode"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			l := jlexer.Lexer{Data: payload}
			l.Interface()
		}

		b.StopTimer()
	})
}
