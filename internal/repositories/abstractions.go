package repositories

import (
	"github.com/echpochmak31/avitotechbackendservice/internal/models"
	"time"
)

type Repository interface {
	GetAllActiveSegments() ([]models.AbstractSegment, error)
	AddSegment(segmentSlug string) error
	RemoveSegment(segmentSlug string) error
	GetUserSegments(userId int64) ([]models.AbstractSegment, error)
	AddUserSegments(userId int64, segmentSlugs []string, expirationDate *time.Time) error
	RemoveUserSegments(userId int64, segmentSlugs []string) error
	DeleteExpiredSegments() error
}

type ReportsRepository interface {
	MakeReportFile(startDate time.Time, endDate time.Time, pathToReport string) error
}
