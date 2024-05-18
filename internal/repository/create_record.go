package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/losdmi/timetracker/internal/model"
)

func (r *Repository) CreateRecord(ctx context.Context, record *model.Record, previousRecord *model.Record) error {
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
	if err != nil {
		return err
	}

	if previousRecord != nil {
		err = r.SaveRecord(ctx, previousRecord)
		if err != nil {
			fmt.Printf("Repository.CreateRecord ошибка при сохранении записи %s", err)
			return err
		}
	}

	return nil
}
