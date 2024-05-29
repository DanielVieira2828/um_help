package res

type LoginResponse struct {
	JWS            string `json:"jws"`
	ExpirationTime int64  `json:"expiration_time"`
	RefreshToken   string `json:"refresh_token"`
}
