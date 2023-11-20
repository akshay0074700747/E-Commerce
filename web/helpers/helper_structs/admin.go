package helperstructs

import "time"

type AdminReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type SalesReportTime struct {
	Starttime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
