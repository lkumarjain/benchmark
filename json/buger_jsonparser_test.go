package json

import (
	"bytes"
	"testing"

	"github.com/buger/jsonparser"
)

func BenchmarkJsonParser(b *testing.B) {
	b.ReportAllocs()
	for _, tt := range tests {
		benchmarkJsonParserIterate(b, tt.name, []byte(tt.payload))
		benchmarkJsonParserStruct(b, tt.name, []byte(tt.payload), tt.size)
		benchmarkJsonParserInterface(b, tt.name, []byte(tt.payload), tt.size)
	}
}

func benchmarkJsonParserIterate(b *testing.B, prefix string, payload []byte) {
	b.Run(testName(prefix, "Iterate@Object"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			jsonparser.ArrayEach(payload, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				jsonparser.ObjectEach(value, func(key, value []byte, vt jsonparser.ValueType, off int) error {
					switch {
					case bytes.Equal(key, idByte):
						jsonparser.ParseInt(value)
					case bytes.Equal(key, nameByte):
						jsonparser.ParseString(value)
					case bytes.Equal(key, cityByte):
						jsonparser.ParseString(value)
					case bytes.Equal(key, countryByte):
						jsonparser.ParseString(value)
					case bytes.Equal(key, currencyByte):
						jsonparser.ParseString(value)
					}
					return nil
				})
			})
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Iterate@Get"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			jsonparser.ArrayEach(payload, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				jsonparser.GetInt(value, idStr)
				jsonparser.GetString(value, nameStr)
				jsonparser.GetString(value, cityStr)
				jsonparser.GetString(value, countryStr)
				jsonparser.GetString(value, currencyStr)
			})
		}
		b.StopTimer()
	})
}

func benchmarkJsonParserStruct(b *testing.B, prefix string, payload []byte, size int) {
	b.Run(testName(prefix, "Struct@Object"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]Payload, size)
			j := 0
			jsonparser.ArrayEach(payload, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				result := Payload{}
				jsonparser.ObjectEach(value, func(key, value []byte, vt jsonparser.ValueType, off int) error {
					switch {
					case bytes.Equal(key, idByte):
						v, _ := jsonparser.ParseInt(value)
						result.ID = int(v)
					case bytes.Equal(key, nameByte):
						v, _ := jsonparser.ParseString(value)
						result.Name = v
					case bytes.Equal(key, cityByte):
						v, _ := jsonparser.ParseString(value)
						result.City = v
					case bytes.Equal(key, countryByte):
						v, _ := jsonparser.ParseString(value)
						result.Country = v
					case bytes.Equal(key, currencyByte):
						v, _ := jsonparser.ParseString(value)
						result.Currency = v
					}

					payloads[j] = result
					return nil
				})
				j++
			})
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Struct@Get"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]Payload, size)
			j := 0
			jsonparser.ArrayEach(payload, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				result := Payload{}
				v, _ := jsonparser.GetInt(value, idStr)
				result.ID = int(v)
				result.Name, _ = jsonparser.GetString(value, nameStr)
				result.City, _ = jsonparser.GetString(value, cityStr)
				result.Country, _ = jsonparser.GetString(value, countryStr)
				result.Currency, _ = jsonparser.GetString(value, currencyStr)
				payloads[j] = result
				j++
			})
		}
		b.StopTimer()
	})
}

func benchmarkJsonParserInterface(b *testing.B, prefix string, payload []byte, size int) {
	b.Run(testName(prefix, "Interface@Object"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]interface{}, size)
			j := 0
			jsonparser.ArrayEach(payload, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				result := map[string]interface{}{}
				jsonparser.ObjectEach(value, func(key, value []byte, vt jsonparser.ValueType, off int) error {
					switch {
					case bytes.Equal(key, idByte):
						result[idStr], _ = jsonparser.ParseInt(value)
					case bytes.Equal(key, nameByte):
						result[nameStr], _ = jsonparser.ParseString(value)
					case bytes.Equal(key, cityByte):
						result[cityStr], _ = jsonparser.ParseString(value)
					case bytes.Equal(key, countryByte):
						result[countryStr], _ = jsonparser.ParseString(value)
					case bytes.Equal(key, currencyByte):
						result[currencyStr], _ = jsonparser.ParseString(value)
					}

					payloads[j] = result
					return nil
				})
				j++
			})
		}
		b.StopTimer()
	})

	b.Run(testName(prefix, "Interface@Get"), func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			payloads := make([]interface{}, size)
			j := 0
			jsonparser.ArrayEach(payload, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				result := map[string]interface{}{}
				result[idStr], _ = jsonparser.GetInt(value, idStr)
				result[nameStr], _ = jsonparser.GetString(value, nameStr)
				result[cityStr], _ = jsonparser.GetString(value, cityStr)
				result[countryStr], _ = jsonparser.GetString(value, countryStr)
				result[currencyStr], _ = jsonparser.GetString(value, currencyStr)
				payloads[j] = result
				j++
			})
		}
		b.StopTimer()
	})
}
