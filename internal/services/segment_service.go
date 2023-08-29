package services

import (
	"github.com/echpochmak31/avitotechbackendservice/internal/models"
	"github.com/echpochmak31/avitotechbackendservice/internal/repositories"
	"log"
	"time"
)

type SegmentService struct {
	repository repositories.Repository
}

func NewSegmentService(rep repositories.Repository) *SegmentService {
	s := new(SegmentService)
	s.repository = rep
	return s
}

func (s *SegmentService) GetActiveUserSegments(userId int64) ([]string, error) {
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

func (s *SegmentService) CreateNewSegment(segmentName string, userPercentage float32) error {
	return s.repository.AddSegment(segmentName, userPercentage)
}

func (s *SegmentService) DeleteSegment(segmentName string) error {
	return s.repository.RemoveSegment(segmentName)
}

func (s *SegmentService) SetUserSegments(
	userId int64,
	segmentsToAdd []models.AbstractSegmentWithTTL,
	segmentsToRemove []string) error {

	activeSegments, err := s.repository.GetAllActiveSegments()
	if err != nil {
		return err
	}
	set := make(map[string]bool)
	for _, activeSegment := range activeSegments {
		set[activeSegment.GetName()] = true
	}

	checkedSegmentsToAdd := make([]models.AbstractSegmentWithTTL, 0)
	checkedSegmentsToRemove := make([]string, 0)
	for _, segment := range segmentsToAdd {
		if set[segment.GetName()] {
			checkedSegmentsToAdd = append(checkedSegmentsToAdd, segment)
		}
	}
	for _, segment := range segmentsToRemove {
		if set[segment] {
			checkedSegmentsToRemove = append(checkedSegmentsToRemove, segment)
		}
	}

	err = s.repository.AddUserSegments(userId, checkedSegmentsToAdd)
	if err != nil {
		return err
	}
	err = s.repository.RemoveUserSegments(userId, checkedSegmentsToRemove)
	return err
}

func (s *SegmentService) SynchronizeSegments(ticker *time.Ticker) {
	for range ticker.C {
		err := s.repository.DeleteExpiredSegments()
		if err != nil {
			log.Fatal("Synchronization failed: ", err)
		}
		log.Println("Synchronization successful")
	}
}
