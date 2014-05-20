package appygram

import (
	"net/http"
)

type AppygramClient struct {
	ApiKey     string
	HttpClient http.Client
}

func (ac *AppygramClient) getTopics() []string {
	s := make([]string, 1)
	return s
}

func (ac *AppygramClient) sendAppygramMessage(m AppygramMessage) (AppygramResult, error) {
	return AppygramResult{false}, nil
}
