package requests

type AddUserToSegmentsRequest struct {
	UserId           int64
	SegmentsToAdd    []string
	SegmentsToRemove []string
}
