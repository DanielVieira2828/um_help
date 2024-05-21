package res

type LoginResponse struct {
	JWS            string `json:"jws"`
	ExpirationTime int64  `json:"expirationTime"`
}
