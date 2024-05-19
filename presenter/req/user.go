package req

type NewUser struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	DocumentNumber string `json:"document_number"`
	Password       string `json:"password"`
}
