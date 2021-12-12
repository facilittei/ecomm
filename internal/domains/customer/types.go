package customer

// Customer personal information
type Customer struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Document string  `json:"document"`
	Address  Address `json:"address"`
}

// Address for a customer who wants to buy a product
type Address struct {
	Street   string `json:"street"`
	Number   string `json:"number"`
	City     string `json:"city"`
	State    string `json:"state"`
	PostCode string `json:"postCode"`
}
