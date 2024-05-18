package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/losdmi/timetracker/internal/model"
)

func (r *Repository) ReadDayRecords(ctx context.Context, now time.Time) (model.Records, error) {
	year, month, day := now.Date()
	day_beginning := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

	b := sq.Select(
		"id",
		"task",
		"description",
		"time_start",
		"time_end",
	).PlaceholderFormat(sq.Dollar).
		From("record").
		Where(sq.And{
			sq.GtOrEq{"time_start": day_beginning.UTC()},
			sq.Lt{"time_start": day_beginning.AddDate(0, 0, 1).UTC()},
		}).
		OrderBy("time_start DESC")

	sql, args, err := b.ToSql()
	if err != nil {
		return nil, err
	}

	var result model.Records
	err = pgxscan.Select(ctx, r.dbpool, &result, sql, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
