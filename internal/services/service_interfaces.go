package services

type AbstractSegmentService interface {
	GetActiveUserSegments(userId int64) ([]string, error)
	CreateNewSegment(segmentName string) error
	DeleteSegment(segmentName string) error
	SetUserSegments(userId int64, segmentsToAdd []string, segmentsToRemove []string) error
}
