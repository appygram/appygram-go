package appygram

import (
	"net/http"
)

func Client(apiKey string, client http.Client) AppygramClient {
	return AppygramClient{apiKey, client}
}

func MakeErrorBackTrace(stackBuff []byte) [][]string {
	return make([][]string, 1)
}
