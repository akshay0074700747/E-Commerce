package helperstructs

type UserReq struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	ReferralId  uint
	RefferedBy  uint   `json:"referred_by"`
	OldPassword string `json:"old_password"`
	Mobile      string `json:"mobile"`
	Otp         string `json:"otp"`
}

type AddressReq struct {
	ID            uint   `json:"id"`
	Email         string `json:"email"`
	HouseName     string `json:"house_name"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	District      string `json:"district"`
	PO            string `json:"po"`
	State         string `json:"state"`
}

type AddrID struct {
	ID uint `json:"id"`
}
