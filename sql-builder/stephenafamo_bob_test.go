package sqlbuilder

import (
	"context"
	"testing"

	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dm"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/bob/dialect/psql/um"
)

func BenchmarkBob(b *testing.B) {
	benchmarkBobInsert(b)
	benchmarkBobSelect(b)
	benchmarkBobUpdate(b)
	benchmarkBobDelete(b)
}

func benchmarkBobInsert(b *testing.B) {
	b.Run("Insert/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			psql.Insert(
				im.Into("dummy_table"),
				im.Values(psql.Arg("XYZ", "location-A", 25)),
			).Build(context.Background())
		}
		b.StopTimer()
	})

	b.Run("Insert/Bulk", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			psql.Insert(
				im.Into("dummy_table", "name", "address", "age"),
				im.Values(psql.Arg("XYZ", "location-A", 25)),
				im.Values(psql.Arg("ABC", "location-B", 35)),
				im.Values(psql.Arg("DEF", "location-C", 40)),
				im.Values(psql.Arg("PQR", "location-D", 50)),
			).Build(context.Background())
		}
		b.StopTimer()
	})
}

func benchmarkBobSelect(b *testing.B) {
	b.Run("Select/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			psql.Select(
				sm.Columns("id", "name", "address"),
				sm.From("dummy_table"),
				sm.Where(psql.Quote("id").In(psql.Arg(100, 200, 300))),
			).Build(context.Background())
		}
		b.StopTimer()
	})

	b.Run("Select/Join", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			psql.Select(
				sm.Columns("id", "name", "address"),
				sm.From("dummy_table"),
				sm.Where(psql.Quote("id").In(psql.Arg(100, 200, 300))),
				sm.LeftJoin("other_table").Using("id"),
			).Build(context.Background())
		}
		b.StopTimer()
	})
}

func benchmarkBobUpdate(b *testing.B) {
	b.Run("Update/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			psql.Update(
				um.Table("dummy_table"),
				um.SetCol("name").ToArg("test"),
				um.SetCol("address").ToArg("Test Addr"),
				um.Where(psql.Quote("id").EQ(psql.Arg(100))),
			).Build(context.Background())
		}
		b.StopTimer()
	})

	b.Run("Update/Complex", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			psql.Update(
				um.Table("dummy_table"),
				um.SetCol("age").To("age + 1"),
				um.From("other_table"),
				um.Where(psql.Quote("other_table", "name").EQ(psql.Arg("XYZ"))),
				um.Where(psql.Quote("dummy_table", "id").EQ(psql.Quote("other_table", "id"))),
			).Build(context.Background())
		}
		b.StopTimer()
	})
}

func benchmarkBobDelete(b *testing.B) {
	b.Run("Delete/Simple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			psql.Delete(
				dm.From("dummy_table"),
				dm.Where(psql.Quote("id").In(psql.Arg(100))),
			).Build(context.Background())
		}
		b.StopTimer()
	})

	b.Run("Delete/Complex", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			psql.Delete(
				dm.From("dummy_table"),
				dm.Using("other_table"),
				dm.Where(psql.Quote("dummy_table", "name").EQ(psql.Arg("XYZ"))),
				dm.Where(psql.Quote("dummy_table", "id").EQ(psql.Quote("other_table", "id"))),
			).Build(context.Background())
		}
		b.StopTimer()
	})
}
