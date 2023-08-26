package models

type AbstractSegment interface {
	GetName() string
}

type SimpleSegment struct {
	Name string
}

func (s SimpleSegment) GetName() string {
	return s.Name
}
