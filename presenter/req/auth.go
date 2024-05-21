package req

type LoginRequest struct {
	DocumentNumber string `json:"document_number"`
	Password       string `json:"password"`
}
