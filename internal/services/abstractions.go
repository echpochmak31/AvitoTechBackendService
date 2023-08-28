package services

import "time"

type AbstractSegmentService interface {
	GetActiveUserSegments(userId int64) ([]string, error)
	CreateNewSegment(segmentName string) error
	DeleteSegment(segmentName string) error
	SetUserSegments(userId int64, segmentsToAdd []string, segmentsToRemove []string) error
	SynchronizeSegments(ticker *time.Ticker)
}

type AbstractReportService interface {
	FormReport(startDate time.Time, endDate time.Time) (string, error)
	GetReportName(startDate time.Time, endDate time.Time) string
	SendReport(handler AbstractReportHandler) error
}

type AbstractReportHandler interface {
	Handle(any any) error
}
