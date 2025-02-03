package json

import (
	"io"
	"testing"

	ugorji "github.com/ugorji/go/codec"
)

func BenchmarkCodec(b *testing.B) {
	for _, tt := range tests {
		benchmarkCodecStruct(b, tt.name, []byte(tt.payload), tt.size)
		benchmarkCodecInterface(b, tt.name, []byte(tt.payload), tt.size)
	}
}

func benchmarkCodecStruct(b *testing.B, prefix string, payload []byte, size int) {
	h := ugorji.JsonHandle{PreferFloat: true, MapKeyAsString: true}
	b.Run(testName(prefix, "Struct@Decode"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]Payload, size)
			dec := ugorji.NewDecoderBytes(payload, &h)
			err := dec.Decode(&payloads)
			if err != nil && err != io.EOF {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}

func benchmarkCodecInterface(b *testing.B, prefix string, payload []byte, size int) {
	h := ugorji.JsonHandle{PreferFloat: true, MapKeyAsString: true}

	b.Run(testName(prefix, "Interface@Decode"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]interface{}, size)
			dec := ugorji.NewDecoderBytes(payload, &h)
			err := dec.Decode(&payloads)
			if err != nil && err != io.EOF {
				b.Fatal(err)
			}
		}

		b.StopTimer()
	})
}
