package json

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func BenchmarkJsoniter(b *testing.B) {
	b.ReportAllocs()
	for _, tt := range tests {
		benchmarkJsoniterIterate(b, tt.name, []byte(tt.payload))
		benchmarkJsoniterStruct(b, tt.name, []byte(tt.payload), tt.size)
		benchmarkJsoniterInterface(b, tt.name, []byte(tt.payload), tt.size)
	}
}

func benchmarkJsoniterIterate(b *testing.B, prefix string, payload []byte) {
	b.Run(testName(prefix, "Iterate@Read"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			jit := jsoniter.ConfigDefault.BorrowIterator(payload)
			jit.ReadArrayCB(func(i *jsoniter.Iterator) bool {
				i.ReadObjectCB(func(i *jsoniter.Iterator, field string) bool {
					switch field {
					case idStr:
						i.ReadInt()
					case nameStr:
						i.ReadString()
					case cityStr:
						i.ReadString()
					case countryStr:
						i.ReadString()
					case currencyStr:
						i.ReadString()
					}
					return true
				})
				return true
			})
			jsoniter.ConfigDefault.ReturnIterator(jit)
		}
		b.StopTimer()
	})
}

func benchmarkJsoniterStruct(b *testing.B, prefix string, payload []byte, size int) {
	b.Run(testName(prefix, "Struct@Read"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]Payload, size)
			j := 0
			jit := jsoniter.ConfigDefault.BorrowIterator(payload)
			jit.ReadArrayCB(func(i *jsoniter.Iterator) bool {
				result := Payload{}
				i.ReadObjectCB(func(i *jsoniter.Iterator, field string) bool {
					switch field {
					case idStr:
						result.ID = i.ReadInt()
					case nameStr:
						result.Name = i.ReadString()
					case cityStr:
						result.City = i.ReadString()
					case countryStr:
						result.Country = i.ReadString()
					case currencyStr:
						result.Currency = i.ReadString()
					}
					return true
				})
				payloads[j] = result
				j++
				return true
			})
			jsoniter.ConfigDefault.ReturnIterator(jit)
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Struct@Unmarshal"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]Payload, size)
			if err := jsoniter.ConfigDefault.Unmarshal(payload, &payloads); err != nil {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}

func benchmarkJsoniterInterface(b *testing.B, prefix string, payload []byte, size int) {
	b.Run(testName(prefix, "Interface@Read"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]interface{}, size)
			j := 0
			jit := jsoniter.ConfigDefault.BorrowIterator(payload)
			jit.ReadArrayCB(func(i *jsoniter.Iterator) bool {
				result := map[string]interface{}{}
				i.ReadObjectCB(func(i *jsoniter.Iterator, field string) bool {
					switch field {
					case idStr:
						result[idStr] = i.ReadInt()
					case nameStr:
						result[nameStr] = i.ReadString()
					case cityStr:
						result[cityStr] = i.ReadString()
					case countryStr:
						result[countryStr] = i.ReadString()
					case currencyStr:
						result[currencyStr] = i.ReadString()
					}
					return true
				})
				payloads[j] = result
				j++
				return true
			})
			jsoniter.ConfigDefault.ReturnIterator(jit)
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Interface@Unmarshal"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]interface{}, size)
			if err := jsoniter.ConfigDefault.Unmarshal(payload, &payloads); err != nil {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}
