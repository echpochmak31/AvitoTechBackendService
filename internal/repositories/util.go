package repositories

import (
	"context"
	"github.com/echpochmak31/avitotechbackendservice/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"strconv"
	"time"
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

func toTimestampStr(time time.Time) string {
	return "to_timestamp(" + strconv.FormatInt(time.Unix(), 10) + ")"
}

func (rep *PgxRepository) synchronizeSegmentAndUsers(segmentName string, userPercentage float32) error {
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

	selectUsersWithoutSegmentSql := "SELECT DISTINCT user_id FROM avito.user_segment WHERE segment != $1"
	countUsersWithSegmentSql := "SELECT COUNT (DISTINCT user_id) FROM avito.user_segment WHERE segment = $1"

	var usersWithSegmentsCount int
	users := make([]int64, 0)
	rows, err := tx.Query(context.TODO(), selectUsersWithoutSegmentSql, segmentName)
	if err != nil {
		return err
	}
	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return err
		}
		users = append(users, id)
	}

	err = tx.QueryRow(context.TODO(), countUsersWithSegmentSql, segmentName).Scan(&usersWithSegmentsCount)
	if err != nil {
		return err
	}

	totalUsers := len(users) + usersWithSegmentsCount
	targetUsersCount := int((userPercentage / 100.0) * float32(totalUsers))
	usersToAddCount := targetUsersCount - usersWithSegmentsCount

	insertUsersStatement := "INSERT INTO avito.user_segment (user_id, segment) VALUES ($1, $2) ON CONFLICT DO NOTHING"
	affectedUsersCount := 0
	for i := 0; i < usersToAddCount && i < len(users); i++ {
		_, err = tx.Exec(context.TODO(), insertUsersStatement, users[i], segmentName)
		if err != nil {
			return err
		}
		affectedUsersCount++
	}

	return nil
}

func (rep *PgxRepository) synchronizeAllSegmentsAndUsers() error {
	sql := "SELECT segment, user_percentage FROM avito.segments"
	rows, err := rep.pgxPool.Query(context.Background(), sql)
	if err != nil {
		return err
	}
	for rows.Next() {
		var segment string
		var percentage float32

		err = rows.Scan(&segment, &percentage)
		if err != nil {
			return err
		}

		if percentage != 0.0 {
			err = rep.synchronizeSegmentAndUsers(segment, percentage)
			if err != nil {
				log.Println("Users and segments synchronization failed")
				return err
			}
		}
	}

	return nil
}
