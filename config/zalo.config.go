package config

type ZaloTokenConfig struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
}

var ZaloToken ZaloTokenConfig
