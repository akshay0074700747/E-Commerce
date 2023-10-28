package helperstructs

import "time"

type ReportReq struct {
	Email       string `json:"email"`
	Description string `json:"description"`
}

type BlockReq struct {
	Email string `json:"email"`
	Time  time.Time `json:"time"`
}
