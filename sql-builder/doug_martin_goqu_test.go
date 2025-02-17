package sqlbuilder

import (
	"testing"

	"github.com/doug-martin/goqu/v9"
)

func BenchmarkGoqu(b *testing.B) {
	benchmarkGoquInsert(b)
	benchmarkGoquSelect(b)
	benchmarkGoquUpdate(b)
	benchmarkGoquDelete(b)
}

func benchmarkGoquInsert(b *testing.B) {
	b.Run("Insert/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			goqu.Insert("dummy_table").
				Cols("name", "address", "age").
				Vals(goqu.Vals{"XYZ", "location-A", 25}).
				ToSQL()
		}
		b.StopTimer()
	})

	b.Run("Insert/Bulk", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			goqu.Insert("dummy_table").
				Cols("name", "address", "age").
				Vals(goqu.Vals{"XYZ", "location-A", 25}).
				Vals(goqu.Vals{"ABC", "location-B", 35}).
				Vals(goqu.Vals{"DEF", "location-C", 40}).
				Vals(goqu.Vals{"PQR", "location-D", 50}).
				ToSQL()
		}
		b.StopTimer()
	})
}

func benchmarkGoquSelect(b *testing.B) {
	b.Run("Select/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			goqu.From("dummy_table").
				Select("id", "name", "address").
				Where(goqu.Ex{"id": 100}).
				ToSQL()
		}
		b.StopTimer()
	})

	b.Run("Select/Join", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			goqu.From("dummy_table").
				Select("id", "name", "address").
				Where(goqu.Ex{"id": 100}).
				Join(goqu.T("other_table"), goqu.On(goqu.I("dummy_table.id").Eq(goqu.I("other_table.id")))).
				ToSQL()
		}
		b.StopTimer()
	})
}

func benchmarkGoquUpdate(b *testing.B) {
	b.Run("Update/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			goqu.Update("dummy_table").
				Set(goqu.Record{"name": "test", "address": "Test Addr"}).
				Where(goqu.Ex{"id": 100}).
				ToSQL()
		}
		b.StopTimer()
	})

	b.Run("Update/Complex", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			goqu.Update("dummy_table").
				Set(goqu.Record{"name": "test", "address": "Test Addr"}).
				Where(goqu.Ex{"id": 100}).
				From("other_table").
				Where(goqu.Ex{"table_one.id": goqu.I("table_two.id")}).
				ToSQL()
		}
		b.StopTimer()
	})
}

func benchmarkGoquDelete(b *testing.B) {
	b.Run("Delete/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			goqu.Delete("dummy_table").
				Where(goqu.Ex{"id": 100}).
				ToSQL()
		}
		b.StopTimer()
	})

	b.Run("Delete/Complex", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s, _, _ := goqu.Select("id").From("other_table").Where(goqu.Ex{"other_table.id": "otherdummy_table_table.id"}).ToSQL()
			goqu.Delete("dummy_table").
				Where(goqu.Ex{"dummy_table.id": s}).
				Where(goqu.Ex{"id": 100}).
				ToSQL()
		}
		b.StopTimer()
	})
}
