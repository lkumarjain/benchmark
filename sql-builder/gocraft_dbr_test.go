package sqlbuilder

import (
	"testing"

	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
)

func BenchmarkDbr(b *testing.B) {
	session := &dbr.Session{Connection: &dbr.Connection{Dialect: dialect.PostgreSQL}}
	benchmarkDbrInsert(b, session)
	benchmarkDbrSelect(b, session)
	benchmarkDbrUpdate(b, session)
	benchmarkDbrDelete(b, session)
}

func benchmarkDbrInsert(b *testing.B, session *dbr.Session) {
	b.Run("Insert/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			session.InsertInto("dummy_table").
				Columns("name", "address", "age").
				Values("XYZ", "location-A", 25).
				Build(dialect.PostgreSQL, dbr.NewBuffer())
		}
		b.StopTimer()
	})

	b.Run("Insert/Bulk", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			session.InsertInto("dummy_table").
				Columns("name", "address", "age").
				Values("XYZ", "location-A", 25).
				Values("ABC", "location-B", 35).
				Values("DEF", "location-C", 40).
				Values("PQR", "location-D", 50).
				Build(dialect.PostgreSQL, dbr.NewBuffer())
		}
		b.StopTimer()
	})
}

func benchmarkDbrSelect(b *testing.B, session *dbr.Session) {
	b.Run("Select/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			session.Select("id", "name", "address").
				From("dummy_table").
				Where("id IN ?", []int{100, 200, 300}).
				Build(dialect.PostgreSQL, dbr.NewBuffer())
		}
		b.StopTimer()
	})

	b.Run("Select/Join", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			session.Select("id", "name", "address").
				From("dummy_table").
				Where("id IN ?", []int{100, 200, 300}).
				LeftJoin("other_table", "dummy_table.id = other_table.id").
				Build(dialect.PostgreSQL, dbr.NewBuffer())
		}
		b.StopTimer()
	})
}

func benchmarkDbrUpdate(b *testing.B, session *dbr.Session) {
	b.Run("Update/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			session.Update("dummy_table").
				Set("name", "test").
				Set("address", "Test Addr").
				Where("id", 100).
				Build(dialect.PostgreSQL, dbr.NewBuffer())
		}
		b.StopTimer()
	})

	b.Run("Update/Complex", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			session.UpdateBySql(`UPDATE "dummy_table" SET "foo"="other_table"."bar" FROM "other_table" WHERE ("dummy_table"."id" = "other_table"."id"`).
				Build(dialect.PostgreSQL, dbr.NewBuffer())
		}
		b.StopTimer()
	})
}

func benchmarkDbrDelete(b *testing.B, session *dbr.Session) {
	b.Run("Delete/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			session.DeleteFrom("dummy_table").
				Where("id", 100).
				Build(dialect.PostgreSQL, dbr.NewBuffer())
		}
		b.StopTimer()
	})

	b.Run("Delete/Complex", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			session.DeleteFrom("dummy_table").
				Where("id", session.Select("id").From("other_table").Where("name", "XYZ")).
				Build(dialect.PostgreSQL, dbr.NewBuffer())
		}
		b.StopTimer()
	})
}
