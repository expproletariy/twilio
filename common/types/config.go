package types

import "net/http"

type Config struct {
	Client             *http.Client
	BaseApiUrl         string
	ApiVersion         string
	TwilioApiKeySid    string
	TwilioApiKeySecret string
}
