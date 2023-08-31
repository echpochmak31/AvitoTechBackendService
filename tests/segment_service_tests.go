package tests

import (
	"github.com/echpochmak31/avitotechbackendservice/internal/models"
	"github.com/echpochmak31/avitotechbackendservice/internal/services"
	"github.com/echpochmak31/avitotechbackendservice/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func PrepareMocks(t *testing.T) *mocks.MockRepository {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)
	return mockRepo
}

func TestGetUserSegments(t *testing.T) {
	var userId int64 = 1001
	var segName1, segName2 = "SEGMENT_ONE", "SEGMENT_TWO"
	seg1 := models.SimpleSegment{
		Name: segName1,
	}
	seg2 := models.SimpleSegment{
		Name: segName2,
	}

	mockRepo := PrepareMocks(t)
	segmentService := services.NewSegmentService(mockRepo)

	mockRepo.EXPECT().GetUserSegments(userId).Return([]models.AbstractSegment{seg1, seg2}, nil)

	segs, err := segmentService.GetActiveUserSegments(userId)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, segs[0], segName1)
	assert.Equal(t, segs[1], segName2)
}

func TestSynchronizeSegments(t *testing.T) {
	var segName1, segName2 = "SEGMENT_ONE", "SEGMENT_TWO"
	seg1WithTTL := models.SimpleSegmentWithTTL{
		Name:           segName1,
		ExpirationDate: time.Time{},
	}
	seg2WithTTL := models.SimpleSegmentWithTTL{
		Name:           segName2,
		ExpirationDate: time.Now().Add(time.Hour * 5),
	}
	seg1 := models.SimpleSegment{
		Name: segName1,
	}
	seg2 := models.SimpleSegment{
		Name: segName2,
	}
	segmentsWithTTLInput := []models.AbstractSegmentWithTTL{seg1WithTTL, seg2WithTTL}
	segmentsWithTTLSlice := make([]models.AbstractSegmentWithTTL, 0)
	activeSegments := []models.AbstractSegment{seg1, seg2}
	var userId int64 = 1001

	mockRepo := PrepareMocks(t)
	segmentService := services.NewSegmentService(mockRepo)

	mockRepo.EXPECT().GetAllActiveSegments().Return(activeSegments, nil).Times(1)

	mockRepo.EXPECT().AddUserSegments(userId, segmentsWithTTLInput).Do(func(int64, []models.AbstractSegmentWithTTL) {
		for _, seg := range segmentsWithTTLInput {
			segmentsWithTTLSlice = append(segmentsWithTTLSlice, seg)
		}
	}).Times(1)

	mockRepo.EXPECT().RemoveUserSegments(userId, []string{}).Return(nil).Times(1)

	mockRepo.EXPECT().DeleteExpiredSegments().Do(func() {
		for i, seg := range segmentsWithTTLSlice {
			if seg.GetExpirationDate().Compare(time.Now()) < 0 {
				segmentsWithTTLSlice[i] = segmentsWithTTLSlice[len(segmentsWithTTLSlice)-1]
				segmentsWithTTLSlice = segmentsWithTTLSlice[:len(segmentsWithTTLSlice)-1]
			}
		}
	}).Times(1)

	err := segmentService.SetUserSegments(userId, segmentsWithTTLInput, nil)
	if err != nil {
		t.Fail()
	}

	ticker := time.NewTicker(2 * time.Second)
	go segmentService.SynchronizeSegments(ticker)
	time.Sleep(3 * time.Second)
	ticker.Stop()

	assert.True(t, len(segmentsWithTTLSlice) == 1)
	assert.Equal(t, segmentsWithTTLSlice[0], seg2WithTTL)
	log.Println(len(segmentsWithTTLSlice))
}
