package responce

type UserData struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Wallet    int    `json:"wallet"`
	Isblocked bool   `json:"isblocked"`
}

type AddressData struct {
	ID            uint   `json:"id"`
	Email         string `json:"email"`
	HouseName     string `json:"house_name"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	District      string `json:"district"`
	PO            string `json:"po"`
	State         string `json:"state"`
}

type UserResponce struct {
	Userdata UserData
	Address  []AddressData
	Orders   []OrderData
}
