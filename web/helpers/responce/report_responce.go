package responce

import (
	"github.com/lib/pq"
)

type ReportResponce struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

// type MultiString []string

// func (s *MultiString) Scan(src interface{}) error {
// 	str, ok := src.(string)
// 	if !ok {
// 		return errors.New("failed to scan multistring field - source is not a string")
// 	}
// 	*s = strings.Split(str, ",")
// 	return nil
// }

// func (s MultiString) Value() (driver.Value, error) {
// 	if s == nil || len(s) == 0 {
// 		return nil, nil
// 	}
// 	return strings.Join(s, ","), nil
// }

type DetailReportResponce struct {
	Email       string         `json:"email"`
	Reports     uint           `json:"reports"`
	Description pq.StringArray `json:"description"`
}
