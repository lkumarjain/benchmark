package json

import (
	"testing"

	"github.com/tidwall/gjson"
)

func BenchmarkGjson(b *testing.B) {
	b.ReportAllocs()
	for _, tt := range tests {
		benchmarkGjsonIterate(b, tt.name, []byte(tt.payload))
		benchmarkGjsonStruct(b, tt.name, []byte(tt.payload), tt.size)
		benchmarkGjsonInterface(b, tt.name, []byte(tt.payload), tt.size)
	}
}

func benchmarkGjsonIterate(b *testing.B, prefix string, payload []byte) {
	b.Run(testName(prefix, "Iterate@Read"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			gjson.GetBytes(payload, "").ForEach(func(key, value gjson.Result) bool {
				value.IsObject()
				value.Get(idStr).Int()
				value.Get(nameStr)
				value.Get(cityStr)
				value.Get(countryStr)
				value.Get(currencyStr)
				return true
			})
		}
		b.StopTimer()
	})
}

func benchmarkGjsonStruct(b *testing.B, prefix string, payload []byte, size int) {
	b.Run(testName(prefix, "Struct@Read"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]Payload, size)
			j := 0
			gjson.GetBytes(payload, "").ForEach(func(key, value gjson.Result) bool {
				result := Payload{}
				value.IsObject()
				result.ID = int(value.Get(idStr).Int())
				result.Name = value.Get(nameStr).String()
				result.City = value.Get(cityStr).String()
				result.Country = value.Get(countryStr).String()
				result.Currency = value.Get(currencyStr).String()
				payloads[j] = result
				j++
				return true
			})
		}
		b.StopTimer()
	})
}

func benchmarkGjsonInterface(b *testing.B, prefix string, payload []byte, size int) {
	b.Run(testName(prefix, "Interface@Read"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]interface{}, size)
			j := 0
			gjson.GetBytes(payload, "").ForEach(func(key, value gjson.Result) bool {
				payloads[j] = value.Value().(map[string]interface{})
				j++
				return true
			})
		}
		b.StopTimer()
	})
}
