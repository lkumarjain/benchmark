package sqlbuilder

import (
	"testing"

	"github.com/Masterminds/squirrel"
)

func BenchmarkSquirrel(b *testing.B) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	benchmarkSquirrelInsert(b, builder)
	benchmarkSquirrelSelect(b, builder)
	benchmarkSquirrelUpdate(b, builder)
	benchmarkSquirrelDelete(b, builder)
}

func benchmarkSquirrelInsert(b *testing.B, builder squirrel.StatementBuilderType) {
	b.Run("Insert/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			builder.Insert("dummy_table").
				Columns("name", "address", "age").
				Values("XYZ", "location-A", 25).
				ToSql()
		}
		b.StopTimer()
	})

	b.Run("Insert/Bulk", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			builder.Insert("dummy_table").
				Columns("name", "address", "age").
				Values("XYZ", "location-A", 25).
				Values("ABC", "location-B", 35).
				Values("DEF", "location-C", 40).
				Values("PQR", "location-D", 50).
				ToSql()
		}
		b.StopTimer()
	})
}

func benchmarkSquirrelSelect(b *testing.B, builder squirrel.StatementBuilderType) {
	b.Run("Select/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			builder.Select("id", "name", "address").
				From("dummy_table").
				Where(squirrel.Eq{"id": 100}).
				ToSql()
		}
		b.StopTimer()
	})

	b.Run("Select/Join", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			builder.Select("id", "name", "address").
				From("dummy_table").
				Where(squirrel.Eq{"id": 100}).
				Join("other_table ON dummy_table.id = other_table.id").
				ToSql()
		}
		b.StopTimer()
	})
}

func benchmarkSquirrelUpdate(b *testing.B, builder squirrel.StatementBuilderType) {
	b.Run("Update/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			builder.Update("dummy_table").
				Set("name", "test").
				Set("address", "Test Addr").
				Where(squirrel.Eq{"id": 100}).
				ToSql()
		}
		b.StopTimer()
	})

	b.Run("Update/Complex", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			builder.Update("dummy_table").
				Set("age", squirrel.Expr("age + 1")).
				From("other_table").
				Where("other_table.name", "XYZ").
				Where("dummy_table.id = other_table.id").
				ToSql()
		}
		b.StopTimer()
	})
}

func benchmarkSquirrelDelete(b *testing.B, builder squirrel.StatementBuilderType) {
	b.Run("Delete/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			builder.Delete("dummy_table").
				Where(squirrel.Eq{"id": 100}).
				ToSql()
		}
		b.StopTimer()
	})

	b.Run("Delete/Complex", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			builder.Delete("dummy_table").
				From("other_table").
				Where("other_table.name", "XYZ").
				Where("dummy_table.id = other_table.id").
				ToSql()
		}
		b.StopTimer()
	})

}
