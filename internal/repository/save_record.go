package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/losdmi/timetracker/internal/model"
)

func (r *Repository) SaveRecord(ctx context.Context, record *model.Record) error {
	b := sq.Update("record").PlaceholderFormat(sq.Dollar).
		SetMap(sq.Eq{
			"task":        record.Task,
			"description": record.Description,
			"time_start":  record.TimeStart,
			"time_end":    record.TimeEnd,
		}).
		Where(sq.Eq{"id": record.ID})

	sql, args, err := b.ToSql()
	if err != nil {
		return err
	}

	_, err = r.dbpool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
