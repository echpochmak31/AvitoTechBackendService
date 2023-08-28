package requests

import "time"

type FormReportRequest struct {
	StartDate time.Time
	EndDate   time.Time
}
