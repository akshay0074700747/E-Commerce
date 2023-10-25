package helperstructs

import "time"

type ReportReq struct {
	Email       string
	Description string `json:"description"`
}

type BlockReq struct {
	Email string
	Time  time.Time `json:"time"`
}
