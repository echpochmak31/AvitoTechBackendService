package requests

import (
	"time"
)

type SegmentWithTTl struct {
	Name           string
	ExpirationDate time.Time
}

type AddUserToSegmentsRequest struct {
	UserId           int64
	SegmentsToAdd    []SegmentWithTTl
	SegmentsToRemove []string
}
