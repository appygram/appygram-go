package appygram

import (
	"net/http"
)

type AppygramClient struct {
	ApiKey     string
	HttpClient http.Client
}
