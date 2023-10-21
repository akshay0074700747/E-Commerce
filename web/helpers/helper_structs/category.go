package helperstructs

type CategoryReq struct {
	Id          uint   `json:"id"`
	Category    string `json:"category"`
	SubCategory string `json:"subcategory"`
	UpdatedBy   string `json:"updatedby"`
}
