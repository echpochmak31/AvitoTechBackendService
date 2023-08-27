package repositories

import (
	"context"
	"github.com/echpochmak31/avitotechbackendservice/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgxRepository(ctx context.Context, connString string) (*PgxRepository, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	cfg.ConnConfig.RuntimeParams["timezone"] = "Europe/Moscow"
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	rep := new(PgxRepository)
	rep.connectionString = connString
	rep.pgxPool = pool

	return rep, nil
}

func mapToSegments(rows pgx.Rows) ([]models.AbstractSegment, error) {
	segments := make([]models.AbstractSegment, 0)
	for rows.Next() {
		var segName string
		err := rows.Scan(&segName)
		if err != nil {
			return segments, err
		}
		seg := models.SimpleSegment{Name: segName}
		segments = append(segments, seg)
	}
	return segments, nil
}
