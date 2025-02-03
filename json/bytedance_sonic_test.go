package json

import (
	"bytes"
	"testing"

	encoding "github.com/bytedance/sonic"
	"github.com/bytedance/sonic/decoder"
)

func BenchmarkSonic(b *testing.B) {
	b.ReportAllocs()
	for _, tt := range tests {
		benchmarkSonicStruct(b, tt.name, []byte(tt.payload), tt.size)
		benchmarkSonicInterface(b, tt.name, []byte(tt.payload), tt.size)
	}
}

func benchmarkSonicStruct(b *testing.B, prefix string, payload []byte, size int) {
	b.Run(testName(prefix, "Struct@Decode"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]Payload, size)
			dec := decoder.NewStreamDecoder(bytes.NewReader(payload))
			j := 0
			for {
				var result Payload
				if err := dec.Decode(&result); err != nil {
					break
				}
				payloads[j] = result
				j++
			}
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Struct@Unmarshal"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]Payload, size)
			if err := encoding.Unmarshal(payload, &payloads); err != nil {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}

func benchmarkSonicInterface(b *testing.B, prefix string, payload []byte, size int) {
	b.Run(testName(prefix, "Interface@Decode"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]interface{}, size)
			dec := decoder.NewStreamDecoder(bytes.NewReader(payload))
			j := 0
			for {
				var result interface{}
				if err := dec.Decode(&result); err != nil {
					break
				}
				payloads[j] = result
				j++
			}
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Interface@Unmarshal"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]interface{}, size)
			if err := encoding.Unmarshal(payload, &payloads); err != nil {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}
