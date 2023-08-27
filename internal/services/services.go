package services

import "github.com/echpochmak31/avitotechbackendservice/internal/repositories"

type SegmentService struct {
	repository repositories.Repository
}

func NewSegmentService(rep repositories.Repository) *SegmentService {
	s := new(SegmentService)
	s.repository = rep
	return s
}

func (s SegmentService) GetActiveUserSegments(userId int64) ([]string, error) {
	segments, err := s.repository.GetUserSegments(userId)
	if err != nil {
		return nil, err
	}
	segmentSlugs := make([]string, len(segments))
	for i, seg := range segments {
		segmentSlugs[i] = seg.GetName()
	}
	return segmentSlugs, nil
}

func (s SegmentService) CreateNewSegment(segmentName string) error {
	return s.repository.AddSegment(segmentName)
}

func (s SegmentService) DeleteSegment(segmentName string) error {
	return s.repository.RemoveSegment(segmentName)
}

func (s SegmentService) SetUserSegments(userId int64, segmentsToAdd []string, segmentsToRemove []string) error {
	// todo handle expiration date
	err := s.repository.AddUserSegments(userId, segmentsToAdd, nil)
	if err != nil {
		return err
	}
	return s.repository.RemoveUserSegments(userId, segmentsToRemove)
}
