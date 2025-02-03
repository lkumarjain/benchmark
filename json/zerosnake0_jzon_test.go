package json

import (
	"testing"

	"github.com/zerosnake0/jzon"
)

func BenchmarkJzon(b *testing.B) {
	b.ReportAllocs()
	for _, tt := range tests {
		benchmarkJzonIterate(b, tt.name, []byte(tt.payload))
		benchmarkJzonStruct(b, tt.name, []byte(tt.payload), tt.size)
		benchmarkJzonInterface(b, tt.name, []byte(tt.payload), tt.size)
	}
}

func benchmarkJzonIterate(b *testing.B, prefix string, payload []byte) {
	b.Run(testName(prefix, "Iterate@Read"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			jzit := jzon.NewIterator()
			defer jzit.Release()
			jzit.ResetBytes(payload)
			jzit.ReadArrayCB(func(i *jzon.Iterator) error {
				i.ReadObjectCB(func(i *jzon.Iterator, field string) error {
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
					return nil
				})
				return nil
			})
		}
		b.StopTimer()
	})
}

func benchmarkJzonStruct(b *testing.B, prefix string, payload []byte, size int) {
	b.Run(testName(prefix, "Struct@Read"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]Payload, size)
			j := 0
			jzit := jzon.NewIterator()
			defer jzit.Release()
			jzit.ResetBytes(payload)
			jzit.ReadArrayCB(func(i *jzon.Iterator) error {
				result := Payload{}
				i.ReadObjectCB(func(i *jzon.Iterator, field string) error {
					switch field {
					case idStr:
						result.ID, _ = i.ReadInt()
					case nameStr:
						result.Name, _ = i.ReadString()
					case cityStr:
						result.City, _ = i.ReadString()
					case countryStr:
						result.Country, _ = i.ReadString()
					case currencyStr:
						result.Currency, _ = i.ReadString()
					}
					return nil
				})
				payloads[j] = result
				j++
				return nil
			})
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Struct@Unmarshal"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]Payload, size)
			if err := jzon.Unmarshal(payload, &payloads); err != nil {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}

func benchmarkJzonInterface(b *testing.B, prefix string, payload []byte, size int) {
	b.Run(testName(prefix, "Interface@Read"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]interface{}, size)
			j := 0
			jzit := jzon.NewIterator()
			defer jzit.Release()
			jzit.ResetBytes(payload)
			jzit.ReadArrayCB(func(i *jzon.Iterator) error {
				result := map[string]interface{}{}
				i.ReadObjectCB(func(i *jzon.Iterator, field string) error {
					switch field {
					case idStr:
						result[idStr], _ = i.ReadInt()
					case nameStr:
						result[nameStr], _ = i.ReadString()
					case cityStr:
						result[cityStr], _ = i.ReadString()
					case countryStr:
						result[countryStr], _ = i.ReadString()
					case currencyStr:
						result[currencyStr], _ = i.ReadString()
					}
					return nil
				})
				payloads[j] = result
				j++
				return nil
			})
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Interface@Unmarshal"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]interface{}, size)
			if err := jzon.Unmarshal(payload, &payloads); err != nil {
				b.Fatal(err)
			}
		}
		b.StopTimer()
	})
}
