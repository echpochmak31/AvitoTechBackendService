package repositories

import (
	"context"
	"github.com/echpochmak31/avitotechbackendservice/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type PgxRepository struct {
	connectionString string
	pgxPool          *pgxpool.Pool
}

func (rep *PgxRepository) Close() {
	rep.pgxPool.Close()
}

func (rep *PgxRepository) GetAllActiveSegments() ([]models.AbstractSegment, error) {
	statement := "SELECT segment FROM avito.segments"
	rows, err := rep.pgxPool.Query(context.Background(), statement)
	if err != nil {
		return make([]models.AbstractSegment, 0), err
	}
	return mapToSegments(rows)
}

func (rep *PgxRepository) AddSegment(segmentSlug string) error {
	statement := "INSERT INTO avito.segments (segment) values ($1) ON CONFLICT DO NOTHING"

	tx, err := rep.pgxPool.BeginTx(context.TODO(), pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(context.TODO())
		} else {
			_ = tx.Commit(context.TODO())
		}
	}()

	_, err = tx.Exec(context.TODO(), statement, segmentSlug)
	if err != nil {
		return err
	}

	return nil
}

func (rep *PgxRepository) RemoveSegment(segmentSlug string) error {
	statement1 := "DELETE FROM avito.segments WHERE segment = ($1)"
	statement2 := "UPDATE avito.user_segment SET deleted_at = NOW() WHERE segment = ($1)"
	tx, err := rep.pgxPool.BeginTx(context.TODO(), pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(context.TODO())
		} else {
			_ = tx.Commit(context.TODO())
		}
	}()

	_, err = tx.Exec(context.TODO(), statement1, segmentSlug)
	if err != nil {
		return err
	}
	_, err = tx.Exec(context.TODO(), statement2, segmentSlug)
	if err != nil {
		return err
	}
	return nil
}

func (rep *PgxRepository) GetUserSegments(userId int64) ([]models.AbstractSegment, error) {
	statement :=
		"SELECT segment FROM avito.user_segment " +
			"WHERE user_id = $1 AND deleted_at IS NULL AND (expired_at IS NULL OR expired_at > NOW())"
	rows, err := rep.pgxPool.Query(context.Background(), statement, userId)
	if err != nil {
		return make([]models.AbstractSegment, 0), err
	}
	return mapToSegments(rows)
}

func (rep *PgxRepository) AddUserSegments(userId int64, segmentSlugs []string, expirationDate *time.Time) error {
	var statement string
	if expirationDate == nil {
		statement = "INSERT INTO avito.user_segment (user_id, segment) VALUES ($1, $2)"
	} else {
		statement = "INSERT INTO avito.user_segment (user_id, segment, expired_at) VALUES ($1, $2, $3)"
	}
	for _, segmentSlug := range segmentSlugs {
		var err error
		if expirationDate == nil {
			_, err = rep.pgxPool.Query(context.Background(), statement, userId, segmentSlug)
		} else {
			_, err = rep.pgxPool.Query(context.Background(), statement, userId, segmentSlug, *expirationDate)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (rep *PgxRepository) RemoveUserSegments(userId int64, segmentSlugs []string) error {
	statement := "UPDATE avito.user_segment SET deleted_at = NOW() WHERE user_id = $1 AND segment = $2"
	for _, segmentSlug := range segmentSlugs {
		_, err := rep.pgxPool.Query(context.Background(), statement, userId, segmentSlug)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rep *PgxRepository) MakeReportFile(startDate time.Time, endDate time.Time, pathToReport string) error {
	statement := "COPY (SELECT user_id, segment, " +
		"CASE WHEN deleted_at IS NULL THEN 'Add' ELSE 'Remove' END AS operation, " +
		"CASE WHEN deleted_at IS NULL THEN created_at ELSE deleted_at END AS date " +
		"FROM avito.user_segment " +
		"WHERE created_at >= " + toTimestampStr(startDate) + " AND created_at <= " + toTimestampStr(endDate) +
		"AND (deleted_at IS NULL OR (deleted_at >= " + toTimestampStr(startDate) +
		"AND deleted_at <= " + toTimestampStr(endDate) + "))" +
		"ORDER BY date)" + "TO '" + pathToReport + "' WITH CSV"

	_, err := rep.pgxPool.Query(context.Background(), statement)
	return err
}

func (rep *PgxRepository) DeleteExpiredSegments() error {
	statement := "UPDATE avito.user_segment SET deleted_at = NOW() WHERE deleted_at IS NULL AND expired_at < NOW()"
	_, err := rep.pgxPool.Query(context.Background(), statement)
	return err
}
