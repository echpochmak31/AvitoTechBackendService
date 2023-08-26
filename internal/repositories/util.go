package repositories

import (
	"context"
	"github.com/echpochmak31/avitotechbackendservice/internal/models"
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
	rep.ConnectionString = connString
	rep.PgxPool = pool

	return rep, nil
}

func (rep *PgxRepository) GetSegmentsWithStatement(statement string, params ...any) ([]models.AbstractSegment, error) {
	rows, err := rep.PgxPool.Query(context.Background(), statement, params)
	if err != nil {
		return make([]models.AbstractSegment, 0), err
	}

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
