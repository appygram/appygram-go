package appygram

import (
	"net/http"
)

func Client(apiKey string, client http.Client) AppygramClient {
	return AppygramClient{apiKey, client}
}
