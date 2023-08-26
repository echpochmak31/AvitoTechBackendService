package repositories

import "github.com/echpochmak31/avitotechbackendservice/internal/models"

type Repository interface {
	GetAllActiveSegments() ([]models.AbstractSegment, error)
	AddSegment(segmentSlug string) error
	RemoveSegment(segmentSlug string) error
	GetUserSegments(userId int64) ([]models.AbstractSegment, error)
	AddUserSegments(userId int64, segmentSlugs []string) error
	RemoveUserSegments(userId int64, segmentSlugs []string) error
}
