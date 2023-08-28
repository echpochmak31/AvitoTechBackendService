package models

import "time"

type AbstractSegment interface {
	GetName() string
}

type SimpleSegment struct {
	Name string
}

func (s SimpleSegment) GetName() string {
	return s.Name
}

type AbstractSegmentWithTTL interface {
	GetName() string
	GetExpirationDate() time.Time
}

type SimpleSegmentWithTTL struct {
	Name           string
	ExpirationDate time.Time
}

func (s SimpleSegmentWithTTL) GetName() string {
	return s.Name
}

func (s SimpleSegmentWithTTL) GetExpirationDate() time.Time {
	return s.ExpirationDate
}
