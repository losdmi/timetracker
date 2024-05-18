package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/losdmi/timetracker/internal/model"
)

func (r *Repository) CreateRecord(ctx context.Context, record model.Record) error {
	b := sq.Insert("record").PlaceholderFormat(sq.Dollar).
		Columns(
			"task",
			"description",
			"time_start",
		).
		Values(
			record.Task,
			record.Description,
			record.TimeStart,
		)

	sql, args, err := b.ToSql()
	if err != nil {
		return err
	}

	_, err = r.dbpool.Exec(ctx, sql, args...)

	return err
}
