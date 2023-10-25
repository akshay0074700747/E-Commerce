package responce

type ReportResponce struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

type DetailReportResponce struct {
	Email       string   `json:"email"`
	Reports     uint     `json:"reports"`
	Description []string `json:"description"`
}
