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
	activeSegments, err := s.repository.GetAllActiveSegments()
	if err != nil {
		return err
	}
	set := make(map[string]bool)
	for _, activeSegment := range activeSegments {
		set[activeSegment.GetName()] = true
	}

	checkedSegmentsToAdd := make([]string, 0)
	checkedSegmentsToRemove := make([]string, 0)
	for _, segment := range segmentsToAdd {
		if set[segment] {
			checkedSegmentsToAdd = append(checkedSegmentsToAdd, segment)
		}
	}
	for _, segment := range segmentsToRemove {
		if set[segment] {
			checkedSegmentsToRemove = append(checkedSegmentsToRemove, segment)
		}
	}

	err = s.repository.AddUserSegments(userId, checkedSegmentsToAdd, nil)
	if err != nil {
		return err
	}
	return s.repository.RemoveUserSegments(userId, checkedSegmentsToRemove)
}
