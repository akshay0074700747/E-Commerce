package helperstructs

type ReviewReq struct {
	ID          uint `json:"id"`
	Product     uint `json:"product"`
	ReviewedBy  string
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
}
